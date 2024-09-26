package service

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/assessments"
	pb "github.com/ghssni/Smartcy-LMS/Enrollment-Service/proto/assessments"
	"net/http"
)

type AssessmentsService struct {
	pb.UnimplementedAssessmentsServiceServer
	assessmentsRepo repository.AssessmentsRepository
}

func NewAssessmentsService(assessmentsRepo repository.AssessmentsRepository) *AssessmentsService {
	return &AssessmentsService{
		assessmentsRepo: assessmentsRepo,
	}
}

func (s *AssessmentsService) CreateAssessment(ctx context.Context, req *assessments.CreateAssessmentRequest) (*assessments.CreateAssessmentResponse, error) {
	assessment, err := s.assessmentsRepo.CreateAssessment(req.EnrollmentId, req.Score, req.AssessmentType, req.TakenAt.AsTime(), req.CreatedAt.AsTime(), req.UpdatedAt.AsTime())
	if err != nil {
		return nil, err
	}

	// response to client
	response := &assessments.CreateAssessmentResponse{
		Meta: &pb.MetaAssessment{
			Code:    uint32(http.StatusCreated),
			Message: "Assessment created successfully",
			Status:  http.StatusText(http.StatusCreated),
		},
		Assessments: assessment,
	}
	return response, nil
}

func (s *AssessmentsService) GetAssessmentByStudentId(ctx context.Context, req *assessments.GetAssessmentByStudentIdRequest) (*assessments.GetAssessmentByStudentIdResponse, error) {
	assessment, err := s.assessmentsRepo.GetAssessmentByStudentId(req.Id, req.EnrollmentId)
	if err != nil {
		return nil, err
	}

	// response to client
	response := &assessments.GetAssessmentByStudentIdResponse{
		Meta: &pb.MetaAssessment{
			Code:    uint32(http.StatusOK),
			Message: "Assessment found",
			Status:  http.StatusText(http.StatusOK),
		},
		Assessments: assessment,
	}
	return response, nil
}
