package service

import (
	"context"
	"course-service/data"
	"course-service/pb"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type LessonService struct{}

func NewLessonService() *LessonService { return &LessonService{} }

func (s *LessonService) CreateLesson(ctx context.Context, in *pb.CreateLessonRequest) (*pb.Lesson, error) {
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
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to create lesson: %v", err))
	}

	res := &pb.Lesson{
		Id:         lessonId,
		CourseId:   courseID,
		Title:      title,
		ContentUrl: contentUrl,
		LessonType: lessonType,
		Sequence:   sequence,
		CreatedAt:  createdAt.String(),
		UpdatedAt:  updatedAt.String(),
	}

	return res, nil
}

func (s *LessonService) GetLesson(ctx context.Context, in *pb.GetLessonRequest) (*pb.Lesson, error) {
	// Get data from request
	lessonID := in.GetId()
	courseID := in.GetCourseId()

	// Get lesson from database
	lesson, err := repo.Lesson.GetLesson(ctx, lessonID, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get lesson: %v", err))
	}

	// Return response
	res := &pb.Lesson{
		Id:         lesson.ID,
		CourseId:   courseID,
		Title:      lesson.Title,
		ContentUrl: lesson.ContentURL,
		LessonType: lesson.LessonType,
		Sequence:   lesson.Sequence,
		CreatedAt:  lesson.CreatedAt.String(),
		UpdatedAt:  lesson.UpdatedAt.String(),
	}

	return res, nil
}

func (s *LessonService) GetLessonBySequence(ctx context.Context, in *pb.GetLessonBySequenceRequest) (*pb.Lesson, error) {
	// Get data from request
	sequence := in.GetSequence()
	courseID := in.GetCourseId()

	// Get lesson from database
	lesson, err := repo.Lesson.GetLessonBySequence(ctx, sequence, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get lesson: %v", err))
	}

	// Return response
	res := &pb.Lesson{
		Id:         lesson.ID,
		CourseId:   courseID,
		Title:      lesson.Title,
		ContentUrl: lesson.ContentURL,
		LessonType: lesson.LessonType,
		Sequence:   lesson.Sequence,
		CreatedAt:  lesson.CreatedAt.String(),
		UpdatedAt:  lesson.UpdatedAt.String(),
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

func (s *LessonService) UpdateLesson(ctx context.Context, in *pb.UpdateLessonRequest) (*emptypb.Empty, error) {
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

	return &emptypb.Empty{}, nil
}

func (s *LessonService) DeleteLesson(ctx context.Context, in *pb.DeleteLessonRequest) (*emptypb.Empty, error) {
	// Get data from request
	lessonID := in.GetId()
	courseID := in.GetCourseId()

	// DeletedAt
	deletedAt := time.Now()

	// Delete lesson from database
	err := repo.Lesson.DeleteLesson(ctx, lessonID, courseID, deletedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to delete lesson: %v", err))
	}

	// Return response
	return &emptypb.Empty{}, nil
}

func (s *LessonService) DeleteLessonByCourseID(ctx context.Context, in *pb.DeleteLessonByCourseIDRequest) (*emptypb.Empty, error) {
	// Get data from request
	courseID := in.GetCourseId()

	// DeletedAt
	deletedAt := time.Now()

	// Delete lesson from database
	err := repo.Lesson.DeleteLessonByCourse(ctx, courseID, deletedAt)
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
