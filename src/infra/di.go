/*
	ここで、依存性注入を行いControllerを作成する。
*/
package infra

import (
	"src/infra/database"
	"src/infra/redis"
	"src/interfaces/controller"
	"src/interfaces/presenters"
	"src/usecase/interactor"
)

// NewUserController ユーザコントローラ取得
func NewUserController(sqlHandler database.SQLHandler) *controller.User {

	userRepository := &database.UserRepository{
		SQLHandler: sqlHandler,
	}

	interactor := interactor.User{
		UserRepository: userRepository,
	}

	return &controller.User{
		Interactor:   interactor,
		Presenters:   presenters.UserOutputPort{},
		LoginSession: redis.LoginSession{Sess: redis.NewSessionHandler()},
	}
}
