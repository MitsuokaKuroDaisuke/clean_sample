package presenters

import (
	"fmt"
	"net/http"
	"time"
	"www/common"
	"www/domain"

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
func (pr *UserOutputPort) OutputAllUser(c echo.Context, users []domain.User) error {

	res := []APIUser{}
	for _, u := range users {
		r := new(APIUser)
		common.CopyStruct(u, r)
		res = append(res, *r)
	}

	return c.JSON(http.StatusOK, res)
}

// OutputUser 指定したユーザの取得
func (pr *UserOutputPort) OutputUser(c echo.Context, user domain.User) error {

	data := struct {
		Users domain.User
	}{
		Users: user,
	}
	return c.Render(http.StatusOK, "user_list", data)
}

func (pr *UserOutputPort) OutputCreateUserDone(c echo.Context, user domain.User) error {

	if user.ID == 0 {
		return fmt.Errorf("ユーザが作成されていません。")
	}
	data := struct {
		User domain.User
	}{
		User: user,
	}
	return c.Render(http.StatusOK, "user_create_done", data)

}
