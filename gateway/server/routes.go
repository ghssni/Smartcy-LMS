package server

import (
	"gateway-service/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	course *handler.CourseHandler
}

func NewHandlers(courseHandler *handler.CourseHandler) *Handlers {
	return &Handlers{
		course: courseHandler,
	}
}

func Routes(e *echo.Echo, handlers *Handlers) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Register Swagger route
	//e.GET("/swagger/*", echoSwagger.WrapHandler)
	//
	//e.POST("/users/register", handlers.user.Register)
	//e.POST("/users/login", handlers.user.Login)
	//
	//jwtConfig := echoJWT.Config{
	//	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	//		return new(middlewares.JWTCustomClaims)
	//	},
	//	SigningKey: []byte(config.Viper.GetString("JWT_SECRET")),
	//}

	// Book Routes
	e.POST("/course", handlers.course.CreateCourse)
	e.GET("/courses", handlers.course.GetAllCourses)
	e.GET("/course/:id", handlers.course.GetCourseByID)
	e.PUT("/course/:id", handlers.course.UpdateCourse)
	e.DELETE("/course/:id", handlers.course.DeleteCourse)

}
