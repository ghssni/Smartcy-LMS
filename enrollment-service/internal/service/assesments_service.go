package service

import (
	"context"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/internal/repository"
	"github.com/ghssni/Smartcy-LMS/Enrollment-Service/pb"
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

func (s *AssessmentsService) CreateAssessment(ctx context.Context, req *pb.CreateAssessmentRequest) (*pb.CreateAssessmentResponse, error) {
	panic("implement me")

}

func (s *AssessmentsService) GetAssessmentByStudentId(ctx context.Context, req *pb.GetAssessmentByStudentIdRequest) (*pb.GetAssessmentByStudentIdResponse, error) {
	panic("implement me")
}
