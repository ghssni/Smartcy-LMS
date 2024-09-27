package handler

import (
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/server/middlewares"
	"gateway-service/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type LearningProgressHandler struct {
	// Add the learning progress service client here
	learningProgressService pb.LearningProgressServiceClient
}

func NewLearningProgressHandler(learningProgressService pb.LearningProgressServiceClient) *LearningProgressHandler {
	return &LearningProgressHandler{
		learningProgressService: learningProgressService,
	}
}

func (h *LearningProgressHandler) MarkLessonAsCompleted(c echo.Context) error {

	// Get userId from JWT
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	userID := claims.UserID

	// Check User is enrolled and paid for the course using enrollment service

	// Bind the mark request struct
	markRequest := new(model.LPRequest)
	err := c.Bind(&markRequest)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the mark request struct
	err = config.Validator.Struct(markRequest)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Call the service to mark the lesson as completed
	_, err = h.learningProgressService.MarkLessonAsCompleted(c.Request().Context(), &pb.MarkLessonAsCompletedRequest{
		EnrollmentId: markRequest.EnrollmentID,
		LessonId:     markRequest.LessonID,
		UserId:       userID,
	})

	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, model.JsonResponse{
		Status:  "success",
		Message: "Lesson marked as completed",
		Data:    nil,
	})

}

func (h *LearningProgressHandler) ResetLessonMark(c echo.Context) error {

	// Bind the reset request struct
	resetRequest := new(model.LPRequest)
	err := c.Bind(&resetRequest)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the reset request struct
	err = config.Validator.Struct(resetRequest)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Call the service to reset all lesson marks
	_, err = h.learningProgressService.ResetLessonMark(c.Request().Context(), &pb.ResetLessonMarkRequest{
		EnrollmentId: resetRequest.EnrollmentID,
		LessonId:     resetRequest.LessonID,
	})

	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to reset lesson mark")
	}

	return c.JSON(http.StatusCreated, model.JsonResponse{
		Status:  "success",
		Message: "Lesson unmarked",
		Data:    nil,
	})

}

func (h *LearningProgressHandler) ResetAllLessonMarks(c echo.Context) error {

	// Bind the reset request struct
	reqPayload := &struct {
		EnrollmentID uint32 `json:"enrollment_id" validate:"required"`
	}{}

	err := c.Bind(&reqPayload)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the reset request struct
	err = config.Validator.Struct(reqPayload)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Call the service to reset all lesson marks
	_, err = h.learningProgressService.ResetAllLessonMarks(c.Request().Context(), &pb.ResetAllLessonMarksRequest{
		EnrollmentId: reqPayload.EnrollmentID,
	})

	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to reset all lesson marks")
	}

	return c.JSON(http.StatusCreated, model.JsonResponse{
		Status:  "success",
		Message: "All lesson marks reset",
		Data:    nil,
	})
}

func (h *LearningProgressHandler) GetTotalCompletedLessons(c echo.Context) error {
	// Get userId from JWT
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	userID := claims.UserID

	enrollmentID := c.Param("enrollment_id")

	// Call the service to get total completed lessons
	res, err := h.learningProgressService.GetTotalCompletedLessons(c.Request().Context(), &pb.GetTotalCompletedLessonsRequest{
		EnrollmentId: utils.StringToUint(enrollmentID),
		UserId:       userID,
	})

	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, err.Error())
	}

	// Map the response to the JSON response
	respSummary := new(model.CompletedProgressResponse)
	respSummary.EnrollmentID = res.EnrollmentId
	respSummary.TotalCompleted = res.TotalCompleted

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  "success",
		Message: "Total completed lessons retrieved",
		Data:    respSummary,
	})
}

func (h *LearningProgressHandler) GetTotalCompletedProgress(c echo.Context) error {
	// Get userId from JWT
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JWTCustomClaims)
	userID := claims.UserID

	log.Printf("User ID: %v", userID)

	// Call the service to get total completed progress
	res, err := h.learningProgressService.GetTotalCompletedProgress(c.Request().Context(), &pb.GetTotalCompletedProgressRequest{
		UserId: userID,
	})

	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to get total completed progress")
	}

	log.Printf("Total completed progress: %v", res)

	// Map the response to the JSON response
	respSummary := make([]*model.CompletedProgressResponse, 0)
	for _, v := range res.Progress {
		respSummary = append(respSummary, &model.CompletedProgressResponse{
			EnrollmentID:   v.EnrollmentId,
			TotalCompleted: v.TotalCompleted,
		})
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  "success",
		Message: "Total completed progress retrieved",
		Data:    respSummary,
	})
}

func (h *LearningProgressHandler) ListLearningProgress(c echo.Context) error {

	//user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(*middlewares.JWTCustomClaims)
	//userID := claims.UserID

	enrollmentID := c.Param("enrollment_id")
	if enrollmentID == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Please provide enrollment id as parameter")
	}

	res, err := h.learningProgressService.ListLearningProgress(c.Request().Context(), &pb.ListLearningProgressRequest{
		EnrollmentId: utils.StringToUint(enrollmentID),
	})

	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to list learning progress")
	}

	respSummary := make([]*model.LearningProgress, len(res.LearningProgress))
	for i, v := range res.LearningProgress {
		log.Printf("Last Accessed: %v", v.LastAccessed.AsTime().Format(time.RFC3339))
		respSummary[i] = &model.LearningProgress{
			EnrollmentID: v.EnrollmentId,
			LessonID:     v.LessonId,
			Status:       v.Status,
			LastAccessed: utils.FormatTimestamp(v.LastAccessed),
			CompletedAt:  utils.FormatTimestamp(v.CompletedAt),
		}
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  "success",
		Message: "Learning progress retrieved",
		Data:    respSummary,
	})

}

func (h *LearningProgressHandler) UpdateLastAccessed(c echo.Context) error {

	reqPayload := new(model.LPRequest)

	err := c.Bind(&reqPayload)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	err = config.Validator.Struct(reqPayload)
	if err != nil {
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	_, err = h.learningProgressService.UpdateLastAccessed(c.Request().Context(), &pb.UpdateLastAccessedRequest{
		EnrollmentId: reqPayload.EnrollmentID,
		LessonId:     reqPayload.LessonID,
	})

	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to update last accessed")
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  "success",
		Message: "Last accessed updated",
		Data:    nil,
	})
}
