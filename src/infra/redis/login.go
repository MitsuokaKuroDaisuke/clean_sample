package redis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v8"
)

type LoginSession struct {
	Sess SessionHandler
}

func (ls LoginSession) getSession(c echo.Context) (*sessions.Session, error) {
	store, err := redisstore.NewRedisStore(context.Background(), ls.Sess.client)
	if err != nil {
		return nil, err
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		MaxAge:   600,
		HttpOnly: true,
	})
	session, err := store.Get(c.Request(), "my-investment")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (ls LoginSession) NewSession(c echo.Context) error {
	session, err := ls.getSession(c)
	session.Values["username"] = "u.Name"                            //ログインユーザ名を付与
	session.Values["auth"] = true                                    // ログイン有無の確認用
	if err := sessions.Save(c.Request(), c.Response()); err != nil { // Session情報の保存
		return fmt.Errorf("エラー")
	}
	return err
}

func (ls LoginSession) GetSession(c echo.Context) (bool, string) {
	session, _ := ls.getSession(c)
	if session.Values["auth"] == true {
		return true, session.Values["username"].(string)
	} else {
		return false, ""
	}
}

func (ls LoginSession) DeleteSession(c echo.Context) error {
	session, _ := ls.getSession(c)
	// ログアクト
	session.Values["auth"] = false
	// セッション削除
	session.Options.MaxAge = -1
	if err := sessions.Save(c.Request(), c.Response()); err != nil {
		return fmt.Errorf("ログアウト失敗")
	}

	return c.Redirect(http.StatusFound, "/user/login")
}
