package main

import (
	"github.com/ghssni/Smartcy-LMS/User-Service/database"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/service"
	pb "github.com/ghssni/Smartcy-LMS/User-Service/pb/proto"
	"github.com/ghssni/Smartcy-LMS/User-Service/pkg"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var db *mongo.Database

func main() {
	// Setup logger
	pkg.SetupLogger()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Initialize database
	var err error
	db, err = database.InitMongoDB()
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	}

	// Set up graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Run gRPC and gRPC Gateway concurrently
	runGrpcServer()
	// Wait for shutdown signal
	<-sigs
	logrus.Println("Shutting down servers...")
}

func runGrpcServer() {
	grpcServer := grpc.NewServer()

	// Initialize repositories and services
	userRepo := repository.NewUserRepo(db)
	userActivityLogRepo := repository.NewUserActivityLogRepo(db)
	userService := service.NewUserService(userRepo, userActivityLogRepo)

	// Register gRPC service
	pb.RegisterUserServiceServer(grpcServer, userService)

	// Get port from environment and start listener
	port := os.Getenv("GRPC_PORT_USER_SERVICE")
	if port == "" {
		port = "50051" // default port if not set
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logrus.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	logrus.Printf("gRPC server is running on port %s", port)

	// Serve gRPC server
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
