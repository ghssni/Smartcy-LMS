package handler

import (
	"fmt"
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/utils"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CourseHandler struct {
	courseService pb.CourseServiceClient
	reviewService pb.ReviewServiceClient
}

func NewCourseHandler(courseService pb.CourseServiceClient, reviewService pb.ReviewServiceClient) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
		reviewService: reviewService,
	}
}

func (h *CourseHandler) CreateCourse(c echo.Context) error {
	course := new(model.Course)

	err := c.Bind(&course)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the order struct
	err = config.Validator.Struct(course)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Check user Role from jwt

	req := pb.CreateCourseRequest{
		Title:        course.Title,
		Description:  course.Description,
		Price:        course.Price,
		ThumbnailUrl: course.ThumbnailURL,
		InstructorId: course.InstructorID,
		Category:     course.Category,
	}

	// Do the gRPC call
	res, err := h.courseService.CreateCourse(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	course.ID = res.Id

	return c.JSON(http.StatusCreated, model.JsonResponse{Status: "success", Message: "Course created", Data: course})
}

func (h *CourseHandler) GetCourseByID(c echo.Context) error {
	courseID := c.Param("id")

	req := pb.GetCourseByIdRequest{
		Id: utils.StringToUint(courseID),
	}

	res, err := h.courseService.GetCourseById(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, err.Error())
	}

	averageRating, err := h.reviewService.GetAverageRatingByCourse(c.Request().Context(), &pb.GetAverageRatingByCourseRequest{CourseId: res.Id})
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	totalReview, err := h.reviewService.GetTotalReviewsByCourse(c.Request().Context(), &pb.GetTotalReviewsByCourseRequest{CourseId: res.Id})
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	log.Printf("averageRating: %v", averageRating)
	log.Printf("totalReview: %v", totalReview)

	course := model.CourseWithReview{
		ID:            res.Id,
		Title:         res.Title,
		Description:   res.Description,
		Price:         res.Price,
		ThumbnailURL:  res.ThumbnailUrl,
		InstructorID:  res.InstructorId,
		Category:      res.Category,
		CreatedAt:     utils.ProtoTimeToDate(res.CreatedAt),
		UpdatedAt:     utils.ProtoTimeToDate(res.CreatedAt),
		AverageRating: res.AverageRating,
		TotalReviews:  res.TotalReviews,
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Course found", Data: course})
}

func (h *CourseHandler) GetAllCourses(c echo.Context) error {

	category := c.QueryParam("category")
	instructorID := c.QueryParam("instructor")

	coursesRes := make([]*pb.Course, 0)

	if category == "" && instructorID == "" {
		req := pb.GetAllCoursesRequest{}
		res, err := h.courseService.GetAllCourses(c.Request().Context(), &req)
		if err != nil {
			return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
		}
		coursesRes = res.Courses
	} else if category != "" && instructorID != "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide only category or instructor query parameter")
	} else if instructorID != "" {
		req := pb.GetCoursesByInstructorIDRequest{
			InstructorId: instructorID,
		}
		res, err := h.courseService.GetCoursesByInstructorID(c.Request().Context(), &req)
		if err != nil {
			return utils.HandleError(c, constans.ErrInternalServerError)
		}
		coursesRes = res.Courses
	} else if category != "" {
		req := pb.GetCoursesByCategoryRequest{
			Category: category,
		}
		res, err := h.courseService.GetCoursesByCategory(c.Request().Context(), &req)
		if err != nil {
			return utils.HandleError(c, constans.ErrInternalServerError)
		}
		coursesRes = res.Courses
	} else {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide category or instructor query parameter")
	}

	courses := make([]model.CourseWithReview, len(coursesRes))
	for i, course := range coursesRes {
		courses[i] = model.CourseWithReview{
			ID:            course.Id,
			Title:         course.Title,
			Description:   course.Description,
			Price:         course.Price,
			ThumbnailURL:  course.ThumbnailUrl,
			InstructorID:  course.InstructorId,
			Category:      course.Category,
			UpdatedAt:     utils.ProtoTimeToDate(course.UpdatedAt),
			AverageRating: course.AverageRating,
			TotalReviews:  course.TotalReviews,
		}
	}

	message := fmt.Sprintf("Found %d courses", len(courses))

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: message, Data: courses})
}

func (h *CourseHandler) UpdateCourse(c echo.Context) error {
	courseID := c.Param("id")

	// check courseID is valid
	if courseID == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide course id as parameter")
	}

	// Get Course by ID
	getCourseReq := pb.GetCourseByIdRequest{
		Id: utils.StringToUint(courseID),
	}

	res, err := h.courseService.GetCourseById(c.Request().Context(), &getCourseReq)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	course := model.Course{
		ID:           res.Id,
		Title:        res.Title,
		Description:  res.Description,
		Price:        res.Price,
		ThumbnailURL: res.ThumbnailUrl,
		InstructorID: res.InstructorId,
		Category:     res.Category,
	}

	err = c.Bind(&course)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the order struct
	err = config.Validator.Struct(course)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	updateCourseReq := pb.UpdateCourseRequest{
		Id:           course.ID,
		Title:        course.Title,
		Description:  course.Description,
		Price:        course.Price,
		ThumbnailUrl: course.ThumbnailURL,
		InstructorId: course.InstructorID,
		Category:     course.Category,
	}

	_, err = h.courseService.UpdateCourse(c.Request().Context(), &updateCourseReq)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Course updated"})
}

func (h *CourseHandler) DeleteCourse(c echo.Context) error {
	courseID := c.Param("id")

	// check courseID is valid
	if courseID == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide course id as parameter")
	}

	// Get Course by ID
	getCourseReq := pb.GetCourseByIdRequest{
		Id: utils.StringToUint(courseID),
	}

	_, err := h.courseService.GetCourseById(c.Request().Context(), &getCourseReq)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	deleteCourseReq := pb.DeleteCourseRequest{
		Id: utils.StringToUint(courseID),
	}

	_, err = h.courseService.DeleteCourse(c.Request().Context(), &deleteCourseReq)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Course deleted"})
}
