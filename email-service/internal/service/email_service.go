package service

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/Email-Service/internal/models"
	"github.com/ghssni/Smartcy-LMS/Email-Service/internal/repository"
	pb "github.com/ghssni/Smartcy-LMS/Email-Service/pb/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	_, err := s.emailRepo.InsertEmail(email)
	if err != nil {
		return &pb.SendPaymentDueEmailResponse{
			Meta: &pb.MetaEmail{
				Code:    int32(codes.Internal),
				Message: "Failed to insert email",
				Status:  http.StatusText(http.StatusInternalServerError),
			},
			Success: false,
		}, status.Errorf(codes.Internal, "failed to insert email: %v", err)
	}

	// send email
	err = SendEmailPayment(req.Email, req.CourseName, req.PaymentLink)
	statusStr := "sent"
	errorMsg := ""
	if err != nil {
		logrus.Println("Error sending email:", err)
		statusStr = "failed"
		errorMsg = err.Error()
		return &pb.SendPaymentDueEmailResponse{
			Meta: &pb.MetaEmail{
				Code:    int32(codes.Internal),
				Message: "Failed to send email",
				Status:  http.StatusText(http.StatusInternalServerError),
			},
			Success: false,
		}, status.Errorf(codes.Internal, "Error sending email: %v", err)
	}

	// log email
	emailLog := &models.EmailLogs{
		UserID:       req.UserId,
		Email:        req.Email,
		Status:       statusStr,
		SentAt:       time.Now(),
		ErrorMessage: errorMsg,
	}

	_, err = s.emailLogRepo.InsertEmailLog(emailLog)
	if err != nil {
		logrus.Println("Error inserting email log:", err)
	} else {
		logrus.Println("Email log inserted successfully for:", req.Email)
	}

	response := &pb.SendPaymentDueEmailResponse{
		Meta: &pb.MetaEmail{
			Code:    int32(codes.OK),
			Message: "Email sent successfully",
			Status:  http.StatusText(http.StatusOK),
		},
		Success: statusStr == "sent",
	}
	logrus.Println("Response Success:", response.Success)
	return response, nil
}

func (s *EmailService) SendForgotPasswordEmail(ctx context.Context, req *pb.SendForgotPasswordEmailRequest) (*pb.SendForgotPasswordEmailResponse, error) {
	email := &models.Email{
		UserID:    req.UserId,
		EmailType: "forgot_password",
		Email:     req.Email,
	}

	_, err := s.emailRepo.InsertEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert email: %v", err)
	}

	// send email
	err = SendEmailForgotPassword(req.Email, req.ResetLink, req.ResetToken)
	statusStr := "sent"
	errorMsg := ""
	if err != nil {
		logrus.Println("Error sending email:", err)
		statusStr = "failed"
		errorMsg = err.Error()
	}

	// log email
	emailLog := &models.EmailLogs{
		UserID:       req.UserId,
		Email:        req.Email,
		Status:       statusStr,
		SentAt:       time.Now(),
		ErrorMessage: errorMsg,
	}

	_, err = s.emailLogRepo.InsertEmailLog(emailLog)

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
		CreatedAt: time.Now(),
	}

	_, err := s.emailRepo.InsertEmail(email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert email: %v", err)
	}

	// send email
	err = SendEmailSuccess(req.Email, req.CourseName)
	statusStr := "sent"
	errorMsg := ""
	if err != nil {
		logrus.Println("Error sending email:", err)
		statusStr = "failed"
		errorMsg = err.Error()
	}

	// log email
	emailLog := &models.EmailLogs{
		UserID:       req.UserId,
		Email:        req.Email,
		Status:       statusStr,
		SentAt:       time.Now(),
		ErrorMessage: errorMsg,
	}

	_, err = s.emailLogRepo.InsertEmailLog(emailLog)

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
