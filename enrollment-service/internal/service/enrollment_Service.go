package service

import (
	"context"
	"errors"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/models"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"net/http"
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

func (s *enrollmentService) CreateEnrollment(ctx context.Context, req *pb.CreateEnrollmentRequest) (*pb.CreateEnrollmentResponse, error) {
	studentId := req.StudentId

	enrollmentInput := &models.EnrollmentInput{
		CourseID:      req.CourseId,
		StudentID:     studentId,
		PaymentStatus: "Pending",
		EnrolledAt:    time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	tx := s.er.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err := s.er.ExistingEnrollment(studentId, enrollmentInput.CourseID)
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.AlreadyExists, "student is already enrolled in this course")
	}

	newEnrollment := &models.Enrollments{
		StudentID:     enrollmentInput.StudentID,
		CourseID:      enrollmentInput.CourseID,
		EnrolledAt:    enrollmentInput.EnrolledAt,
		PaymentStatus: enrollmentInput.PaymentStatus,
		CreatedAt:     enrollmentInput.CreatedAt,
		UpdatedAt:     enrollmentInput.UpdatedAt,
	}

	if err := s.er.CreateEnrollment(newEnrollment); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "failed to create enrollment: %v", err)
	}

	// Create payment
	paymentInput := &models.Payments{
		EnrollmentID: newEnrollment.ID,
		Status:       "Pending",
		Created:      time.Now(),
		Updated:      time.Now(),
	}

	if err := s.payment.CreatePayment(ctx, paymentInput); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "failed to create payment: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

	response := &pb.CreateEnrollmentResponse{
		Meta: &pb.MetaEnrollment{
			Message: "enrollment created successfully",
			Code:    uint32(http.StatusCreated),
			Status:  http.StatusText(http.StatusCreated),
		},
		Data: &pb.Enrollment{
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

func (s *enrollmentService) GetEnrollmentsByStudentId(ctx context.Context, req *pb.GetEnrollmentsByStudentIdRequest) (*pb.GetEnrollmentsByStudentIdResponse, error) {
	studentId := req.StudentId

	enrollments, err := s.er.GetEnrollmentsByStudentId(studentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "No enrollments found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get enrollments: %v", err)
	}

	enrollmentsResponse := make([]*pb.Enrollment, 0, len(enrollments))

	for _, enrollmentRequest := range enrollments {
		enrollmentsResponse = append(enrollmentsResponse, &pb.Enrollment{
			Id:            enrollmentRequest.ID,
			StudentId:     enrollmentRequest.StudentID,
			CourseId:      enrollmentRequest.CourseID,
			PaymentStatus: enrollmentRequest.PaymentStatus,
			EnrolledAt:    timestamppb.New(enrollmentRequest.EnrolledAt),
			CreatedAt:     timestamppb.New(enrollmentRequest.CreatedAt),
			UpdatedAt:     timestamppb.New(enrollmentRequest.UpdatedAt),
		})
	}

	response := &pb.GetEnrollmentsByStudentIdResponse{
		Meta: &pb.MetaEnrollment{
			Message: "enrollments retrieved successfully",
			Code:    uint32(codes.OK),
			Status:  codes.OK.String(),
		},
		Data: enrollmentsResponse,
	}

	return response, nil
}

func (s *enrollmentService) DeleteEnrollmentById(ctx context.Context, req *pb.DeleteEnrollmentByIdRequest) (*pb.DeleteEnrollmentByIdResponse, error) {
	enrollment, err := s.er.GetEnrollmentsById(req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "Enrollment not found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get enrollment: %v", err)
	}

	if err := s.er.DeleteEnrollmentById(enrollment); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete enrollment: %v", err)
	}

	response := &pb.DeleteEnrollmentByIdResponse{
		Meta: &pb.MetaEnrollment{
			Message: "Enrollment deleted successfully",
			Code:    uint32(http.StatusOK),
			Status:  http.StatusText(http.StatusOK),
		},
	}

	return response, nil
}

func NewEnrollmentService(er repository.EnrollmentRepository, payment repository.PaymentRepository) EnrollmentService {
	return &enrollmentService{
		er:      er,
		payment: payment,
	}
}
