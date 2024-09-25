package data

import (
	"errors"
	"golang.org/x/net/context"
	"time"
)

// Course represents the courses table in PostgreSQL
type Course struct {
	ID           uint32     `json:"id" db:"id"`
	Title        string     `json:"title" db:"title"`
	Description  string     `json:"description" db:"description"`
	Price        float64    `json:"price" db:"price"`
	ThumbnailURL string     `json:"thumbnail_url" db:"thumbnail_url"`
	InstructorID string     `json:"instructor_id" db:"instructor_id"`
	Category     string     `json:"category" db:"category"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`

	// Additional fields
	AverageRating float32 `db:"average_rating" json:"average_rating"` // Calculated
	TotalReviews  uint32  `db:"total_reviews" json:"total_reviews"`   // Calculated
}

func (c *Course) CreateCourse(ctx context.Context, course *Course, createdAt, updatedAt time.Time) (uint32, error) {
	sqlStatement := `INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var lastInsertID uint32
	err := db.QueryRowContext(ctx, sqlStatement, course.Title, course.Description, course.Price, course.ThumbnailURL, course.InstructorID, course.Category, createdAt, updatedAt).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (c *Course) GetCourseByID(ctx context.Context, courseID uint32) (*Course, error) {
	var course Course

	err := db.GetContext(ctx, &course, `SELECT
    c.id,
    c.title,
    c.description,
    c.price,
    c.thumbnail_url,
    c.instructor_id,
    c.category,
    c.updated_at,
    COALESCE(AVG(r.rating), 0) AS average_rating,
    COUNT(r.id) AS total_reviews
FROM courses c
         LEFT JOIN reviews r ON c.id = r.course_id
WHERE c.id = $1 AND c.deleted_at IS NULL
GROUP BY c.id;
`, courseID)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *Course) GetCourseByInstructorID(ctx context.Context, instructorID string) ([]Course, error) {
	sqlStatement := `
        SELECT 
            c.id, 
            c.title, 
            c.description,
            c.price,
            c.thumbnail_url,
            c.instructor_id,
            c.category,
            c.updated_at, 
            COALESCE(AVG(r.rating), 0) AS average_rating, 
            COUNT(r.id) AS total_reviews
        FROM courses c
        LEFT JOIN reviews r ON c.id = r.course_id
        WHERE c.instructor_id = $1 AND c.deleted_at IS NULL
        GROUP BY c.id
        ORDER BY total_reviews DESC, average_rating DESC;
    `

	var courses []Course
	err := db.SelectContext(ctx, &courses, sqlStatement, instructorID)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Course) GetCourseByCategory(ctx context.Context, category string) ([]Course, error) {

	sqlStatement := `
        SELECT 
            c.id, 
            c.title, 
            c.description, 
            c.price,
            c.thumbnail_url,
            c.instructor_id,
            c.category,
            c.updated_at, 
            COALESCE(AVG(r.rating), 0) AS average_rating, 
            COUNT(r.id) AS total_reviews
        FROM courses c
        LEFT JOIN reviews r ON c.id = r.course_id
        WHERE c.category = $1  AND c.deleted_at IS NULL
        GROUP BY c.id
        ORDER BY total_reviews DESC, average_rating DESC;
    `

	var courses []Course
	err := db.SelectContext(ctx, &courses, sqlStatement, category)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Course) GetAllCourses(ctx context.Context) ([]Course, error) {
	sqlStatement := `
        SELECT 
            c.id, 
            c.title, 
            c.description, 
            c.price,
            c.thumbnail_url,
            c.instructor_id,
            c.category,
            c.updated_at, 
            COALESCE(AVG(r.rating), 0) AS average_rating, 
            COUNT(r.id) AS total_reviews
        FROM courses c
        LEFT JOIN reviews r ON c.id = r.course_id
        WHERE c.deleted_at IS NULL
        GROUP BY c.id
        ORDER BY average_rating DESC, total_reviews DESC;
    `

	var courses []Course
	err := db.SelectContext(ctx, &courses, sqlStatement)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Course) UpdateCourse(ctx context.Context, course *Course, updatedAt time.Time) error {
	sqlStatement := `UPDATE courses SET title = $1, description = $2, price = $3, thumbnail_url = $4, instructor_id = $5, category = $6, updated_at = $7 WHERE id = $8 AND deleted_at IS NULL`

	result, err := db.ExecContext(ctx, sqlStatement, course.Title, course.Description, course.Price, course.ThumbnailURL, course.InstructorID, course.Category, updatedAt, course.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("failed to update course")
	}

	return nil
}

func (c *Course) DeleteCourse(ctx context.Context, courseID uint32, deletedAt time.Time) error {
	sqlStatement := `UPDATE courses SET deleted_at = $1 WHERE id = $2`

	result, err := db.ExecContext(ctx, sqlStatement, deletedAt, courseID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("failed to delete course")
	}

	return nil
}

func (c *Course) CheckCourseByID(ctx context.Context, courseID uint32) (bool, error) {
	sqlStatement := `SELECT 1 FROM courses WHERE id = $1`

	var count int
	err := db.GetContext(ctx, &count, sqlStatement, courseID)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
