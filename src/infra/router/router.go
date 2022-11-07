package router

import (
	"net/http"
	"www/infra"

	"github.com/labstack/echo/v4"
)

// Init ルーティング初期化
func Init() *echo.Echo {

	e := echo.New()
	e.Renderer = infra.InitTmp()

	setStatic(e)
	setUserRoutes(e)
	setMemberRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "/users")
	})

	return e
}

func setStatic(e *echo.Echo) {
	e.Static("/public/css/", "/var/www/public/css")
	e.Static("/public/js/", "/var/www/public/js/")
	e.Static("/public/img/", "/var/www/public/img/")
}
