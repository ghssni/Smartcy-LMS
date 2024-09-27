package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Review struct {
	ID         uint32     `db:"id"`
	CourseID   uint32     `db:"course_id"`
	StudentID  string     `db:"student_id"`
	Rating     uint32     `db:"rating"`
	ReviewText string     `db:"review_text"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}

func (r *Review) CreateReview(ctx context.Context, review *Review, createdAt, updatedAt time.Time) (uint32, error) {
	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Check if the student is enrolled in the course
	var enrolled bool
	enrollmentCheckQuery := `
        SELECT EXISTS(
            SELECT 1 FROM enrollments WHERE course_id = $1 AND student_id = $2 AND payment_status = 'paid'
        )`
	err = tx.QueryRowContext(ctx, enrollmentCheckQuery, review.CourseID, review.StudentID).Scan(&enrolled)
	if err != nil {
		return 0, err
	}

	if !enrolled {
		return 0, errors.New("student is not enrolled in the course or has not paid")
	}

	// Check if the student has completed 50% of the lessons
	var totalLessons, completedLessons int
	lessonCountQuery := `SELECT COUNT(*) FROM lessons WHERE course_id = $1 AND deleted_at IS NULL`
	err = tx.QueryRowContext(ctx, lessonCountQuery, review.CourseID).Scan(&totalLessons)
	if err != nil {
		return 0, err
	}

	completedLessonQuery := `
        SELECT COUNT(*) FROM learning_progress lp
        INNER JOIN enrollments e ON lp.enrollment_id = e.id
        WHERE e.course_id = $1 AND e.student_id = $2 AND lp.status = true`
	err = tx.QueryRowContext(ctx, completedLessonQuery, review.CourseID, review.StudentID).Scan(&completedLessons)
	if err != nil {
		return 0, err
	}

	if completedLessons < (totalLessons / 2) {
		return 0, errors.New("student has not completed 50% of the course")
	}

	// Check if a review already exists and is deleted
	var existingReviewID uint32
	reviewCheckQuery := `
        SELECT id FROM reviews WHERE course_id = $1 AND student_id = $2 AND deleted_at IS NOT NULL`
	err = tx.QueryRowContext(ctx, reviewCheckQuery, review.CourseID, review.StudentID).Scan(&existingReviewID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	if existingReviewID != 0 {
		// Update the existing review if it was deleted
		updateReviewQuery := `
            UPDATE reviews SET rating = $1, review_text = $2, created_at = $3, updated_at = $4, deleted_at = NULL
            WHERE id = $5`
		_, err = tx.ExecContext(ctx, updateReviewQuery, review.Rating, review.ReviewText, createdAt, updatedAt, existingReviewID)
		if err != nil {
			return 0, err
		}
		review.ID = existingReviewID
	} else {
		// Insert the review
		insertReviewQuery := `
            INSERT INTO reviews (course_id, student_id, rating, review_text, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
		err = tx.QueryRowContext(ctx, insertReviewQuery, review.CourseID, review.StudentID, review.Rating, review.ReviewText, createdAt, updatedAt).Scan(&review.ID)
		if err != nil {
			return 0, err
		}
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return review.ID, nil
}

func (r *Review) GetReview(ctx context.Context, reviewID uint32) (*Review, error) {
	var review Review

	err := db.GetContext(ctx, &review, `SELECT id, course_id, student_id, rating, review_text, created_at, updated_at FROM reviews WHERE id = $1 AND deleted_at IS NULL`, reviewID)
	if err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *Review) GetReviewsByStudent(ctx context.Context, courseID uint32, studentID string) ([]Review, error) {
	sqlStatement := `SELECT id, course_id, student_id, rating, review_text, created_at, updated_at FROM reviews WHERE course_id = $1 AND student_id = $2 AND deleted_at IS NULL`

	var reviews []Review
	err := db.SelectContext(ctx, &reviews, sqlStatement, courseID, studentID)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *Review) GetTotalReviewsByCourse(ctx context.Context, courseID uint32) (uint32, error) {
	var totalReviews uint32
	sqlStatement := `SELECT COUNT(*) FROM reviews WHERE course_id = $1 AND deleted_at IS NULL`

	err := db.GetContext(ctx, &totalReviews, sqlStatement, courseID)
	if err != nil {
		return 0, err
	}

	return totalReviews, nil
}

func (r *Review) GetAverageRatingByCourse(ctx context.Context, courseID uint32) (float32, error) {
	var averageRating float32
	sqlStatement := `SELECT AVG(rating) FROM reviews WHERE course_id = $1 AND deleted_at IS NULL`

	err := db.GetContext(ctx, &averageRating, sqlStatement, courseID)
	if err != nil {
		return 0, err
	}

	return averageRating, nil
}

func (r *Review) ListReviews(ctx context.Context, courseID uint32) ([]Review, error) {
	sqlStatement := `SELECT id, course_id, student_id, rating, review_text, created_at, updated_at FROM reviews WHERE course_id = $1 AND deleted_at IS NULL`

	var reviews []Review
	err := db.SelectContext(ctx, &reviews, sqlStatement, courseID)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *Review) UpdateReview(ctx context.Context, review *Review, updatedAt time.Time) error {
	sqlStatement := `UPDATE reviews SET rating = $1, review_text = $2, updated_at = $3 WHERE id = $4`

	_, err := db.ExecContext(ctx, sqlStatement, review.Rating, review.ReviewText, updatedAt, review.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Review) DeleteReview(ctx context.Context, reviewID uint32, deletedAt time.Time) error {
	sqlStatement := `UPDATE reviews SET deleted_at = $1 WHERE id = $2`

	_, err := db.ExecContext(ctx, sqlStatement, deletedAt, reviewID)
	if err != nil {
		return err
	}

	return nil
}
