package main

import (
	"course-service/config"
	"course-service/data"
	"course-service/pb"
	"course-service/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config.InitViper()
	postgresDb := config.InitDB()

	dbModel := data.New(postgresDb)

	service.InitService(dbModel)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(service.Unary()),
		grpc.StreamInterceptor(service.Stream()),
	)

	courseServer := service.NewCourseService()
	lessonServer := service.NewLessonService()
	pb.RegisterCourseServiceServer(grpcServer, courseServer)
	pb.RegisterLessonServiceServer(grpcServer, lessonServer)

	gRPCPort := config.Viper.GetString("PORT")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on port :", gRPCPort)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}

}
