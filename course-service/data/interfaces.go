package data

import (
	"context"
	"time"
)

type CourseInterfaces interface {
	CreateCourse(ctx context.Context, course *Course, createdAt, updatedAt time.Time) (uint32, error)
	GetCourseByID(ctx context.Context, courseID uint32) (*Course, error)
	GetCourseByInstructorID(ctx context.Context, instructorID string) ([]Course, error)
	GetCourseByCategory(ctx context.Context, category string) ([]Course, error)
	GetAllCourses(ctx context.Context) ([]Course, error)
	UpdateCourse(ctx context.Context, course *Course, updatedAt time.Time) error
	DeleteCourse(ctx context.Context, courseID uint32, deletedAt time.Time) error
	CheckCourseByID(ctx context.Context, courseID uint32) (bool, error)
}
type LessonInterfaces interface {
	CreateLesson(ctx context.Context, lesson *Lesson, createdAt, updatedAt time.Time) (uint32, error)
	GetLesson(ctx context.Context, lessonID uint32) (*Lesson, error)
	GetLessonBySequence(ctx context.Context, sequence, courseID uint32) (*Lesson, error)
	GetAllLessons(ctx context.Context, courseID uint32) ([]LessonSummary, error)
	UpdateLesson(ctx context.Context, lesson *Lesson, updatedAt time.Time) error
	DeleteLesson(ctx context.Context, lessonID uint32, deletedAt time.Time) error
	DeleteLessonByCourse(ctx context.Context, courseID uint32, deletedAt time.Time) error
	SearchLessonByTitle(ctx context.Context, courseID uint32, title string) ([]Lesson, error)
	SearchLessonByType(ctx context.Context, courseID uint32, lessonType string) ([]Lesson, error)
	GetTotalLessonsByCourseID(ctx context.Context, courseID uint32) (uint32, error)
	GetTotalLessonsByCourseIDAndType(ctx context.Context, courseID uint32, lessonType string) (uint32, error)
}
type ReviewInterfaces interface {
	CreateReview(ctx context.Context, review *Review, createdAt, updatedAt time.Time) (uint32, error)
	GetReview(ctx context.Context, reviewID uint32) (*Review, error)
	GetReviewsByStudent(ctx context.Context, courseID uint32, studentID string) ([]Review, error)
	GetTotalReviewsByCourse(ctx context.Context, courseID uint32) (uint32, error)
	GetAverageRatingByCourse(ctx context.Context, courseID uint32) (float32, error)
	ListReviews(ctx context.Context, courseID uint32) ([]Review, error)
	UpdateReview(ctx context.Context, review *Review, updatedAt time.Time) error
	DeleteReview(ctx context.Context, reviewID uint32, deletedAt time.Time) error
}
type LearningProgressInterfaces interface {
	MarkLessonAsCompleted(ctx context.Context, userID string, enrollmentID, lessonID uint32, lastAccessed, completedAt time.Time) error
	ResetLessonMark(ctx context.Context, userID string, enrollmentID, lessonID uint32) error
	ResetAllLessonMarks(ctx context.Context, userID string, enrollmentID uint32) error
	GetTotalCompletedLessons(ctx context.Context, userID string, enrollmentID uint32) (*CompletedProgress, error)
	GetTotalCompletedProgress(ctx context.Context, userID string) ([]CompletedProgress, error)
	ListLearningProgress(ctx context.Context, userID string, enrollmentID uint32) ([]LearningProgress, error)
	UpdateLastAccessed(ctx context.Context, userID string, enrollmentID, lessonID uint32, lastAccessed time.Time) error
}
