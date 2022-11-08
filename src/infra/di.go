package infra

import (
	"src/infra/database"
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
		Interactor: interactor,
		Presenters: presenters.UserOutputPort{},
	}
}

func NewMemberController(sqlHandler database.SQLHandler) *controller.Member {

	interactor := interactor.Member{
		MemberRepository: &database.MemberRepository{
			SQLHandler: sqlHandler,
		},
		MemberDetailRepository: &database.MemberDetailRepository{
			SQLHandler: sqlHandler,
		},
		MemberMatchRepository: &database.MemberMatchRepository{
			SQLHandler: sqlHandler,
		},
		OriginalMemberRepository: &database.OriginalMemberRepository{
			SQLHandler: sqlHandler,
		},
		Transaction: &database.Transaction{
			SQLHandler: sqlHandler,
		},
	}

	return &controller.Member{
		Interactor: interactor,
		Presenters: presenters.MemberOutputPort{},
	}
}
