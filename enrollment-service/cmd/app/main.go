package main

import (
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/config"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/database"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/service"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/pb"
	helper "github.com/ghssni/Smartcy-LMS/Enrollment-Service/pkg"
	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
	"os"
)

var db *gorm.DB

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed to load env file: %v", err)
	}
	helper.SetupLogger()
	var err error
	db, err = database.InitDB()
	if err != nil {
		logrus.Println("Failed to connect to database: %v", err)
	}

	go runGrpcServer() // Run gRPC server on port 50052

	// run scheduler
	go func() {
		conn, err := grpc.Dial(":50054", grpc.WithInsecure())
		if err != nil {
			logrus.Fatalf("Failed to dial gRPC server: %v", err)
		}
		defer conn.Close()

		paymentClient := pb.NewPaymentsServiceClient(conn)
		scheduler := config.NewScheduler(paymentClient)
		if err := scheduler.Scheduler(); err != nil {
			logrus.Fatalf("Failed to run scheduler: %v", err)
		}
	}()

	select {}
}

func runGrpcServer() {
	// Run gRPC server
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT_ENROLLMENT_SERVICE"))
	if err != nil {
		logrus.Fatalf("Failed to listen on port 50052: %v", err)
	}

	//accessKey := os.Getenv("CRON_ACCESS_KEY")

	grpcServer := grpc.NewServer(
		//grpc.ChainUnaryInterceptor(
		//	middleware.AccessKeyInterceptor(accessKey),
		//),
		grpc.UnaryInterceptor(service.Unary()),
	)

	enrollmentRepo := repository.NewEnrollmentRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	assessmentsRepo := repository.NewAssessmentsRepository(db)
	//certificateRepo := repository.NewCertificateRepository(db)

	// Register gRPC server from service
	pb.RegisterEnrollmentServiceServer(grpcServer, service.NewEnrollmentService(enrollmentRepo, paymentRepo))

	//register gRPC server from service
	pb.RegisterAssessmentsServiceServer(grpcServer, service.NewAssessmentsService(assessmentsRepo))

	//pb.RegisterCertificateServiceServer(grpcServer, service.)

	pb.RegisterPaymentsServiceServer(grpcServer, service.NewPaymentService(paymentRepo))

	// Start gRPC server in a goroutine
	go func() {
		logrus.Println("Starting gRPC server on port 50052")
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	select {}
}
