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

type ReviewService struct{}

func NewReviewService() *ReviewService { return &ReviewService{} }

func (r *ReviewService) CreateReview(ctx context.Context, in *pb.CreateReviewRequest) (*pb.Review, error) {
	// Get data from request
	courseID := in.GetCourseId()
	studentID := in.GetStudentId()
	rating := in.GetRating()
	reviewText := in.GetReviewText()

	// Insert data to database
	review := &data.Review{
		CourseID:   courseID,
		StudentID:  studentID,
		Rating:     rating,
		ReviewText: reviewText,
	}

	// CreateAt and UpdatedAt
	createdAt, updatedAt := time.Now(), time.Now()

	// Insert data to database
	reviewId, err := repo.Review.CreateReview(ctx, review, createdAt, updatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to create review: %v", err))
	}

	res := &pb.Review{
		Id:         reviewId,
		CourseId:   courseID,
		StudentId:  studentID,
		Rating:     rating,
		ReviewText: reviewText,
		CreatedAt:  createdAt.String(),
		UpdatedAt:  updatedAt.String(),
	}

	return res, nil
}

func (r *ReviewService) GetReview(ctx context.Context, in *pb.GetReviewRequest) (*pb.Review, error) {
	// Get data from request
	reviewID := in.GetReviewId()

	// Find review by id
	review, err := repo.Review.GetReview(ctx, reviewID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Review with ID %d not found", reviewID))
	}

	res := &pb.Review{
		Id:         review.ID,
		CourseId:   review.CourseID,
		StudentId:  review.StudentID,
		Rating:     review.Rating,
		ReviewText: review.ReviewText,
		CreatedAt:  review.CreatedAt.String(),
		UpdatedAt:  review.UpdatedAt.String(),
	}

	return res, nil
}

func (r *ReviewService) GetReviewsByStudent(ctx context.Context, in *pb.GetReviewsByStudentRequest) (*pb.GetReviewsByStudentResponse, error) {
	courseID := in.GetCourseId()
	studentID := in.GetStudentId()

	// Get reviews by student
	reviews, err := repo.Review.GetReviewsByStudent(ctx, courseID, studentID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get reviews: %v", err))
	}

	var resReviews []*pb.Review
	for _, review := range reviews {
		resReview := &pb.Review{
			Id:         review.ID,
			CourseId:   review.CourseID,
			StudentId:  review.StudentID,
			Rating:     review.Rating,
			ReviewText: review.ReviewText,
			CreatedAt:  review.CreatedAt.String(),
			UpdatedAt:  review.UpdatedAt.String(),
		}
		resReviews = append(resReviews, resReview)
	}

	res := &pb.GetReviewsByStudentResponse{
		Reviews: resReviews,
	}

	return res, nil
}

func (r *ReviewService) GetAverageRatingByCourse(ctx context.Context, in *pb.GetAverageRatingByCourseRequest) (*pb.GetAverageRatingByCourseResponse, error) {
	courseID := in.GetCourseId()

	// Get average rating by course
	averageRating, err := repo.Review.GetAverageRatingByCourse(ctx, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get average rating: %v", err))
	}

	res := &pb.GetAverageRatingByCourseResponse{
		AverageRating: averageRating,
	}

	return res, nil
}

func (r *ReviewService) GetTotalReviewsByCourse(ctx context.Context, in *pb.GetTotalReviewsByCourseRequest) (*pb.GetTotalReviewsByCourseResponse, error) {
	courseID := in.GetCourseId()

	// Get total reviews by course
	totalReviews, err := repo.Review.GetTotalReviewsByCourse(ctx, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get total reviews: %v", err))
	}

	res := &pb.GetTotalReviewsByCourseResponse{
		TotalReviews: totalReviews,
	}

	return res, nil
}

func (r *ReviewService) ListReviews(ctx context.Context, in *pb.ListReviewsRequest) (*pb.ListReviewsResponse, error) {
	courseID := in.GetCourseId()

	// Get all reviews by course
	reviews, err := repo.Review.ListReviews(ctx, courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Failed to get reviews: %v", err))
	}

	var resReviews []*pb.Review
	for _, review := range reviews {
		resReview := &pb.Review{
			Id:         review.ID,
			CourseId:   review.CourseID,
			StudentId:  review.StudentID,
			Rating:     review.Rating,
			ReviewText: review.ReviewText,
			CreatedAt:  review.CreatedAt.String(),
			UpdatedAt:  review.UpdatedAt.String(),
		}
		resReviews = append(resReviews, resReview)
	}

	res := &pb.ListReviewsResponse{
		Reviews: resReviews,
	}

	return res, nil
}

func (r *ReviewService) UpdateReview(ctx context.Context, in *pb.UpdateReviewRequest) (*pb.Review, error) {
	// Get data from request
	reviewID := in.GetReviewId()
	rating := in.GetRating()
	reviewText := in.GetReviewText()

	// Find review by id
	review, err := repo.Review.GetReview(ctx, reviewID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Review with ID %d not found", reviewID))
	}

	// Update review
	review.Rating = rating
	review.ReviewText = reviewText
	updatedAt := time.Now()

	err = repo.Review.UpdateReview(ctx, review, updatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to update review: %v", err))
	}

	res := &pb.Review{
		Id:         review.ID,
		CourseId:   review.CourseID,
		StudentId:  review.StudentID,
		Rating:     review.Rating,
		ReviewText: review.ReviewText,
		CreatedAt:  review.CreatedAt.String(),
		UpdatedAt:  updatedAt.String(),
	}

	return res, nil
}

func (r *ReviewService) DeleteReview(ctx context.Context, in *pb.DeleteReviewRequest) (*emptypb.Empty, error) {
	// Get data from request
	reviewID := in.GetReviewId()

	// Find review by id
	review, err := repo.Review.GetReview(ctx, reviewID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Review with ID %d not found", reviewID))
	}

	// Delete review
	deletedAt := time.Now()

	err = repo.Review.DeleteReview(ctx, review.ID, deletedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to delete review: %v", err))
	}

	return &emptypb.Empty{}, nil
}
