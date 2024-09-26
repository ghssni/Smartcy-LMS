package service

import (
	"context"
	"course-service/data"
	"course-service/pb"
	"course-service/utils"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type LessonService struct{}

func NewLessonService() *LessonService { return &LessonService{} }

func (s *LessonService) CreateLesson(ctx context.Context, in *pb.CreateLessonRequest) (*pb.CreateLessonResponse, error) {
	// Get data from request
	courseID := in.GetCourseId()
	title := in.GetTitle()
	contentUrl := in.GetContentUrl()
	lessonType := in.GetLessonType()
	sequence := in.GetSequence()

	// Insert data to database
	lesson := &data.Lesson{
		CourseID:   courseID,
		Title:      title,
		ContentURL: contentUrl,
		LessonType: lessonType,
		Sequence:   sequence,
	}

	// CreateAt and UpdatedAt
	createdAt, updatedAt := time.Now(), time.Now()

	// Insert data to database
	lessonId, err := repo.Lesson.CreateLesson(ctx, lesson, createdAt, updatedAt)
	if err != nil {
		return nil, utils.HandlePostgresError(err)
	}

	res := &pb.CreateLessonResponse{
		Id: lessonId,
	}

	return res, nil
}

func (s *LessonService) GetLesson(ctx context.Context, in *pb.GetLessonRequest) (*pb.GetLessonResponse, error) {
	// Get data from request
	lessonID := in.GetId()

	// Get lesson from database
	lesson, err := repo.Lesson.GetLesson(ctx, lessonID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get lesson: %v", err))
	}

	// Return response
	res := &pb.GetLessonResponse{
		Lesson: &pb.Lesson{
			Id:         lesson.ID,
			CourseId:   lesson.CourseID,
			Title:      lesson.Title,
			ContentUrl: lesson.ContentURL,
			LessonType: lesson.LessonType,
			Sequence:   lesson.Sequence,
			CreatedAt:  lesson.CreatedAt.Format("02-01-2006"),
			UpdatedAt:  lesson.UpdatedAt.Format("02-01-2006"),
		},
	}

	return res, nil
}

func (s *LessonService) GetLessonBySequence(ctx context.Context, in *pb.GetLessonBySequenceRequest) (*pb.GetLessonBySequenceResponse, error) {
	// Get data from request
	sequence := in.GetSequence()
	courseID := in.GetCourseId()

	// Get lesson from database
	lesson, err := repo.Lesson.GetLessonBySequence(ctx, sequence, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get lesson: %v", err))
	}

	// Return response
	res := &pb.GetLessonBySequenceResponse{
		Lesson: &pb.Lesson{
			Id:         lesson.ID,
			CourseId:   courseID,
			Title:      lesson.Title,
			ContentUrl: lesson.ContentURL,
			LessonType: lesson.LessonType,
			Sequence:   lesson.Sequence,
			CreatedAt:  lesson.CreatedAt.Format("02-01-2006"),
			UpdatedAt:  lesson.UpdatedAt.Format("02-01-2006"),
		},
	}

	return res, nil
}

func (s *LessonService) ListLessons(ctx context.Context, in *pb.ListLessonsRequest) (*pb.ListLessonsResponse, error) {
	// Get data from request
	courseID := in.GetCourseId()

	// Get all lessons from database
	lessons, err := repo.Lesson.GetAllLessons(ctx, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get all lessons: %v", err))
	}

	var lessonsRes []*pb.LessonSummary

	for _, lesson := range lessons {
		lessonsRes = append(lessonsRes, &pb.LessonSummary{
			Id:         lesson.ID,
			Title:      lesson.Title,
			LessonType: lesson.LessonType,
			Sequence:   lesson.Sequence,
		})
	}

	res := &pb.ListLessonsResponse{
		Lessons: lessonsRes,
	}

	return res, nil
}

func (s *LessonService) UpdateLesson(ctx context.Context, in *pb.UpdateLessonRequest) (*pb.UpdateLessonResponse, error) {
	// Get data from request
	lessonID := in.GetId()
	courseID := in.GetCourseId()
	title := in.GetTitle()
	contentUrl := in.GetContentUrl()
	lessonType := in.GetLessonType()
	sequence := in.GetSequence()

	// Update data to database
	lesson := &data.Lesson{
		ID:         lessonID,
		CourseID:   courseID,
		Title:      title,
		ContentURL: contentUrl,
		LessonType: lessonType,
		Sequence:   sequence,
	}

	// UpdatedAt
	updatedAt := time.Now()

	// Update data to database
	err := repo.Lesson.UpdateLesson(ctx, lesson, updatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to update lesson: %v", err))
	}

	// Return response
	res := &pb.UpdateLessonResponse{
		Lesson: &pb.Lesson{
			Id:         lessonID,
			CourseId:   courseID,
			Title:      title,
			ContentUrl: contentUrl,
			LessonType: lessonType,
			Sequence:   sequence,
			CreatedAt:  lesson.CreatedAt.Format("02-01-2006"),
			UpdatedAt:  updatedAt.Format("02-01-2006"),
		},
	}

	return res, nil
}

func (s *LessonService) DeleteLesson(ctx context.Context, in *pb.DeleteLessonRequest) (*emptypb.Empty, error) {
	// Get data from request
	lessonID := in.GetId()

	// DeletedAt
	deletedAt := time.Now()

	// Delete lesson from database
	err := repo.Lesson.DeleteLesson(ctx, lessonID, deletedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to delete lesson: %v", err))
	}

	// Return response
	return &emptypb.Empty{}, nil
}

func (s *LessonService) SearchLessonsByTitle(ctx context.Context, in *pb.SearchLessonsByTitleRequest) (*pb.ListLessonsResponse, error) {
	// Get data from request
	courseID := in.GetCourseId()
	title := in.GetTitle()

	// Search lesson from database
	lessons, err := repo.Lesson.SearchLessonByTitle(ctx, courseID, title)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to search lesson: %v", err))
	}

	// Return response
	var lessonsRes []*pb.LessonSummary

	for _, lesson := range lessons {
		lessonsRes = append(lessonsRes, &pb.LessonSummary{
			Id:         lesson.ID,
			Title:      lesson.Title,
			LessonType: lesson.LessonType,
			Sequence:   lesson.Sequence,
		})
	}

	res := &pb.ListLessonsResponse{
		Lessons: lessonsRes,
	}

	return res, nil
}

