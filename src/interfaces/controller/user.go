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
	"src/interfaces/session"
	"src/usecase/interactor"

	"github.com/labstack/echo/v4"
)

// User ユーザコントローラ
type User struct {
	Interactor   interactor.User
	Presenters   presenters.UserOutputPort
	LoginSession session.Login
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

func (controller *User) ShowLoginUser(c echo.Context) error {
	return c.Render(http.StatusOK, "login", "")
}

func (controller *User) LoginUser(c echo.Context) error {
	prm := entity.User{}
	c.Bind(&prm)
	err := controller.Interactor.Login(prm)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	controller.LoginSession.NewSession(c)
	return c.Redirect(http.StatusFound, "/user/home")
}

func (controller *User) ShowHome(c echo.Context) error {
	auth, name := controller.LoginSession.GetSession(c)
	if !auth {
		return c.Redirect(http.StatusFound, "/user/login")
	}

	data := struct {
		Name string
	}{
		Name: name,
	}

	return c.Render(http.StatusOK, "home", data)
}

func (controller *User) LogoutUser(c echo.Context) error {
	controller.LoginSession.DeleteSession(c)
	return c.Render(http.StatusOK, "home", "")
}
