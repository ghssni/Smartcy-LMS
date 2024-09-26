package main

import (
	"github.com/ghssni/Smartcy-LMS/Email-Service/database"
	"github.com/ghssni/Smartcy-LMS/Email-Service/helper"
	"github.com/ghssni/Smartcy-LMS/Email-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/Email-Service/internal/service"
	pb "github.com/ghssni/Smartcy-LMS/Email-Service/pb/proto"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var db *gorm.DB

func main() {
	// Setup logger
	helper.SetupLogger()
	//.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Config Db
	var err error
	db, err = database.InitDB()
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	}

	runGrpcServer()
}

func runGrpcServer() {
	grpcServer := grpc.NewServer()

	// initialize all services
	grpcHost := os.Getenv("GRPC_HOST")
	if grpcHost == "" {
		grpcHost = "localhost"
	}
	grpcPort := os.Getenv("GRPC_PORT_EMAIL_SERVICE")
	if grpcPort == "" {
		grpcPort = "50053"
	}
	address := grpcHost + ":" + grpcPort

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s:%s - %v", grpcHost, grpcPort, err)
	}

	grpcServer = grpc.NewServer()

	//register service
	emailRepo := repository.NewEmailsRepository(db)
	emailLogRepo := repository.NewEmailsLogRepository(db)
	// Register gRPC services
	pb.RegisterEmailServiceServer(grpcServer, service.NewEmailService(emailRepo, emailLogRepo))

	// Graceful shutdown for gRPC server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Gracefully stopping gRPC server...")
		grpcServer.GracefulStop()
	}()
	// Start gRPC server
	log.Printf("gRPC server listening on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
