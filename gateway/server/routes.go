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
	lesson *handler.LessonHandler
}

func NewHandlers(
	userHandler *handler.UserHandler,
	courseHandler *handler.CourseHandler,
	lessonHandler *handler.LessonHandler,
) *Handlers {
	return &Handlers{
		user:   userHandler,
		course: courseHandler,
		lesson: lessonHandler,
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

}
