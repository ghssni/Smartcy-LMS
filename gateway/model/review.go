package model

type Review struct {
	ID         uint32 `json:"id,omitempty"`
	CourseID   uint32 `json:"course_id" validate:"required"`
	StudentID  string `json:"student_id" validate:"required"`
	Rating     uint32 `json:"rating" validate:"required"`
	ReviewText string `json:"review_text"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
