package main

import (
	"github.com/ghssni/Smartcy-LMS/User-Service/database"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/service"
	"github.com/ghssni/Smartcy-LMS/User-Service/pb"
	"github.com/ghssni/Smartcy-LMS/User-Service/pkg"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var db *mongo.Database

func main() {
	// Setup logger
	pkg.SetupLogger()
	//.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

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
	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT_USER_SERVICE"))
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
