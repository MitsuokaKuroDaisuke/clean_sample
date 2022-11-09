package controller

import (
	"net/url"
	"src/infra/database"
	"src/interfaces/controller"
	"src/interfaces/presenters"
	"src/usecase/interactor"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUserController_CreateUser(t *testing.T) {
	type fields struct {
		Interactor interactor.User
		Presenters presenters.UserOutputPort
	}
	type args struct {
		c echo.Context
	}

	f := fields{
		Interactor: interactor.User{
			UserRepository: &database.UserRepository{
				SQLHandler: database.NewTestSQLHandler(),
			},
		},
		Presenters: presenters.UserOutputPort{},
	}

	// echo.Context作成
	form := make(url.Values)
	form.Set("name", "CreateTest")
	form.Set("email", "mail@example.com")

	a := args{
		c: initContext(form),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "登録完了して完了画面まで表示できるかのテスト",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &controller.User{
				Interactor: tt.fields.Interactor,
				Presenters: tt.fields.Presenters,
			}
			if err := controller.CreateUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserController.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
