/*
	package controller
	レイヤーはInterface Adapters層。
	受け取ったリクエストを取り出し、それぞれのレイヤーで利用できる形に変換した上でusecaseの処理を呼び出す。
	その後usecase層からの戻り値をpresentersに渡す。
*/
package controller

import (
	"net/http"
	"src/domain/entity"
	"src/interfaces/presenters"
	"src/usecase/interactor"

	"github.com/labstack/echo/v4"
)

// User ユーザコントローラ
type User struct {
	Interactor interactor.User
	Presenters presenters.UserOutputPort
}

// GetAllUser 取得
func (controller *User) GetAllUser(c echo.Context) error {

	res := controller.Interactor.GetAllUser()
	return controller.Presenters.OutputAllUser(c, res)

}

// GetUser 取得
func (controller *User) GetUser(c echo.Context) error {

	id := c.Param("id")

	res, err := controller.Interactor.GetInfo(id)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return controller.Presenters.OutputUser(c, res)
}

func (controller *User) ShowCreateUser(c echo.Context) error {
	return c.Render(http.StatusOK, "user_create", "")
}

// CreateUser ※このメソッドは未完成状態！！
// 現状はあくまでもコンパイルを通すためだけの処理が記載
func (controller *User) CreateUser(c echo.Context) error {
	res := entity.User{}
	return controller.Presenters.OutputCreateUserDone(c, res)
}
