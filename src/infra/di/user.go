/*
	ここで、依存性注入を行いControllerを作成する。
*/
package di

import (
	"src/infra/database"
	"src/interface/controller"
	"src/interface/presenter"
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
		Interactor: interactor,
		Presenter:  presenter.UserOutputPort{},
	}
}
