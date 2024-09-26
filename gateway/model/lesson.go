package model

type Lesson struct {
	ID         uint32 `json:"id,omitempty" db:"id"`
	Title      string `json:"title" db:"title"`
	ContentURL string `json:"content_url" db:"content_url"`
	LessonType string `json:"lesson_type" db:"lesson_type"`
	Sequence   uint32 `json:"sequence" db:"sequence"`
	CourseID   uint32 `json:"course_id" db:"course_id"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
	DeletedAt  string `json:"deleted_at,omitempty" db:"deleted_at"`
}
