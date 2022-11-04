package controller

import (
	"net/url"
	"testing"
	"www/infra/database"
	"www/interfaces/presenters"
	"www/usecase"

	"github.com/labstack/echo/v4"
)

func TestUserController_CreateUser(t *testing.T) {
	type fields struct {
		Interactor usecase.UserInteractor
		Presenters presenters.UserOutputPort
	}
	type args struct {
		c echo.Context
	}

	f := fields{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: database.NewTestSQLHandler(),
			},
		},
		Presenters: presenters.UserOutputPort{},
	}

	// echo.Context作成
	form := make(url.Values)
	form.Set("name", "CreateTest")
	form.Set("email", "mail@mai.jj")

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
			name:    "test1",
			fields:  f,
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &UserController{
				Interactor: tt.fields.Interactor,
				Presenters: tt.fields.Presenters,
			}
			if err := controller.CreateUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserController.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