func (s *LessonService) SearchLessonsByType(ctx context.Context, in *pb.SearchLessonsByTypeRequest) (*pb.ListLessonsResponse, error) {
	// Get data from request
	courseID := in.GetCourseId()
	lessonType := in.GetLessonType()

	// Search lesson from database
	lessons, err := repo.Lesson.SearchLessonByType(ctx, courseID, lessonType)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to search lesson: %v", err))
	}

	// Return response
	var lessonsRes []*pb.LessonSummary

	for _, lesson := range lessons {
		lessonsRes = append(lessonsRes, &pb.LessonSummary{
			Id:         lesson.ID,
			Title:      lesson.Title,
			LessonType: lesson.LessonType,
			Sequence:   lesson.Sequence,
		})
	}

	res := &pb.ListLessonsResponse{
		Lessons: lessonsRes,
	}

	return res, nil
}

func (s *LessonService) GetTotalLessons(ctx context.Context, in *pb.GetTotalLessonsRequest) (*pb.GetTotalLessonsResponse, error) {
	// Get data from request
	courseID := in.GetCourseId()

	// Get total lessons from database
	total, err := repo.Lesson.GetTotalLessonsByCourseID(ctx, courseID)

	// Return response
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to get total lessons: %v", err))
	}

	// Return response
	res := &pb.GetTotalLessonsResponse{
		Total: total,
	}

	return res, nil
}

func (s *LessonService) GetTotalLessonsByType(ctx context.Context, in *pb.GetTotalLessonsByTypeRequest) (*pb.GetTotalLessonsResponse, error) {
	// Get data from request
	courseID := in.GetCourseId()
	lessonType := in.GetLessonType()

	// Get lesson from database
	total, err := repo.Lesson.GetTotalLessonsByCourseIDAndType(ctx, courseID, lessonType)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to get total lessons: %v", err))
	}

	// Return response
	res := &pb.GetTotalLessonsResponse{
		Total: total,
	}

	return res, nil
}
