package data

import (
	"context"
	"time"
)

type Lesson struct {
	ID         uint32     `json:"id,omitempty" db:"id"`
	Title      string     `json:"title" db:"title"`
	ContentURL string     `json:"content_url" db:"content_url"`
	LessonType string     `json:"lesson_type" db:"lesson_type"`
	Sequence   uint32     `json:"sequence" db:"sequence"`
	CourseID   uint32     `json:"course_id" db:"course_id"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// LessonSummary represents a summary of a lesson, containing only the essential details.
type LessonSummary struct {
	ID         uint32 `db:"id"`
	Title      string `db:"title"`
	LessonType string `db:"lesson_type"`
	Sequence   uint32 `db:"sequence"`
}

func (l *Lesson) CreateLesson(ctx context.Context, lesson *Lesson, createdAt, updatedAt time.Time) (uint32, error) {
	sqlStatement := `INSERT INTO lessons (title, content_url, lesson_type, sequence, course_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var lastInsertID uint32
	err := db.QueryRowContext(ctx, sqlStatement, lesson.Title, lesson.ContentURL, lesson.LessonType, lesson.Sequence, lesson.CourseID, createdAt, updatedAt).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (l *Lesson) GetLesson(ctx context.Context, lessonID, courseID uint32) (*Lesson, error) {
	var lesson Lesson

	// Execute the SQL query to retrieve the lesson details
	err := db.GetContext(ctx, &lesson, `SELECT id, title, content_url, lesson_type, sequence, course_id, created_at, updated_at FROM lessons WHERE id = $1 AND course_id = $2 AND deleted_at IS NULL`, lessonID, courseID)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (l *Lesson) GetLessonBySequence(ctx context.Context, sequence, courseID uint32) (*Lesson, error) {
	var lesson Lesson

	// Execute the SQL query to retrieve the lesson details
	err := db.GetContext(ctx, &lesson, `SELECT id, title, content_url, lesson_type, sequence, course_id, created_at, updated_at FROM lessons WHERE sequence = $1 AND course_id = $2 AND deleted_at IS NULL`, sequence, courseID)
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (l *Lesson) GetAllLessons(ctx context.Context, courseID uint32) ([]LessonSummary, error) {
	sqlStatement := `SELECT id, title, lesson_type, sequence 
                     FROM lessons 
                     WHERE course_id = $1 AND deleted_at IS NULL 
                     ORDER BY sequence ASC`

	var lessons []LessonSummary
	err := db.SelectContext(ctx, &lessons, sqlStatement, courseID)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func (l *Lesson) UpdateLesson(ctx context.Context, lesson *Lesson, updatedAt time.Time) error {
	sql := `UPDATE lessons SET title=$1, content_url=$2, lesson_type=$3, sequence=$4, updated_at=$5 WHERE id=$6 AND course_id=$7 AND deleted_at IS NULL`

	_, err := db.ExecContext(ctx, sql, lesson.Title, lesson.ContentURL, lesson.LessonType, lesson.Sequence, updatedAt, lesson.ID, lesson.CourseID)
	return err
}

func (l *Lesson) DeleteLesson(ctx context.Context, lessonID, courseID uint32, deletedAt time.Time) error {
	sql := `UPDATE lessons SET deleted_at=$1 WHERE id=$2 AND course_id=$3 AND deleted_at IS NULL`

	_, err := db.ExecContext(ctx, sql, deletedAt, lessonID, courseID)
	return err
}

func (l *Lesson) DeleteLessonByCourse(ctx context.Context, courseID uint32, deletedAt time.Time) error {
	sql := `UPDATE lessons SET deleted_at=$1 WHERE course_id=$2 AND deleted_at IS NULL`

	_, err := db.ExecContext(ctx, sql, deletedAt, courseID)
	return err
}

func (l *Lesson) SearchLessonByTitle(ctx context.Context, courseID uint32, title string) ([]Lesson, error) {
	sqlStatement := `SELECT id, title, content_url, lesson_type, sequence, course_id, created_at, updated_at FROM lessons WHERE course_id = $1 AND title ILIKE '%' || $2 || '%' AND (deleted_at IS NULL OR $3);`

	var lessons []Lesson
	err := db.SelectContext(ctx, &lessons, sqlStatement, courseID, title)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func (l *Lesson) GetTotalLessonsByCourseID(ctx context.Context, courseID uint32) (uint32, error) {
	var total uint32
	err := db.GetContext(ctx, &total, `SELECT COUNT(id) FROM lessons WHERE course_id = $1 AND deleted_at IS NULL`, courseID)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (l *Lesson) GetTotalLessonsByCourseIDAndType(ctx context.Context, courseID uint32, lessonType string) (uint32, error) {
	var total uint32
	err := db.GetContext(ctx, &total, `SELECT COUNT(id) FROM lessons WHERE course_id = $1 AND lesson_type = $2 AND deleted_at IS NULL`, courseID, lessonType)
	if err != nil {
		return 0, err
	}
	return total, nil
}
