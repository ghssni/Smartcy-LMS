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

type CourseService struct{}

func NewCourseService() *CourseService { return &CourseService{} }

func (s *CourseService) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	// Get data from request
	courseTitle := in.GetTitle()
	courseDesc := in.GetDescription()
	coursePrice := in.GetPrice()
	courseThumbnail := in.GetThumbnailUrl()
	instructorID := in.GetInstructorId()
	category := in.GetCategory()

	// Insert data to database
	course := &data.Course{
		Title:        courseTitle,
		Description:  courseDesc,
		Price:        coursePrice,
		ThumbnailURL: courseThumbnail,
		InstructorID: instructorID,
		Category:     category,
	}

	// createdAt and updatedAt
	createdAt, updatedAt := time.Now(), time.Now()

	courseID, err := repo.Course.CreateCourse(ctx, course, createdAt, updatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to create course: %v", err))
	}

	// Return response
	res := &pb.CreateCourseResponse{
		Id: courseID,
	}

	return res, nil
}

func (s *CourseService) GetCourseById(ctx context.Context, in *pb.GetCourseByIdRequest) (*pb.Course, error) {
	// Get data from request
	courseID := in.GetId()

	// Find course by id
	course, err := repo.Course.GetCourseByID(ctx, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Course with ID %d not found", courseID))
	}

	// Return response
	res := &pb.Course{
		Id:           course.ID,
		Title:        course.Title,
		Description:  course.Description,
		Price:        course.Price,
		ThumbnailUrl: course.ThumbnailURL,
		InstructorId: course.InstructorID,
		Category:     course.Category,
	}

	return res, nil
}

func (s *CourseService) GetCoursesByInstructorID(ctx context.Context, in *pb.GetCoursesByInstructorIDRequest) (*pb.GetCoursesByInstructorIDResponse, error) {
	instructorID := in.GetInstructorId()

	courses, err := repo.Course.GetCourseByInstructorID(ctx, instructorID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get courses by instructor ID: %v", err))
	}

	res := &pb.GetCoursesByInstructorIDResponse{
		Courses: make([]*pb.Course, len(courses)),
	}

	for i, course := range courses {
		res.Courses[i] = &pb.Course{
			Id:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Price:        course.Price,
			ThumbnailUrl: course.ThumbnailURL,
			InstructorId: course.InstructorID,
			Category:     course.Category,
		}
	}

	return res, nil
}

func (s *CourseService) GetCoursesByCategory(ctx context.Context, in *pb.GetCoursesByCategoryRequest) (*pb.GetCoursesByCategoryResponse, error) {
	courses, err := repo.Course.GetCourseByCategory(ctx, in.GetCategory())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get courses by category: %v", err))
	}

	res := &pb.GetCoursesByCategoryResponse{
		Courses: make([]*pb.Course, len(courses)),
	}

	for i, course := range courses {
		res.Courses[i] = &pb.Course{
			Id:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Price:        course.Price,
			ThumbnailUrl: course.ThumbnailURL,
			InstructorId: course.InstructorID,
			Category:     course.Category,
		}
	}

	return res, nil
}

func (s *CourseService) GetAllCourses(ctx context.Context, in *pb.GetAllCoursesRequest) (*pb.GetAllCoursesResponse, error) {
	courses, err := repo.Course.GetAllCourses(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get all courses: %v", err))
	}

	res := &pb.GetAllCoursesResponse{
		Courses: make([]*pb.Course, len(courses)),
	}

	for i, course := range courses {
		res.Courses[i] = &pb.Course{
			Id:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Price:        course.Price,
			ThumbnailUrl: course.ThumbnailURL,
			InstructorId: course.InstructorID,
			Category:     course.Category,
		}
	}

	return res, nil
}

func (s *CourseService) CheckCourseByID(ctx context.Context, in *pb.CheckCourseByIDRequest) (*pb.CheckCourseByIDResponse, error) {
	exists, err := repo.Course.CheckCourseByID(ctx, in.GetCourseId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to check course by ID: %v", err))
	}

	return &pb.CheckCourseByIDResponse{Success: exists}, nil
}

func (s *CourseService) UpdateCourse(ctx context.Context, in *pb.UpdateCourseRequest) (*emptypb.Empty, error) {
	// Get data from request
	courseID := in.GetId()
	courseTitle := in.GetTitle()
	courseDesc := in.GetDescription()
	coursePrice := in.GetPrice()
	courseThumbnail := in.GetThumbnailUrl()
	instructorID := in.GetInstructorId()
	category := in.GetCategory()

	// Update course
	course := &data.Course{
		ID:           courseID,
		Title:        courseTitle,
		Description:  courseDesc,
		Price:        coursePrice,
		ThumbnailURL: courseThumbnail,
		InstructorID: instructorID,
		Category:     category,
	}

	// updatedAt
	updatedAt := time.Now()

	err := repo.Course.UpdateCourse(ctx, course, updatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to update course: %v", err))
	}

	// Return response
	return &emptypb.Empty{}, nil
}

func (s *CourseService) DeleteCourse(ctx context.Context, in *pb.DeleteCourseRequest) (*emptypb.Empty, error) {
	// Get data from request
	courseID := in.GetId()

	// DeletedAt
	deletedAt := time.Now()

	// Delete course
	err := repo.Course.DeleteCourse(ctx, courseID, deletedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to delete course: %v", err))
	}

	// Return response
	return &emptypb.Empty{}, nil
}
