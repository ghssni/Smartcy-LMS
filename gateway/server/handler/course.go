package handler

import (
	"fmt"
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/server/middlewares"
	"gateway-service/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CourseHandler struct {
	courseService pb.CourseServiceClient
	lessonService pb.LessonServiceClient
}

func NewCourseHandler(courseService pb.CourseServiceClient, lessonService pb.LessonServiceClient) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
		lessonService: lessonService,
	}
}

func (h *CourseHandler) CreateCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	role := claims.Role

	if role != "instructor" {
		return utils.HandleError(c, constans.ErrForbidden, "Only instructor can create course")
	}

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
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to create course")
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
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Course with ID %s not found", courseID))
	}

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
			return utils.HandleError(c, constans.ErrInternalServerError, "Failed to get courses")
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
			return utils.HandleError(c, constans.ErrInternalServerError, "Failed to get courses by instructor")
		}
		coursesRes = res.Courses
	} else if category != "" {
		req := pb.GetCoursesByCategoryRequest{
			Category: category,
		}
		res, err := h.courseService.GetCoursesByCategory(c.Request().Context(), &req)
		if err != nil {
			return utils.HandleError(c, constans.ErrInternalServerError, "Failed to get courses by category")
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	role := claims.Role

	if role != "instructor" {
		return utils.HandleError(c, constans.ErrForbidden, "Only instructor can create course")
	}

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
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Course with ID %s not found", courseID))
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
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to update course")
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Course updated"})
}

func (h *CourseHandler) DeleteCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	role := claims.Role

	if role != "instructor" {
		return utils.HandleError(c, constans.ErrForbidden, "Only instructor can create course")
	}

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
		return utils.HandleError(c, constans.ErrNotFound, fmt.Sprintf("Course with ID %s not found", courseID))
	}

	deleteCourseReq := pb.DeleteCourseRequest{
		Id: utils.StringToUint(courseID),
	}

	_, err = h.courseService.DeleteCourse(c.Request().Context(), &deleteCourseReq)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to delete course")
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Course deleted"})
}
