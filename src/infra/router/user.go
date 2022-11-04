package router

import (
	"www/infra"
	"www/infra/database"

	"github.com/labstack/echo/v4"
)

func setUserRoutes(e *echo.Echo) {

	userController := infra.NewUserController(database.NewSQLHandler())

	e.GET("/users", func(c echo.Context) error {
		return userController.GetAllUser(c)
	})
	e.GET("/user/:id", func(c echo.Context) error {
		return userController.GetUser(c)
	})
	e.GET("/user/create", func(c echo.Context) error {
		return userController.ShowCreateUser(c)
	})
	e.POST("/user/create", func(c echo.Context) error {
		return userController.CreateUser(c)
	})
}
