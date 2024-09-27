package service

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/Email-Service/internal/models"
	"github.com/ghssni/Smartcy-LMS/Email-Service/internal/repository"
	pb "github.com/ghssni/Smartcy-LMS/Email-Service/pb/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net/http"
	"time"
)

type EmailService struct {
	pb.UnimplementedEmailServiceServer
	emailRepo    repository.EmailsRepository
	emailLogRepo repository.EmailsLogRepository
}

func NewEmailService(emailRepo repository.EmailsRepository, emailLogRepo repository.EmailsLogRepository) *EmailService {
	return &EmailService{
		emailRepo:    emailRepo,
		emailLogRepo: emailLogRepo,
	}
}

func (s *EmailService) SendPaymentDueEmail(ctx context.Context, req *pb.SendPaymentDueEmailRequest) (*pb.SendPaymentDueEmailResponse, error) {
	email := &models.Email{
		UserID:    req.UserId,
		EmailType: "payment_due",
		Email:     req.Email,
	}

	emailLog := &models.EmailLogs{
		UserID: req.UserId,
		Email:  req.Email,
	}

	// Refactored to use the helper function
	statusStr, _ := s.sendAndLogEmail(email, emailLog, func() error {
		return SendEmailPayment(req.Email, req.CourseName, req.PaymentLink)
	})

	response := &pb.SendPaymentDueEmailResponse{
		Meta: &pb.MetaEmail{
			Code:    int32(codes.OK),
			Message: "Email sent successfully",
			Status:  http.StatusText(http.StatusOK),
		},
		Success: statusStr == "sent",
	}
	return response, nil
}

func (s *EmailService) SendForgotPasswordEmail(ctx context.Context, req *pb.SendForgotPasswordEmailRequest) (*pb.SendForgotPasswordEmailResponse, error) {
	email := &models.Email{
		UserID:    req.UserId,
		EmailType: "forgot_password",
		Email:     req.Email,
	}

	emailLog := &models.EmailLogs{
		UserID: req.UserId,
		Email:  req.Email,
	}

	statusStr, _ := s.sendAndLogEmail(email, emailLog, func() error {
		return SendEmailForgotPassword(req.Email, req.ResetLink, req.ResetToken)
	})

	response := &pb.SendForgotPasswordEmailResponse{
		Meta: &pb.MetaEmail{
			Code:    int32(codes.OK),
			Message: "Email sent successfully",
			Status:  http.StatusText(http.StatusOK),
		},
		Success: statusStr == "sent",
	}
	return response, nil
}

func (s *EmailService) SendPaymentSuccessEmail(ctx context.Context, req *pb.SendPaymentSuccessEmailRequest) (*pb.SendPaymentSuccessEmailResponse, error) {
	email := &models.Email{
		UserID:    req.UserId,
		EmailType: "payment_success",
		Email:     req.Email,
	}

	emailLog := &models.EmailLogs{
		UserID: req.UserId,
		Email:  req.Email,
	}

	// Refactored to use the helper function
	statusStr, _ := s.sendAndLogEmail(email, emailLog, func() error {
		return SendEmailSuccess(req.Email, req.CourseName)
	})

	response := &pb.SendPaymentSuccessEmailResponse{
		Meta: &pb.MetaEmail{
			Code:    int32(codes.OK),
			Message: "Email sent successfully",
			Status:  http.StatusText(http.StatusOK),
		},
		Success: statusStr == "sent",
	}
	return response, nil
}

func (s *EmailService) logEmail(emailLog *models.EmailLogs) {
	_, err := s.emailLogRepo.InsertEmailLog(emailLog)
	if err != nil {
		logrus.Println("Error inserting email log:", err)
	} else {
		logrus.Println("Email log inserted successfully for:", emailLog.Email)
	}
}

func (s *EmailService) sendAndLogEmail(email *models.Email, emailLog *models.EmailLogs, sendEmailFunc func() error) (string, string) {
	// Insert email record
	_, err := s.emailRepo.InsertEmail(email)
	if err != nil {
		logrus.Println("Failed to insert email:", err)
		return "failed", err.Error()
	}

	// Send email
	err = sendEmailFunc()
	statusStr := "sent"
	errorMsg := ""
	if err != nil {
		logrus.Println("Error sending email:", err)
		statusStr = "failed"
		errorMsg = err.Error()
	}

	// Log the email result
	emailLog.Status = statusStr
	emailLog.ErrorMessage = errorMsg
	emailLog.SentAt = time.Now()
	s.logEmail(emailLog)

	return statusStr, errorMsg
}
