package data

import (
	"context"
	"time"
)

type LearningProgress struct {
	ID           uint32     `db:"id"`
	EnrollmentID uint32     `db:"enrollment_id"`
	LessonID     uint32     `db:"lesson_id"`
	Status       bool       `db:"status"`
	LastAccessed *time.Time `db:"last_accessed"`
	CompletedAt  *time.Time `db:"completed_at"`
}

type CompletedProgress struct {
	EnrollmentID   uint32 `db:"enrollment_id"`
	TotalCompleted uint32 `db:"total_completed"`
}

// MarkLessonAsCompleted marks a lesson as completed for a specific enrollment.
// It inserts a new record into the learning_progress table with the provided details.
//
// Parameters:
// - ctx: The context for the database operation.
// - enrollmentID: The ID of the enrollment.
// - lessonID: The ID of the lesson.
// - lastAccessed: The timestamp when the lesson was last accessed.
// - completedAt: The timestamp when the lesson was completed.
//
// Returns:
// - error: An error object if the operation fails, otherwise nil.
func (lp *LearningProgress) MarkLessonAsCompleted(ctx context.Context, enrollmentID, lessonID uint32, lastAccessed, completedAt time.Time) error {
	query := `INSERT INTO learning_progress (enrollment_id, lesson_id, status, last_accessed, completed_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.ExecContext(ctx, query, enrollmentID, lessonID, true, lastAccessed, completedAt)
	if err != nil {
		return err
	}

	return nil
}

// ResetLessonMark resets the completion status of a specific lesson for a specific enrollment.
// It updates the status to false, sets the completed_at timestamp to NULL, and sets the last accessed timestamp to NULL
// for the record where the enrollment_id and lesson_id match.
//
// Parameters:
// - ctx: The context for the database operation.
// - enrollmentID: The ID of the enrollment.
// - lessonID: The ID of the lesson.
//
// Returns:
// - error: An error object if the operation fails, otherwise nil.
func (lp *LearningProgress) ResetLessonMark(ctx context.Context, enrollmentID, lessonID uint32) error {

	query := `UPDATE learning_progress SET status = false, completed_at = NULL, last_accessed = NULL WHERE enrollment_id = $1 AND lesson_id = $2`

	_, err := db.ExecContext(ctx, query, enrollmentID, lessonID)
	if err != nil {
		return err
	}

	return nil
}

// ResetAllLessonMarks resets the completion status of all lessons for a specific enrollment.
// It updates the status to false, sets the last accessed timestamp to NULL, and sets the completed_at timestamp to NULL
// for all records where the enrollment_id matches.
//
// Parameters:
// - ctx: The context for the database operation.
// - enrollmentID: The ID of the enrollment.
//
// Returns:
// - error: An error object if the operation fails, otherwise nil.
func (lp *LearningProgress) ResetAllLessonMarks(ctx context.Context, enrollmentID uint32) error {
	query := `UPDATE learning_progress SET status = false, last_accessed = NULL, completed_at = NULL WHERE enrollment_id = $1`

	_, err := db.ExecContext(ctx, query, enrollmentID)
	if err != nil {
		return err
	}

	return nil
}

// GetTotalCompletedLessons retrieves the total number of completed lessons for a specific enrollment.
// It counts the records in the learning_progress table where the status is true and deleted_at is NULL.
//
// Parameters:
// - ctx: The context for the database operation.
// - enrollmentID: The ID of the enrollment.
//
// Returns:
// - *CompletedProgress: A pointer to a CompletedProgress struct containing the total completed lessons.
// - error: An error object if the operation fails, otherwise nil.
func (lp *LearningProgress) GetTotalCompletedLessons(ctx context.Context, enrollmentID uint32) (*CompletedProgress, error) {
	query := `SELECT COUNT(*) FROM learning_progress WHERE enrollment_id = $1 AND status = true `

	completedProgress := new(CompletedProgress)

	err := db.GetContext(ctx, &completedProgress, query, enrollmentID)
	if err != nil {
		return nil, err
	}

	return completedProgress, nil
}

// GetTotalCompletedProgress retrieves the total number of completed lessons for all enrollments.
// It groups the records in the learning_progress table by enrollment_id where the status is true and deleted_at is NULL.
//
// Parameters:
// - ctx: The context for the database operation.
//
// Returns:
// - []CompletedProgress: A slice of CompletedProgress structs containing the enrollment ID and total completed lessons.
// - error: An error object if the operation fails, otherwise nil.
func (lp *LearningProgress) GetTotalCompletedProgress(ctx context.Context) ([]CompletedProgress, error) {
	query := `SELECT enrollment_id, COUNT(*) FROM learning_progress WHERE status = true IS NULL GROUP BY enrollment_id`

	var completedProgress []CompletedProgress
	err := db.SelectContext(ctx, &completedProgress, query)
	if err != nil {
		return nil, err
	}

	return completedProgress, nil
}

// ListLearningProgress lists all learning progress records for a specific enrollment.
// It retrieves records from the learning_progress table where the enrollment_id matches
// and the deleted_at field is NULL.
//
// Parameters:
// - ctx: The context for the database operation.
// - enrollmentID: The ID of the enrollment.
//
// Returns:
// - []LearningProgress: A slice of LearningProgress structs containing the progress records.
// - error: An error object if the operation fails, otherwise nil.
func (lp *LearningProgress) ListLearningProgress(ctx context.Context, enrollmentID uint32) ([]LearningProgress, error) {
	query := `SELECT id, enrollment_id, lesson_id, status, last_accessed, completed_at FROM learning_progress WHERE enrollment_id = $1`

	var learningProgress []LearningProgress
	err := db.SelectContext(ctx, &learningProgress, query, enrollmentID)
	if err != nil {
		return nil, err
	}

	return learningProgress, nil
}

// UpdateLastAccessed updates the last accessed timestamp for a specific lesson in a specific enrollment.
// It updates the last_accessed field in the learning_progress table where the enrollment_id and lesson_id match
// and the deleted_at field is NULL.
//
// Parameters:
// - ctx: The context for the database operation.
// - enrollmentID: The ID of the enrollment.
// - lessonID: The ID of the lesson.
// - lastAccessed: The timestamp when the lesson was last accessed.
//
// Returns:
// - error: An error object if the operation fails, otherwise nil.
func (lp *LearningProgress) UpdateLastAccessed(ctx context.Context, enrollmentID, lessonID uint32, lastAccessed time.Time) error {
	query := `UPDATE learning_progress SET last_accessed = $1 WHERE enrollment_id = $2 AND lesson_id = $3`

	_, err := db.ExecContext(ctx, query, lastAccessed, enrollmentID, lessonID)
	if err != nil {
		return err
	}

	return nil
}
