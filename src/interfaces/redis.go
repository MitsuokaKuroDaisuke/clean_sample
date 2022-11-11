package interfaces

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v8"
)

type LoginSession struct {
}

func (ls LoginSession) getSession(c echo.Context) *sessions.Session {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	store, err := redisstore.NewRedisStore(context.Background(), client)
	if err != nil {
		log.Fatal("Failed cannot connect redis", err)
		//return err
	}
	store.KeyPrefix("session_")
	store.Options(sessions.Options{
		MaxAge:   600,
		HttpOnly: true,
	})
	session, err := store.Get(c.Request(), "my-investment")
	if err != nil {
		log.Fatal("Failed cannot get session", err)
	}
	return session
}

func (ls LoginSession) NewSession(c echo.Context) error {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("ランダムな文字作成時にエラーが発生しました。")
	}

	session := ls.getSession(c)
	session.Values["username"] = "u.Name"                            //ログインユーザ名を付与
	session.Values["auth"] = true                                    // ログイン有無の確認用
	if err := sessions.Save(c.Request(), c.Response()); err != nil { // Session情報の保存
		return fmt.Errorf("エラー")
	}
	return nil
}

func (ls LoginSession) GetSession(c echo.Context) (bool, string) {
	session := ls.getSession(c)
	if session.Values["auth"] == true {
		return true, session.Values["username"].(string)
	} else {
		return false, ""
	}
}

func (ls LoginSession) DeleteSession(c echo.Context) error {
	session := ls.getSession(c)
	// ログアクト
	session.Values["auth"] = false
	// セッション削除
	session.Options.MaxAge = -1
	if err := sessions.Save(c.Request(), c.Response()); err != nil {
		return fmt.Errorf("ログアウト失敗")
	}

	return c.Redirect(http.StatusFound, "/user/login")
}
