package model

// Course represents the courses table in PostgreSQL
type Course struct {
	ID           uint32  `json:"id,omitempty"`
	Title        string  `json:"title" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	ThumbnailURL string  `json:"thumbnail_url" validate:"required"`
	InstructorID string  `json:"instructor_id" validate:"required"`
	Category     string  `json:"category" validate:"required"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

type CourseWithReview struct {
	ID           uint32  `json:"id,omitempty"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ThumbnailURL string  `json:"thumbnail_url"`
	InstructorID string  `json:"instructor_id"`
	Category     string  `json:"category"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
	// Additional fields
	AverageRating float32 `json:"average_rating"` // Calculated
	TotalReviews  uint32  `json:"total_reviews"`  // Calculated
}
