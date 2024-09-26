package main

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/User-Service/database"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/User-Service/internal/service"
	pb "github.com/ghssni/Smartcy-LMS/User-Service/pb/proto"
	"github.com/ghssni/Smartcy-LMS/User-Service/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	middlewareEcho "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	go runGrpcServer()
	go runGrpcGatewayServer()

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

func runGrpcGatewayServer() {
	// Initialize Echo instance
	e := echo.New()
	e.Use(middlewareEcho.Logger())
	e.Use(middlewareEcho.Recover())

	// Setup gRPC-Gateway mux
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Get gRPC server port from environment
	grpcPort := os.Getenv("GRPC_PORT_USER_SERVICE")
	if grpcPort == "" {
		grpcPort = "50051" // default gRPC server port
	}

	// Register gRPC Gateway service
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+grpcPort, opts)
	if err != nil {
		logrus.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	// Register custom HTTP handler (reset password)
	userService := service.NewUserService(repository.NewUserRepo(db), repository.NewUserActivityLogRepo(db))
	e.POST("/reset-password", userService.NewPasswordHTTP)

	// Bind Echo to the same port for HTTP Gateway
	echoPort := os.Getenv("GATEWAY_PORT")
	if echoPort == "" {
		echoPort = "8080" // default Echo port
	}

	logrus.Printf("Echo server with gRPC-Gateway is running on port %s", echoPort)
	if err := e.Start(":" + echoPort); err != nil {
		logrus.Fatalf("Failed to start Echo server: %v", err)
	}
}
