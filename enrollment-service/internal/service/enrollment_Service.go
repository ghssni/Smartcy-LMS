package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/middleware"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/models"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/proto/enrollment"
	pb "github.com/ghssni/Smartcy-LMS/enrollment-service/proto/enrollment"
	"github.com/ghssni/Smartcy-LMS/enrollment-service/proto/meta"
	"github.com/ghssni/Smartcy-LMS/pkg"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type EnrollmentService interface {
	pb.EnrollmentServiceServer
}

type enrollmentService struct {
	pb.UnimplementedEnrollmentServiceServer
	er      repository.EnrollmentRepository
	payment repository.PaymentRepository
}

func (s *enrollmentService) CreateEnrollment(ctx context.Context, req *enrollment.CreateEnrollmentRequest) (*enrollment.CreateEnrollmentResponse, error) {
	studentId, err := middleware.GetUserIDFromToken(ctx)
	email, err := middleware.GetEmailFromToken(ctx)
	if err != nil {
		return nil, err
	}

	enrollmentInput := &models.EnrollmentInput{
		CourseID:      req.CourseId,
		StudentID:     studentId,
		PaymentStatus: "Pending",
		EnrolledAt:    time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// validate enrollment input
	if err := enrollmentInput.Validate(); err != nil {
		var validationError validator.ValidationErrors
		if errors.As(err, &validationError) {
			formatterError := pkg.FormatValidationError(enrollmentInput, validationError)
			if formatterError != "" {
				return nil, status.Errorf(codes.InvalidArgument, formatterError)
			}
		}
	}
	tx := s.er.BeginTransaction()

	// check if student is already enrolled example courseID = 1
	_, err = s.er.ExistingEnrollment(studentId, req.CourseId)
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.AlreadyExists, "student is already enrolled in this course")
	}

	// check if student is already enrolled example courseID = 1
	newEnrollment := &models.Enrollments{
		StudentID:     enrollmentInput.StudentID,
		CourseID:      enrollmentInput.CourseID,
		EnrolledAt:    enrollmentInput.EnrolledAt,
		PaymentStatus: enrollmentInput.PaymentStatus,
		CreatedAt:     enrollmentInput.CreatedAt,
		UpdatedAt:     enrollmentInput.UpdatedAt,
	}

	// create enrollment
	if err := s.er.CreateEnrollment(newEnrollment); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "failed to create enrollment: %v", err)
	}

	invoiceUrl, invoiceId, price, err := CreateInvoiceAndSendEmailPayment(studentId, email, "example course", 100000)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create invoice and send email: %v", err)
	}
	description := fmt.Sprintf("Payment for %s with amount: IDR %.2f", "example course", price)
	payment := &models.Payments{
		EnrollmentID: newEnrollment.ID,
		ExternalID:   invoiceId,
		UserID:       studentId,
		IsHigh:       false,
		Status:       "PENDING",
		Amount:       price,
		PaidAmount:   0,
		PayerEmail:   email,
		Description:  description,
		Updated:      time.Now(),
		InvoiceUrl:   invoiceUrl,
	}

	// create payment
	if err := s.payment.CreatePayment(ctx, payment); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "failed to create payment: %v", err)
	}
	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}
	// log user activity
	err = s.logUserActivity(ctx, studentId, "enrollment Course "+"example course")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to log user activity")
	}

	response := &enrollment.CreateEnrollmentResponse{
		Meta: &meta.Meta{
			Message: "enrollment created successfully",
			Code:    int32(http.StatusCreated),
			Status:  http.StatusText(http.StatusCreated),
		},
		Data: &enrollment.Enrollment{
			Id:            newEnrollment.ID,
			StudentId:     newEnrollment.StudentID,
			CourseId:      newEnrollment.CourseID,
			PaymentStatus: newEnrollment.PaymentStatus,
			EnrolledAt:    timestamppb.New(newEnrollment.EnrolledAt),
			CreatedAt:     timestamppb.New(newEnrollment.CreatedAt),
			UpdatedAt:     timestamppb.New(newEnrollment.UpdatedAt),
		},
	}

	return response, nil
}

func (s *enrollmentService) GetEnrollmentsByStudentId(ctx context.Context, req *pb.GetEnrollmentsByStudentIdRequest) (*enrollment.GetEnrollmentsByStudentIdResponse, error) {
	studentId, err := middleware.GetUserIDFromToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user id from token: %v", err)
	}

	enrollments, err := s.er.GetEnrollmentsByStudentId(studentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "No enrollments found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get enrollments: %v", err)
	}

	err = s.logUserActivity(ctx, studentId, "get enrollments")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to log user activity")
	}

	enrollmentsResponse := make([]*enrollment.Enrollment, 0, len(enrollments))

	for _, enrollmentRequest := range enrollments {
		enrollmentsResponse = append(enrollmentsResponse, &enrollment.Enrollment{
			Id:            enrollmentRequest.ID,
			StudentId:     enrollmentRequest.StudentID,
			CourseId:      enrollmentRequest.CourseID,
			PaymentStatus: enrollmentRequest.PaymentStatus,
			EnrolledAt:    timestamppb.New(enrollmentRequest.EnrolledAt),
			CreatedAt:     timestamppb.New(enrollmentRequest.CreatedAt),
			UpdatedAt:     timestamppb.New(enrollmentRequest.UpdatedAt),
		})
	}

	response := &enrollment.GetEnrollmentsByStudentIdResponse{
		Meta: &meta.Meta{
			Message: "enrollments retrieved successfully",
			Code:    int32(codes.OK),
			Status:  codes.OK.String(),
		},
		Data: enrollmentsResponse,
	}

	return response, nil
}

func (s *enrollmentService) DeleteEnrollmentById(ctx context.Context, req *enrollment.DeleteEnrollmentByIdRequest) (*enrollment.DeleteEnrollmentByIdResponse, error) {
	studentId, err := middleware.GetUserIDFromToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user id from token: %v", err)
	}

	enrollmentRequest, err := s.er.ExistingEnrollment(studentId, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get enrollment: %v", err)
	}

	if enrollmentRequest == nil {
		return nil, status.Error(codes.NotFound, "enrollment not found")
	}

	if err := s.er.DeleteEnrollmentById(enrollmentRequest); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete enrollment: %v", err)
	}

	err = s.logUserActivity(ctx, studentId, "delete enrollment ID : "+fmt.Sprint(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to log user activity")
	}

	response := &enrollment.DeleteEnrollmentByIdResponse{
		Meta: &meta.Meta{
			Message: "enrollment deleted successfully",
			Code:    int32(http.StatusOK),
			Status:  http.StatusText(http.StatusOK),
		},
	}

	return response, nil
}

func (s *enrollmentService) logUserActivity(ctx context.Context, studentId string, activityType string) error {
	logRequest := map[string]interface{}{
		"user_id":       studentId,
		"course_id":     "",
		"activity_type": activityType,
	}

	// Log user activity
	jsonData, err := json.Marshal(logRequest)
	if err != nil {
		return err
	}

	// URL user-service
	url := os.Getenv("API_USER_SERVICE_URL")
	if url == "" {
		return fmt.Errorf("API_USER_SERVICE_URL is not set")
	}

	// Get JWT token from context or metadata
	token, err := middleware.GetTokenFromContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to get JWT token from context: %v", err)
	}

	// HTTP POST request to log user activity
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to log user activity: %s", resp.Status)
	}
	return nil

}

func NewEnrollmentService(er repository.EnrollmentRepository, payment repository.PaymentRepository) EnrollmentService {
	return &enrollmentService{
		er:      er,
		payment: payment,
	}
}
