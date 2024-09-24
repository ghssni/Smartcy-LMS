package service

import (
	"context"
	"course-service/pb"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type LearningProgressService struct{}

func NewLearningProgressService() *LearningProgressService { return &LearningProgressService{} }

func (s *LearningProgressService) MarkLessonAsCompleted(ctx context.Context, in *pb.MarkLessonAsCompletedRequest) (*pb.MarkLessonAsCompletedResponse, error) {
	// Get data from request
	enrollmentID := in.GetEnrollmentId()
	lessonID := in.GetLessonId()

	// lastAccessed and completedAt
	lastAccessed, completedAt := time.Now(), time.Now()

	// Insert data to database
	err := repo.LearningProgress.MarkLessonAsCompleted(ctx, enrollmentID, lessonID, lastAccessed, completedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to mark lesson as completed: %v", err))
	}

	// Return response
	res := &pb.MarkLessonAsCompletedResponse{
		Message: "Lesson marked as completed",
	}

	return res, nil
}

func (s *LearningProgressService) ResetLessonMark(ctx context.Context, in *pb.ResetLessonMarkRequest) (*emptypb.Empty, error) {
	// Get data from request
	enrollmentID := in.GetEnrollmentId()
	lessonID := in.GetLessonId()

	// Insert data to database
	err := repo.LearningProgress.ResetLessonMark(ctx, enrollmentID, lessonID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to reset lesson mark: %v", err))
	}

	return &emptypb.Empty{}, nil
}

func (s *LearningProgressService) ResetAllLessonMarks(ctx context.Context, in *pb.ResetAllLessonMarksRequest) (*emptypb.Empty, error) {
	// Get data from request
	enrollmentID := in.GetEnrollmentId()

	// Update last accessed
	err := repo.LearningProgress.ResetAllLessonMarks(ctx, enrollmentID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to reset all lesson marks: %v", err))
	}

	// Return response
	return &emptypb.Empty{}, nil
}

func (s *LearningProgressService) GetTotalCompletedLessons(ctx context.Context, in *pb.GetTotalCompletedLessonsRequest) (*pb.CompletedProgress, error) {
	// Get data from request
	enrollmentID := in.GetEnrollmentId()

	// Get total completed lessons
	totalCompleted, err := repo.LearningProgress.GetTotalCompletedLessons(ctx, enrollmentID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to get total completed lessons: %v", err))
	}

	// Return response
	res := &pb.CompletedProgress{
		EnrollmentId:   totalCompleted.EnrollmentID,
		TotalCompleted: totalCompleted.TotalCompleted,
	}

	return res, nil
}

func (s *LearningProgressService) GetTotalCompletedProgress(ctx context.Context, in *pb.GetTotalCompletedProgressRequest) (*pb.GetTotalCompletedProgressResponse, error) {
	// Get total completed lessons
	totalCompleted, err := repo.LearningProgress.GetTotalCompletedProgress(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to get total completed lessons: %v", err))
	}

	// Return response
	res := &pb.GetTotalCompletedProgressResponse{
		Progress: make([]*pb.CompletedProgress, 0),
	}

	for _, v := range totalCompleted {
		res.Progress = append(res.Progress, &pb.CompletedProgress{
			EnrollmentId:   v.EnrollmentID,
			TotalCompleted: v.TotalCompleted,
		})
	}

	return res, nil
}

func (s *LearningProgressService) ListLearningProgress(ctx context.Context, in *pb.ListLearningProgressRequest) (*pb.ListLearningProgressResponse, error) {
	// Get data from request
	enrollmentID := in.GetEnrollmentId()

	// Get learning progress from database
	learningProgress, err := repo.LearningProgress.ListLearningProgress(ctx, enrollmentID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to list learning progress: %v", err))
	}

	// Return response
	res := &pb.ListLearningProgressResponse{
		LearningProgress: make([]*pb.LearningProgress, len(learningProgress)),
	}

	for i, v := range learningProgress {
		res.LearningProgress[i] = &pb.LearningProgress{
			EnrollmentId: v.EnrollmentID,
			LessonId:     v.LessonID,
			Status:       v.Status,
			LastAccessed: timestamppb.New(*v.LastAccessed),
			CompletedAt:  timestamppb.New(*v.CompletedAt),
		}
	}

	return res, nil
}

func (s *LearningProgressService) UpdateLastAccessed(ctx context.Context, in *pb.UpdateLastAccessedRequest) (*emptypb.Empty, error) {
	// Get data from request
	enrollmentID := in.GetEnrollmentId()
	lessonID := in.GetLessonId()

	// lastAccessed
	lastAccessed := time.Now()

	// Update last accessed
	err := repo.LearningProgress.UpdateLastAccessed(ctx, enrollmentID, lessonID, lastAccessed)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to update last accessed: %v", err))
	}

	// Return response
	return &emptypb.Empty{}, nil
}
