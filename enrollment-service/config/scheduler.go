package config

import (
	"context"
	pb "github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/payments"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"os"
	"time"
)

type SchedulerCronjob struct {
	paymentService pb.PaymentsServiceClient // Menggunakan gRPC client untuk payment service
}

func NewScheduler(paymentService pb.PaymentsServiceClient) *SchedulerCronjob {
	return &SchedulerCronjob{
		paymentService: paymentService,
	}
}

func (s *SchedulerCronjob) Scheduler() error {
	c := cron.New()

	_, err := c.AddFunc("@daily", func() {
		logrus.Println("Running cron job every 1 minute")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		accessKey := os.Getenv("CRON_ACCESS_KEY")
		md := metadata.Pairs("X-ACCESS-KEY", accessKey)
		ctx = metadata.NewOutgoingContext(ctx, md)

		// Update expired payment status
		req := &pb.UpdateExpiredPaymentStatusRequest{}

		_, err := s.paymentService.UpdateExpiredPaymentStatus(ctx, req)
		if err != nil {
			logrus.Errorf("Failed to update expired payments: %v", err)
		} else {
			logrus.Println("Successfully updated expired payments")
		}
	})

	if err != nil {
		return err
	}

	c.Start()

	select {}
}
