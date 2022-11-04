package controller

import (
	"net/http"
	"www/domain"
	"www/interfaces/presenters"
	"www/usecase"

	"github.com/labstack/echo/v4"
)

// UserController ユーザコントローラ
type UserController struct {
	Interactor usecase.UserInteractor
	Presenters presenters.UserOutputPort
}

// GetAllUser 取得
func (controller *UserController) GetAllUser(c echo.Context) error {

	res := controller.Interactor.GetAllUser()
	return controller.Presenters.OutputAllUser(c, res)

}

// GetUser 取得
func (controller *UserController) GetUser(c echo.Context) error {

	id := c.Param("id")

	res := controller.Interactor.GetInfo(id)
	return controller.Presenters.OutputUser(c, res)
}

func (controller *UserController) ShowCreateUser(c echo.Context) error {
	return c.Render(http.StatusOK, "user_create", "")
}

func (controller *UserController) CreateUser(c echo.Context) error {
	res := domain.User{}
	return controller.Presenters.OutputCreateUserDone(c, res)
}
