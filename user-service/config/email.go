package config

import (
	"context"
	pb "github.com/ghssni/Smartcy-LMS/Email-Service/pb/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func SendEmailForgotPassword(email, resetURL, resetToken string) error {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEmailServiceClient(conn)
	_, err = client.SendForgotPasswordEmail(context.Background(), &pb.SendForgotPasswordEmailRequest{
		Email:      email,
		ResetLink:  resetURL,
		ResetToken: resetToken,
	})
	if err != nil {
		logrus.Fatalf("Could not send email: %v", err)
	}
	return nil
}
