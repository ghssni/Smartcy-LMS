package main

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/config"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/database"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/middleware"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/service"
	helper "github.com/ghssni/Smartcy-LMS/Enrollment-Service/pkg"
	"github.com/joho/godotenv"

	pbAssessments "github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/assessments"
	pbCertificate "github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/certificate"
	pbEnrollment "github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/enrollment"
	pbPayments "github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/payments"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	middlewareEcho "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// init xendit
	config.InitXendit()

	go runGrpcServer() // Run gRPC server on port 50051
	go runGrpcGatewayServer()

	// run scheduler
	go func() {
		conn, err := grpc.Dial(":50054", grpc.WithInsecure())
		if err != nil {
			logrus.Fatalf("Failed to dial gRPC server: %v", err)
		}
		defer conn.Close()

		paymentClient := pbPayments.NewPaymentsServiceClient(conn)
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

	accessKey := os.Getenv("CRON_ACCESS_KEY")

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.JWTInterceptor(os.Getenv("JWT_SECRET")),
			middleware.AccessKeyInterceptor(accessKey),
		),
	)

	enrollmentRepo := repository.NewEnrollmentRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	assessmentsRepo := repository.NewAssessmentsRepository(db)

	// Register gRPC server from service
	pbEnrollment.RegisterEnrollmentServiceServer(grpcServer, service.NewEnrollmentService(enrollmentRepo, paymentRepo))

	//register gRPC server from service
	pbAssessments.RegisterAssessmentsServiceServer(grpcServer, service.NewAssessmentsService(assessmentsRepo))

	pbCertificate.RegisterCertificateServiceServer(grpcServer, pbCertificate.UnimplementedCertificateServiceServer{})

	pbPayments.RegisterPaymentsServiceServer(grpcServer, service.NewPaymentService(paymentRepo))

	// Start gRPC server in a goroutine
	go func() {
		logrus.Println("Starting gRPC server on port 50052")
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	select {}
}

func runGrpcGatewayServer() {
	// Inisialisasi Echo
	e := echo.New()
	e.Use(middlewareEcho.Logger())
	e.Use(middlewareEcho.Recover())
	// Setup gRPC-Gateway mux
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register HTTP handlers for Echo
	servicePayments := service.NewPaymentService(repository.NewPaymentRepository(db))

	e.POST("/v1/payments/webhook", servicePayments.HandleWebhookHTTP)

	// Register HTTP handlers for gRPC-Gateway
	err := pbPayments.RegisterPaymentsServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		logrus.Fatalf("Failed to register gRPC Gateway for Payments service: %v", err)
	}

	e.Any("/*", echo.WrapHandler(mux))

	logrus.Println("Echo server with gRPC-Gateway is running on port 8081")
	if err := e.Start(":8081"); err != nil {
		logrus.Fatalf("Failed to serve Echo server with gRPC-Gateway: %v", err)
	}

	//err := pbEnrollment.RegisterEnrollmentServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	//if err != nil {
	//	logrus.Fatalf("Failed to register gRPC Gateway for Enrollment service: %v", err)
	//}
	//
	//err = pbAssessments.RegisterAssessmentsServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	//if err != nil {
	//	logrus.Fatalf("Failed to register gRPC Gateway for Assessments service: %v", err)
	//}
	//
	//err = pbCertificate.RegisterCertificateServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	//if err != nil {
	//	logrus.Fatalf("Failed to register gRPC Gateway for Certificate service: %v", err)
	//}
	//
	//err = pbLearningProgress.RegisterLearningProgressServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	//if err != nil {
	//	logrus.Fatalf("Failed to register gRPC Gateway for LearningProgress service: %v", err)
	//}

}
