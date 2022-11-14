/*
	package router
	ユーザからのリクエストに対するハンドラの登録を行う。
	コントローラ単位でファイルを分ける。
	チームでの開発にてrouter.goを編集するリスクを少しでも下げるため。
*/
package router

import (
	"src/infra/database"
	"src/infra/di"

	"github.com/labstack/echo/v4"
)

func setUserRoutes(e *echo.Echo) {

	userController := di.NewUserController(database.NewSQLHandler())

	e.GET("/users", func(c echo.Context) error {
		return userController.GetAllUser(c)
	})
	e.GET("/user/:id", func(c echo.Context) error {
		return userController.GetUser(c)
	})
	e.GET("/user/create", func(c echo.Context) error {
		return userController.ShowCreateUser(c)
	})
}
