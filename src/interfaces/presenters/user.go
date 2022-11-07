package presenters

import (
	"net/http"
	"src/domain/entity"
	"src/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

type UserOutputPort struct {
}

// APIUser API返却用ユーザ
type APIUser struct {
	ID        int       `json:"id"`
	Email     string    `json:"e_mail"`
	Name      string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

// OutputAllUser ユーザ全取得の表示
func (pr *UserOutputPort) OutputAllUser(c echo.Context, users []entity.User) error {

	res := []APIUser{}
	for _, u := range users {
		r := new(APIUser)
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
