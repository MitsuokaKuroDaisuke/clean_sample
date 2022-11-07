package router

import (
	"src/infra"
	"src/infra/database"

	"github.com/labstack/echo/v4"
)

func setMemberRoutes(e *echo.Echo) {

	memberController := infra.NewMemberController(database.NewSQLHandler())

	e.GET("/member/:id", func(c echo.Context) error {
		return memberController.GetMember(c)
	})
	e.GET("/original/member/:id", func(c echo.Context) error {
		return memberController.GetOriginalMember(c)
	})
}
