package handler

import (
	"fmt"
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/utils"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
)

type LessonHandler struct {
	lessonService pb.LessonServiceClient
}

func NewLessonHandler(lessonService pb.LessonServiceClient) *LessonHandler {
	return &LessonHandler{
		lessonService: lessonService,
	}
}

func (h *LessonHandler) CreateLesson(c echo.Context) error {
	lesson := new(model.Lesson)

	// Bind the lesson struct
	err := c.Bind(&lesson)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the lesson struct
	err = config.Validator.Struct(lesson)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Check user Role from jwt

	// Check if the course exist

	req := pb.CreateLessonRequest{
		CourseId:   lesson.CourseID,
		Title:      lesson.Title,
		ContentUrl: lesson.ContentURL,
		LessonType: lesson.LessonType,
		Sequence:   lesson.Sequence,
	}

	// Do the gRPC call
	res, err := h.lessonService.CreateLesson(c.Request().Context(), &req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.AlreadyExists:
				return utils.HandleError(c, constans.ErrConflict, st.Message())
			default:
				return utils.HandleError(c, constans.ErrNotFound, st.Message())
			}
		} else {
			return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
		}
	}

	lesson.ID = res.Id

	return c.JSON(http.StatusCreated, model.JsonResponse{Status: "success", Message: "Lesson created", Data: lesson})
}

func (h *LessonHandler) GetLesson(c echo.Context) error {
	id := c.Param("id")

	// Check user Role from jwt

	req := pb.GetLessonRequest{
		Id: utils.StringToUint(id),
	}

	// Do the gRPC call
	res, err := h.lessonService.GetLesson(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, "Lesson not found")
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Lesson found", Data: res.Lesson})
}

func (h *LessonHandler) GetLessonBySequence(c echo.Context) error {
	seq := c.Param("sequence")
	courseId := c.Param("course_id")

	log.Println("seq", seq)
	log.Println("courseId", courseId)

	// Check user Role from jwt

	req := pb.GetLessonBySequenceRequest{
		CourseId: utils.StringToUint(courseId),
		Sequence: utils.StringToUint(seq),
	}

	// Do the gRPC call
	res, err := h.lessonService.GetLessonBySequence(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Lesson found", Data: res.Lesson})
}

func (h *LessonHandler) GetAllLessons(c echo.Context) error {
	courseId := c.Param("course_id")

	lessonType := c.QueryParam("type")
	title := c.QueryParam("title")

	lessonRes := make([]*pb.LessonSummary, 0)

	if lessonType == "" && title == "" {
		req := pb.ListLessonsRequest{
			CourseId: utils.StringToUint(courseId),
		}

		// Do the gRPC call
		res, err := h.lessonService.ListLessons(c.Request().Context(), &req)
		if err != nil {
			return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
		}

		lessonRes = res.Lessons
	} else if lessonType != "" && title != "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide only title or type query parameter")
	} else if lessonType != "" {
		req := pb.SearchLessonsByTypeRequest{
			CourseId:   utils.StringToUint(courseId),
			LessonType: lessonType,
		}

		// Do the gRPC call
		res, err := h.lessonService.SearchLessonsByType(c.Request().Context(), &req)
		if err != nil {
			return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
		}

		lessonRes = res.Lessons
	} else if title != "" {
		req := pb.SearchLessonsByTitleRequest{
			CourseId: utils.StringToUint(courseId),
			Title:    title,
		}

		// Do the gRPC call
		res, err := h.lessonService.SearchLessonsByTitle(c.Request().Context(), &req)
		if err != nil {
			return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
		}

		lessonRes = res.Lessons
	} else {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid query parameter")
	}

	// check if lesson is empty
	if len(lessonRes) == 0 {
		return utils.HandleError(c, constans.ErrNotFound, "No lesson found")
	}

	// Check user Role from jwt

	message := fmt.Sprintf("Found %d lessons", len(lessonRes))

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: message, Data: lessonRes})
}

func (h *LessonHandler) UpdateLesson(c echo.Context) error {
	id := c.Param("id")
	lesson := new(model.Lesson)

	// Bind the lesson struct
	err := c.Bind(&lesson)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the lesson struct
	err = config.Validator.Struct(lesson)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Check user Role from jwt

	req := pb.UpdateLessonRequest{
		Id:         utils.StringToUint(id),
		CourseId:   lesson.CourseID,
		Title:      lesson.Title,
		ContentUrl: lesson.ContentURL,
		LessonType: lesson.LessonType,
		Sequence:   lesson.Sequence,
	}

	// Do the gRPC call
	res, err := h.lessonService.UpdateLesson(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Lesson updated", Data: res.Lesson})
}

func (h *LessonHandler) DeleteLesson(c echo.Context) error {
	id := c.Param("id")

	// Check user Role from jwt

	req := pb.DeleteLessonRequest{
		Id: utils.StringToUint(id),
	}

	// Do the gRPC call
	_, err := h.lessonService.DeleteLesson(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Lesson deleted"})
}

func (h *LessonHandler) GetTotalLessonsByCourseID(c echo.Context) error {
	lesson := new(model.Lesson)

	// Bind the lesson struct
	err := c.Bind(&lesson)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the lesson struct
	err = config.Validator.Struct(lesson)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Check user Role from jwt

	req := pb.GetTotalLessonsRequest{
		CourseId: lesson.CourseID,
	}

	// Do the gRPC call
	res, err := h.lessonService.GetTotalLessons(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JsonResponse{Status: "success", Message: "Total lessons found", Data: res.Total})
}
