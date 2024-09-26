package model

type Lesson struct {
	ID         uint32 `json:"id,omitempty"`
	Title      string `json:"title" validate:"required"`
	ContentURL string `json:"content_url" validate:"required"`
	LessonType string `json:"lesson_type" validate:"required"`
	Sequence   uint32 `json:"sequence" validate:"required"`
	CourseID   uint32 `json:"course_id" validate:"required"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}
