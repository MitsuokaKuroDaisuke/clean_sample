package controller

import (
	"net/http"
	"src/interfaces/presenters"
	"src/usecase/interactor"

	"github.com/labstack/echo/v4"
)

type Member struct {
	Interactor interactor.Member
	Presenters presenters.MemberOutputPort
}

func (controller *Member) GetMember(c echo.Context) error {

	id := c.Param("id")

	res, err := controller.Interactor.GetMember(id)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return controller.Presenters.Member(c, res)
}

func (controller *Member) GetOriginalMember(c echo.Context) error {

	id := c.Param("id")

	res, err := controller.Interactor.GetOriginalMember(id)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return controller.Presenters.OriginalMember(c, res)
}
