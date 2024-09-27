package handler

import (
	"gateway-service/config"
	"gateway-service/constans"
	"gateway-service/model"
	"gateway-service/pb"
	"gateway-service/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ReviewHandler struct {
	reviewService pb.ReviewServiceClient
}

func NewReviewHandler(reviewService pb.ReviewServiceClient) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	// Get data from request
	review := new(model.Review)
	err := c.Bind(&review)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	// Validate the review struct
	err = config.Validator.Struct(review)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	req := pb.CreateReviewRequest{
		CourseId:   review.CourseID,
		StudentId:  review.StudentID,
		Rating:     review.Rating,
		ReviewText: review.ReviewText,
	}

	res, err := h.reviewService.CreateReview(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}

func (h *ReviewHandler) ListReviews(c echo.Context) error {

	courseId := c.Param("course_id")
	if courseId == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Course ID is required")
	}

	req := pb.ListReviewsRequest{
		CourseId: utils.StringToUint(courseId),
	}

	res, err := h.reviewService.ListReviews(c.Request().Context(), &req)
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, "Failed to get reviews")
	}

	reviews := make([]*model.Review, 0)
	for _, review := range res.Reviews {
		reviews = append(reviews, &model.Review{
			ID:         review.Id,
			CourseID:   review.CourseId,
			StudentID:  review.StudentId,
			Rating:     review.Rating,
			ReviewText: review.ReviewText,
			CreatedAt:  review.CreatedAt,
			UpdatedAt:  review.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  constans.ResponseStatusSuccess,
		Message: "Reviews found",
		Data:    reviews,
	})
}

func (h *ReviewHandler) UpdateReviewRequest(c echo.Context) error {
	reqPayload := new(model.Review)

	err := c.Bind(&reqPayload)
	if err != nil {
		return utils.HandleError(c, constans.ErrBadRequest, "Invalid request body")
	}

	err = config.Validator.Struct(reqPayload)
	if err != nil {
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Get review by id
	res1 := pb.GetReviewRequest{
		ReviewId: reqPayload.ID,
	}

	_, err = h.reviewService.GetReview(c.Request().Context(), &res1)
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, "Review not found")
	}

	req2 := pb.UpdateReviewRequest{
		ReviewId:   reqPayload.ID,
		CourseId:   reqPayload.CourseID,
		Rating:     reqPayload.Rating,
		ReviewText: reqPayload.ReviewText,
	}

	res2, err := h.reviewService.UpdateReview(c.Request().Context(), &req2)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to update review")
	}

	review := &model.Review{
		ID:         res2.Id,
		CourseID:   res2.CourseId,
		StudentID:  res2.StudentId,
		Rating:     res2.Rating,
		ReviewText: res2.ReviewText,
		CreatedAt:  res2.CreatedAt,
		UpdatedAt:  res2.UpdatedAt,
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  constans.ResponseStatusSuccess,
		Message: "Review updated",
		Data:    review,
	})
}

func (h *ReviewHandler) DeleteReview(c echo.Context) error {
	reviewID := c.Param("review_id")

	if reviewID == "" {
		return utils.HandleError(c, constans.ErrBadRequest, "Review ID is required")
	}

	// Get review by id
	res1 := pb.GetReviewRequest{
		ReviewId: utils.StringToUint(reviewID),
	}

	_, err := h.reviewService.GetReview(c.Request().Context(), &res1)
	if err != nil {
		return utils.HandleError(c, constans.ErrNotFound, "Review not found")
	}

	req2 := pb.DeleteReviewRequest{
		ReviewId: utils.StringToUint(reviewID),
	}

	_, err = h.reviewService.DeleteReview(c.Request().Context(), &req2)
	if err != nil {
		return utils.HandleError(c, constans.ErrInternalServerError, "Failed to delete review")
	}

	return c.JSON(http.StatusOK, model.JsonResponse{
		Status:  constans.ResponseStatusSuccess,
		Message: "Review deleted",
	})
}
