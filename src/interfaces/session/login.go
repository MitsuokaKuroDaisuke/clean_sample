package session

import (
	"github.com/labstack/echo/v4"
)

type Login interface {
	NewSession(c echo.Context) error
	GetSession(c echo.Context) (bool, string)
	DeleteSession(c echo.Context) error
}
