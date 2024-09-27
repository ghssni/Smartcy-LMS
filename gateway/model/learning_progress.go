package model

type LearningProgress struct {
	ID           uint32 `json:"id,omitempty""`
	EnrollmentID uint32 `json:"enrollment_id"`
	LessonID     uint32 `json:"lesson_id"`
	Status       bool   `json:"status"`
	LastAccessed string `json:"last_accessed"`
	CompletedAt  string `json:"completed_at"`
}

type CompletedProgress struct {
	EnrollmentID   uint32 `db:"enrollment_id"`
	TotalCompleted uint32 `db:"total_completed"`
}

type LPRequest struct {
	EnrollmentID uint32 `json:"enrollment_id" validate:"required"`
	LessonID     uint32 `json:"lesson_id" validate:"required"`
}

type CompletedProgressResponse struct {
	EnrollmentID   uint32 `json:"enrollment_id"`
	TotalCompleted uint32 `json:"total_completed"`
}
