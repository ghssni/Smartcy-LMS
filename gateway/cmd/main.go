package main

import (
	"fmt"
	"gateway-service/config"
	"gateway-service/pb"
	"gateway-service/server"
	"gateway-service/server/handler"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"log"
)

func main() {
	config.InitViper()
	config.InitValidator()

	courseServiceAddress := config.Viper.GetString("COURSE_SERVICE_ADDRESS")
	userServiceAddress := config.Viper.GetString("USER_SERVICE_ADDRESS")
	emailServiceAddress := config.Viper.GetString("EMAIL_SERVICE_ADDRESS")
	enrollmentServiceAddress := config.Viper.GetString("ENROLLMENT_SERVICE_ADDRESS")

	courseServiceDial, err := grpc.Dial(courseServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial course service: %v", err)
	}

	userServiceDial, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial user service: %v", err)
	}

	emailServiceDial, err := grpc.Dial(emailServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial email service: %v", err)
	}

	enrollmentServiceDial, err := grpc.Dial(enrollmentServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial enrollment service: %v", err)
	}

	// Initialize the client
	userServiceClient := pb.NewUserServiceClient(userServiceDial)
	courseServiceClient := pb.NewCourseServiceClient(courseServiceDial)
	reviewServiceClient := pb.NewReviewServiceClient(courseServiceDial)
	learningProgressServiceClient := pb.NewLearningProgressServiceClient(courseServiceDial)
	lessonServiceClient := pb.NewLessonServiceClient(courseServiceDial)
	emailServiceClient := pb.NewEmailServiceClient(emailServiceDial)
	enrollmentServiceClient := pb.NewEnrollmentServiceClient(enrollmentServiceDial)
	paymentsServiceClient := pb.NewPaymentsServiceClient(enrollmentServiceDial)

	// Initialize the handler
	userHandler := handler.NewUserHandler(userServiceClient, emailServiceClient)
	courseHandler := handler.NewCourseHandler(courseServiceClient, lessonServiceClient)
	reviewHandler := handler.NewReviewHandler(reviewServiceClient)
	lessonHandler := handler.NewLessonHandler(lessonServiceClient)
	learningProgressHandler := handler.NewLearningProgressHandler(learningProgressServiceClient)

	// Initialize the handler enrollment
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentServiceClient, userServiceClient, courseServiceClient)
	paymentsHandler := handler.NewPaymentsHandler(paymentsServiceClient, userServiceClient, emailServiceClient)

	e := echo.New()

	handlers := server.NewHandlers(userHandler, courseHandler, reviewHandler, lessonHandler,  learningProgressHandler, paymentsHandler, enrollmentHandler)

	server.Routes(e, handlers)

	//env := config.Viper.GetString("APP_ENV")
	port := "8080"
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))

}
