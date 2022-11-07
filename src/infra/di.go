package infra

import (
	"www/infra/database"
	"www/interfaces/controller"
	"www/interfaces/presenters"
	"www/usecase/interactor"
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
		Presenters: presenters.UserOutputPort{},
	}
}

func NewMemberController(sqlHandler database.SQLHandler) *controller.Member {

	memberRepository := &database.MemberRepository{
		SQLHandler: sqlHandler,
	}

	interactor := interactor.Member{
		MemberRepository: memberRepository,
		OriginalMemberRepository: &database.OriginalMemberRepository{
			SQLHandler: sqlHandler,
		},
	}

	return &controller.Member{
		Interactor: interactor,
		Presenters: presenters.MemberOutputPort{},
	}
}
