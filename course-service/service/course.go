package service

import (
	"context"
	"course-service/data"
	"course-service/pb"
	"log"
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

	courseID, err := repo.Course.CreateCourse(course)
	if err != nil {
		return nil, err
	}

	// Return response
	res := &pb.CreateCourseResponse{
		Success: true,
		Id:      uint32(courseID),
	}

	return res, nil
}

func (s *CourseService) GetCourseById(ctx context.Context, in *pb.GetCourseByIdRequest) (*pb.GetCourseByIdResponse, error) {
	// Get data from request
	courseID := in.GetId()

	// Find course by id
	course, err := repo.Course.GetCourseByID(courseID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Return response
	res := &pb.GetCourseByIdResponse{
		Success: true,
		Course: &pb.Course{
			Id:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Price:        course.Price,
			ThumbnailUrl: course.ThumbnailURL,
			InstructorId: course.InstructorID,
			Category:     course.Category,
		},
	}

	return res, nil
}

func (s *CourseService) GetCoursesByInstructorID(ctx context.Context, in *pb.GetCoursesByInstructorIDRequest) (*pb.GetCoursesByInstructorIDResponse, error) {
	// Get data from request
	instructorID := in.GetInstructorId()

	// Find course by instructor id
	courses, err := repo.Course.GetCourseByInstructorID(instructorID)
	if err != nil {
		return nil, err
	}

	// Return response
	res := &pb.GetCoursesByInstructorIDResponse{
		Success: true,
		Courses: []*pb.Course{},
	}

	for _, course := range courses {
		res.Courses = append(res.Courses, &pb.Course{
			Id:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Price:        course.Price,
			ThumbnailUrl: course.ThumbnailURL,
			InstructorId: course.InstructorID,
			Category:     course.Category,
		})
	}

	return res, nil
}

func (s *CourseService) GetCoursesByCategory(ctx context.Context, in *pb.GetCoursesByCategoryRequest) (*pb.GetCoursesByCategoryResponse, error) {
	// Get data from request
	category := in.GetCategory()

	// Find course by category
	courses, err := repo.Course.GetCourseByCategory(category)
	if err != nil {
		return nil, err
	}

	// Return response
	res := &pb.GetCoursesByCategoryResponse{
		Success: true,
		Courses: []*pb.Course{},
	}

	for _, course := range courses {
		res.Courses = append(res.Courses, &pb.Course{
			Id:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Price:        course.Price,
			ThumbnailUrl: course.ThumbnailURL,
			InstructorId: course.InstructorID,
			Category:     course.Category,
		})
	}

	return res, nil
}

func (s *CourseService) GetAllCourses(ctx context.Context, in *pb.GetAllCoursesRequest) (*pb.GetAllCoursesResponse, error) {
	// Find all courses
	courses, err := repo.Course.GetAllCourses()
	if err != nil {
		return nil, err
	}

	// Return response
	res := &pb.GetAllCoursesResponse{
		Success: true,
		Courses: []*pb.Course{},
	}

	for _, course := range courses {
		res.Courses = append(res.Courses, &pb.Course{
			Id:           course.ID,
			Title:        course.Title,
			Description:  course.Description,
			Price:        course.Price,
			ThumbnailUrl: course.ThumbnailURL,
			InstructorId: course.InstructorID,
			Category:     course.Category,
		})
	}

	return res, nil
}

func (s *CourseService) CheckCourseByID(ctx context.Context, in *pb.CheckCourseByIDRequest) (*pb.CheckCourseByIDResponse, error) {
	// Get data from request
	courseID := in.GetCourseId()

	// Check if course ID exists
	exists, err := repo.Course.CheckCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	// Return response
	res := &pb.CheckCourseByIDResponse{
		Success: exists,
	}

	return res, nil
}

func (s *CourseService) UpdateCourse(ctx context.Context, in *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
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
		Title:        courseTitle,
		Description:  courseDesc,
		Price:        coursePrice,
		ThumbnailURL: courseThumbnail,
		InstructorID: instructorID,
		Category:     category,
	}

	err := repo.Course.UpdateCourse(courseID, course)
	if err != nil {
		return nil, err
	}

	// Return response
	res := &pb.UpdateCourseResponse{
		Success: true,
	}

	return res, nil
}

func (s *CourseService) DeleteCourse(ctx context.Context, in *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	// Get data from request
	courseID := in.GetId()

	// Delete course
	err := repo.Course.DeleteCourse(courseID)
	if err != nil {
		return nil, err
	}

	// Return response
	res := &pb.DeleteCourseResponse{
		Success: true,
	}

	return res, nil
}
