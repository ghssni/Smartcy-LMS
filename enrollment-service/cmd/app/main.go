package main

import (
	"github.com/ghssni/Smartcy-LMS/enrollment-service/database"
	pbAssessments "github.com/ghssni/Smartcy-LMS/enrollment-service/github.com/ghssni/Smartcy-LMS/enrollment-service/proto/assessments"
	pbCertificate "github.com/ghssni/Smartcy-LMS/enrollment-service/github.com/ghssni/Smartcy-LMS/enrollment-service/proto/certificate"
	pbEnrollment "github.com/ghssni/Smartcy-LMS/enrollment-service/github.com/ghssni/Smartcy-LMS/enrollment-service/proto/enrollment"
	pbLearningProgress "github.com/ghssni/Smartcy-LMS/enrollment-service/github.com/ghssni/Smartcy-LMS/enrollment-service/proto/learningProgress"
	pbPayments "github.com/ghssni/Smartcy-LMS/enrollment-service/github.com/ghssni/Smartcy-LMS/enrollment-service/proto/payments"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	go runGrpcServer() // Run gRPC server on port 50051
	runGrpcGatewayServer()
}

func runGrpcServer() {
	// Run gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Println("Failed to listen: %v", err)
	}

	// initDB
	_, err = database.InitDB()
	if err != nil {
		logrus.Println("Failed to connect to database: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.IpWhiteListInterceptor),
	)

	// Register gRPC server from service
	pbEnrollment.RegisterEnrollmentServiceServer(grpcServer, pbEnrollment.UnimplementedEnrollmentServiceServer{})

	//register gRPC server from service
	pbAssessments.RegisterAssessmentsServiceServer(grpcServer, pbAssessments.UnimplementedAssessmentsServiceServer{})

	pbCertificate.RegisterCertificateServiceServer(grpcServer, pbCertificate.UnimplementedCertificateServiceServer{})

	pbLearningProgress.RegisterLearningProgressServiceServer(grpcServer, pbLearningProgress.UnimplementedLearningProgressServiceServer{})

	pbPayments.RegisterPaymentsServiceServer(grpcServer, pbPayments.UnimplementedPaymentsServiceServer{})

	// Start gRPC server in a goroutine
	go func() {
		logrus.Println("gRPC server is running on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
}

func runGrpcGatewayServer() {
	// Set up gRPC-Gateway mux (HTTP to gRPC)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register HTTP handlers for gRPC-Gateway
	err := pbEnrollment.RegisterEnrollmentServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		logrus.Fatalf("Failed to register gRPC Gateway for Enrollment service: %v", err)
	}

	err = pbAssessments.RegisterAssessmentsServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		logrus.Fatalf("Failed to register gRPC Gateway for Assessments service: %v", err)
	}

	err = pbCertificate.RegisterCertificateServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		logrus.Fatalf("Failed to register gRPC Gateway for Certificate service: %v", err)
	}

	err = pbLearningProgress.RegisterLearningProgressServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		logrus.Fatalf("Failed to register gRPC Gateway for LearningProgress service: %v", err)
	}

	err = pbPayments.RegisterPaymentsServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		logrus.Fatalf("Failed to register gRPC Gateway for Payments service: %v", err)
	}

	// Start gRPC-Gateway server on a separate port (e.g., 8081)
	logrus.Println("gRPC-Gateway server is running on port 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		logrus.Fatalf("Failed to serve gRPC-Gateway server: %v", err)
	}
}
