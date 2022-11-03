package infra

import (
	"www/infra/database"
	"www/interfaces/controller"
	"www/interfaces/presenters"
	"www/usecase"
)

// NewUserController ユーザコントローラ取得
func NewUserController(sqlHandler database.SQLHandler) *controller.UserController {

	userRepository := &database.UserRepository{
		SQLHandler: sqlHandler,
	}

	interactor := usecase.UserInteractor{
		UserRepository: userRepository,
	}

	return &controller.UserController{
		Interactor: interactor,
		Presenters: presenters.UserOutputPort{},
	}
}
