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
	lesson      *handler.LessonHandler
	payments    *handler.PaymentsHandler
	enrollments *handler.EnrollmentHandler
}

func NewHandlers(
	userHandler *handler.UserHandler,
	courseHandler *handler.CourseHandler,
	lessonHandler *handler.LessonHandler,
	paymentsHandler *handler.PaymentsHandler,
	enrollmentHandler *handler.EnrollmentHandler,
) *Handlers {
	return &Handlers{
		user:        userHandler,
		course:      courseHandler,
		lesson:      lessonHandler,
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

	// enrollment routes
	e.POST("/enrollment", handlers.enrollments.CreateEnrollment, echoJWT.WithConfig(jwtConfig))
	e.GET("/enrollment/:id", handlers.enrollments.GetEnrollmentsByStudentId, echoJWT.WithConfig(jwtConfig))
	e.DELETE("/enrollment/:id", handlers.enrollments.DeleteEnrollmentById, echoJWT.WithConfig(jwtConfig))
	e.GET("/enrollment/student/:student_id", handlers.enrollments.GetEnrollmentsByStudentId, echoJWT.WithConfig(jwtConfig))

	// payment routes

	e.GET("/payments/:id", handlers.payments.GetPaymentByEnrollmentId, echoJWT.WithConfig(jwtConfig))
	
}
