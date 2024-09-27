package server

import (
	"gateway-service/server/handler"
	"gateway-service/server/middlewares"
	"github.com/golang-jwt/jwt/v5"
	echoJWT "github.com/labstack/echo-jwt/v4"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	user        *handler.UserHandler
	course      *handler.CourseHandler
	review      *handler.ReviewHandler
	lesson      *handler.LessonHandler
	lp          *handler.LearningProgressHandler
	payments    *handler.PaymentsHandler
	enrollments *handler.EnrollmentHandler
}

func NewHandlers(
	userHandler *handler.UserHandler,
	courseHandler *handler.CourseHandler,
	reviewHandler *handler.ReviewHandler,
	lessonHandler *handler.LessonHandler,
	lpHandler *handler.LearningProgressHandler,
	paymentsHandler *handler.PaymentsHandler,
	enrollmentHandler *handler.EnrollmentHandler,
) *Handlers {
	return &Handlers{
		user:        userHandler,
		course:      courseHandler,
		review:      reviewHandler,
		lesson:      lessonHandler,
		lp:          lpHandler,
		payments:    paymentsHandler,
		enrollments: enrollmentHandler,
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
	e.POST("/user/reset-password", handlers.user.NewPassword)
	e.POST("/user/forgot-password", handlers.user.ForgotPassword)
	//

	//webhook
	e.POST("/webhook", handlers.payments.HandleWebhook)

	// Get Secret Key
	jwtConfig := echoJWT.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middlewares.JWTCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	// user routes
	e.GET("/user/profile/:id", handlers.user.GetUserProfile, echoJWT.WithConfig(jwtConfig))
	e.PUT("/user/profile/:id", handlers.user.UpdateUserProfile, echoJWT.WithConfig(jwtConfig))

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

	// enrollment routes
	e.POST("/enrollment", handlers.enrollments.CreateEnrollment, echoJWT.WithConfig(jwtConfig))
	e.GET("/enrollment/:id", handlers.enrollments.GetEnrollmentsByStudentId, echoJWT.WithConfig(jwtConfig))
	e.DELETE("/enrollment/:id", handlers.enrollments.DeleteEnrollmentById, echoJWT.WithConfig(jwtConfig))
	e.GET("/enrollment/student/:student_id", handlers.enrollments.GetEnrollmentsByStudentId, echoJWT.WithConfig(jwtConfig))

	// payment routes

	e.GET("/payments/:id", handlers.payments.GetPaymentByEnrollmentId, echoJWT.WithConfig(jwtConfig))

}
