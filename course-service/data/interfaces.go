package data

type CourseInterfaces interface {
	CreateCourse(course *Course) (uint32, error)
	GetCourseByID(courseID uint32) (*Course, error)
	GetCourseByInstructorID(instructorID string) ([]Course, error)
	GetCourseByCategory(category string) ([]Course, error)
	GetAllCourses() ([]Course, error)
	UpdateCourse(courseID uint32, course *Course) error
	DeleteCourse(courseID uint32) error
	CheckCourseByID(courseID uint32) (bool, error)
}
