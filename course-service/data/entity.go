package data

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func New(dbPool *sqlx.DB) *Models {
	db = dbPool

	return &Models{
		Course:           &Course{},
		Lesson:           &Lesson{},
		Review:           &Review{},
		LearningProgress: &LearningProgress{},
	}
}

type Models struct {
	Course           CourseInterfaces
	Lesson           LessonInterfaces
	Review           ReviewInterfaces
	LearningProgress LearningProgressInterfaces
}
