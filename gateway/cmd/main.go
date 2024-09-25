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

	courseServiceDial, err := grpc.Dial(courseServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to dial course service: %v", err)
	}

	courseServiceClient := pb.NewCourseServiceClient(courseServiceDial)
	reviewServiceClient := pb.NewReviewServiceClient(courseServiceDial)
	courseHandler := handler.NewCourseHandler(courseServiceClient, reviewServiceClient)

	e := echo.New()

	handlers := server.NewHandlers(courseHandler)
	server.Routes(e, handlers)

	//env := config.Viper.GetString("APP_ENV")
	port := "8080"
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))

}
