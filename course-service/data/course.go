package data

import (
	"log"
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
}

func (c *Course) CreateCourse(course *Course) (uint32, error) {
	createdAt := time.Now()
	updatedAt := time.Now()
	sqlStatement := `INSERT INTO courses (title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var lastInsertID uint32
	err := db.QueryRow(sqlStatement, course.Title, course.Description, course.Price, course.ThumbnailURL, course.InstructorID, course.Category, createdAt, updatedAt).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (c *Course) GetCourseByID(courseID uint32) (*Course, error) {
	var course Course

	err := db.Get(&course, `SELECT id, title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at, deleted_at FROM courses WHERE id = $1`, courseID)
	if err != nil {
		log.Printf("Error while getting course by ID: %v", err)
		return nil, err
	}

	return &course, nil
}

func (c *Course) GetCourseByInstructorID(instructorID string) ([]Course, error) {
	sqlStatement := `SELECT id, title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at, deleted_at FROM courses WHERE instructor_id = $1`

	var courses []Course
	err := db.Select(&courses, sqlStatement, instructorID)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Course) GetCourseByCategory(category string) ([]Course, error) {
	sqlStatement := `SELECT id, title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at, deleted_at FROM courses WHERE category = $1`

	var courses []Course
	err := db.Select(&courses, sqlStatement, category)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Course) GetAllCourses() ([]Course, error) {
	sqlStatement := `SELECT id, title, description, price, thumbnail_url, instructor_id, category, created_at, updated_at, deleted_at FROM courses`

	var courses []Course
	err := db.Select(&courses, sqlStatement)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *Course) UpdateCourse(courseID uint32, course *Course) error {
	sqlStatement := `UPDATE courses SET title = $1, description = $2, price = $3, thumbnail_url = $4, instructor_id = $5, category = $6, updated_at = NOW() WHERE id = $7`

	_, err := db.Exec(sqlStatement, course.Title, course.Description, course.Price, course.ThumbnailURL, course.InstructorID, course.Category, courseID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Course) DeleteCourse(courseID uint32) error {
	sqlStatement := `UPDATE courses SET deleted_at = NOW() WHERE id = $1`

	_, err := db.Exec(sqlStatement, courseID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Course) CheckCourseByID(courseID uint32) (bool, error) {
	sqlStatement := `SELECT 1 FROM courses WHERE id = $1`

	var count int
	err := db.Get(&count, sqlStatement, courseID)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
