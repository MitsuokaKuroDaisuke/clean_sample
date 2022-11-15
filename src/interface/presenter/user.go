/*
	package presenter
	レイヤーはInterface Adapters層。
	controllerから受け取った値をユーザに返却する形に変換する。
*/
package presenter

import (
	"net/http"
	"src/domain/entity"
	"src/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type UserOutputPort struct {
}

// OutputAllUser ユーザ全取得の表示
func (pr *UserOutputPort) OutputAllUser(c echo.Context, users []entity.User) error {

	type response struct {
		ID        int       `json:"id"`
		Email     string    `json:"e_mail"`
		Name      string    `json:"user_name"`
		CreatedAt time.Time `json:"created_at"`
	}
	res := []response{}
	for _, u := range users {
		r := new(response)
		usecase.CopyStruct(u, r)
		res = append(res, *r)
	}

	return c.JSON(http.StatusOK, res)
}

// OutputUser 指定したユーザの取得
func (pr *UserOutputPort) OutputUser(c echo.Context, user entity.User) error {

	data := struct {
		Users entity.User
	}{
		Users: user,
	}
	return c.Render(http.StatusOK, "user_list", data)

}
