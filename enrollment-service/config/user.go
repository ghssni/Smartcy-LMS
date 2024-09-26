package config

import (
	"context"
	"fmt"
	"github.com/ghssni/Smartcy-LMS/User-Service/pb"
	"google.golang.org/grpc"
)

func GetUserFromEmail(email string) (string, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return "", fmt.Errorf("failed to connect to user service: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	resp, err := client.GetUserByEmail(context.Background(), &pb.GetUserByEmailRequest{Email: email})
	if err != nil {
		return "", fmt.Errorf("failed to get user from user service: %v", err)
	}
	return resp.User.Email, nil
}

func LogUserActivity(userId string, activity string) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserActivityLogServiceClient(conn)
	_, err = client.CreateUserActivityLog(context.Background(), &pb.CreateUserActivityLogRequest{UserId: userId, ActivityType: activity})
	if err != nil {
		return fmt.Errorf("failed to log user activity: %v", err)
	}
	return nil
}
