package server

import (
	"gateway-service/server/handler"
	"gateway-service/server/middlewares"
	"github.com/golang-jwt/jwt/v5"
	echoJWT "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	user   *handler.UserHandler
	course *handler.CourseHandler
	review *handler.ReviewHandler
	lesson *handler.LessonHandler
	lp     *handler.LearningProgressHandler
}

func NewHandlers(
	userHandler *handler.UserHandler,
	courseHandler *handler.CourseHandler,
	reviewHandler *handler.ReviewHandler,
	lessonHandler *handler.LessonHandler,
	lpHandler *handler.LearningProgressHandler,
) *Handlers {
	return &Handlers{
		user:   userHandler,
		course: courseHandler,
		review: reviewHandler,
		lesson: lessonHandler,
		lp:     lpHandler,
	}
}

func Routes(e *echo.Echo, handlers *Handlers) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Register Swagger route
	//e.GET("/swagger/*", echoSwagger.WrapHandler)
	//
	e.POST("/user/register", handlers.user.Register)
	e.POST("/user/login", handlers.user.Login)
	//

	// Get Secret Key
	jwtConfig := echoJWT.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middlewares.JWTCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	// course routes
	e.POST("/course", handlers.course.CreateCourse, echoJWT.WithConfig(jwtConfig))
	e.GET("/courses", handlers.course.GetAllCourses)
	e.GET("/course/:id", handlers.course.GetCourseByID)
	e.PUT("/course/:id", handlers.course.UpdateCourse, echoJWT.WithConfig(jwtConfig))
	e.DELETE("/course/:id", handlers.course.DeleteCourse, echoJWT.WithConfig(jwtConfig))

	// lesson routes
	e.GET("course/detail/:course_id", handlers.lesson.GetAllLessons)
	e.POST("/lesson", handlers.lesson.CreateLesson, echoJWT.WithConfig(jwtConfig))
	e.GET("/lesson/s/:sequence/c/:course_id", handlers.lesson.GetLessonBySequence)
	e.GET("/lesson/id/:id", handlers.lesson.GetLesson)
	e.PUT("/lesson/id/:id", handlers.lesson.UpdateLesson, echoJWT.WithConfig(jwtConfig))
	e.DELETE("/lesson/id/:id", handlers.lesson.DeleteLesson, echoJWT.WithConfig(jwtConfig))

	// learning progress routes
	e.POST("/learning-progress/mark-completed", handlers.lp.MarkLessonAsCompleted, echoJWT.WithConfig(jwtConfig))
	e.POST("/learning-progress/reset-mark", handlers.lp.ResetLessonMark, echoJWT.WithConfig(jwtConfig))
	e.POST("/learning-progress/reset-all-marks", handlers.lp.ResetAllLessonMarks, echoJWT.WithConfig(jwtConfig))
	e.GET("/learning-progress/total-completed-lessons/:enrollment_id", handlers.lp.GetTotalCompletedLessons, echoJWT.WithConfig(jwtConfig))
	e.GET("/learning-progress/total-completed-progress/:enrollment_id", handlers.lp.GetTotalCompletedProgress, echoJWT.WithConfig(jwtConfig))
	e.GET("/learning-progress/list/:enrollment_id", handlers.lp.ListLearningProgress, echoJWT.WithConfig(jwtConfig))
	e.POST("/learning-progress/update-last-accessed", handlers.lp.UpdateLastAccessed, echoJWT.WithConfig(jwtConfig))

	// Review routes
	e.POST("/review", handlers.review.CreateReview, echoJWT.WithConfig(jwtConfig))
	e.GET("/reviews/c/:course_id", handlers.review.ListReviews)
	e.PUT("/review", handlers.review.UpdateReviewRequest, echoJWT.WithConfig(jwtConfig))
	e.DELETE("/review/:review_id", handlers.review.DeleteReview, echoJWT.WithConfig(jwtConfig))
}
