package main

import (
<<<<<<< HEAD
	"context"
	"github.com/ghssni/Smartcy-LMS/pkg"
	"github.com/ghssni/Smartcy-LMS/user-service/database"
=======
	"github.com/ghssni/Smartcy-LMS/User-Service/database"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/service"
	pb "github.com/ghssni/Smartcy-LMS/User-Service/pb/proto"
	"github.com/ghssni/Smartcy-LMS/User-Service/pkg"
>>>>>>> user-service
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
<<<<<<< HEAD

	"os"
	"os/signal"
	"time"
=======
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
>>>>>>> user-service
)

var db *mongo.Database

func main() {
	//.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	// Setup logger
	pkg.SetupLogger()

	// Config Db
	var err error
	db, err = database.InitMongoDB()
	if err != nil {
		logrus.Fatalf("Error connecting to database: %v", err)
	}

	runGrpcServer()
}

func runGrpcServer() {
	grpcServer := grpc.NewServer()

	// initialize all services
	userRepo := repository.NewUserRepo(db)
	userActivityLogRepo := repository.NewUserActivityLogRepo(db)

	pb.RegisterUserServiceServer(grpcServer, service.NewUserService(userRepo, userActivityLogRepo))
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
