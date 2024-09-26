package config

import (
	"context"
	"fmt"
	pb "github.com/ghssni/Smartcy-LMS/Email-Service/pb/proto"
	"google.golang.org/grpc"
)

func SendPaymentDueEmail(email, courseName, invoiceURL string) error {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewEmailServiceClient(conn)
	_, err = client.SendPaymentDueEmail(context.Background(), &pb.SendPaymentDueEmailRequest{
		Email:       email,
		CourseName:  courseName,
		PaymentLink: invoiceURL,
	})
	if err != nil {
		return fmt.Errorf("failed to send payment due email: %v", err)
	}
	return nil
}

func SendPaymentSuccessEmail(email, description, userId, invoiceId string, amount float32) error {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewEmailServiceClient(conn)
	_, err = client.SendPaymentSuccessEmail(context.Background(), &pb.SendPaymentSuccessEmailRequest{
		Email:      email,
		CourseName: description,
		InvoiceId:  invoiceId,
		Amount:     amount,
		UserId:     userId,
	})
	if err != nil {
		return fmt.Errorf("failed to send payment success email: %v", err)
	}

	return nil
}
