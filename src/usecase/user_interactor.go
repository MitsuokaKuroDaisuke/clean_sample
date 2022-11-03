package usecase

import (
	"strconv"
	"www/domain"
)

// UserInteractor Userビジネスロジック
type UserInteractor struct {
	UserRepository domain.UserRepository
}

// GetAllUser 取得
func (interactor *UserInteractor) GetAllUser() []domain.User {
	users := interactor.UserRepository.SelectAll()
	return users
}

// GetInfo 取得
func (interactor *UserInteractor) GetInfo(id string) domain.User {
	uid, err := strconv.Atoi(id)
	if err != nil {
		uid = 0
	}
	res := interactor.UserRepository.FindByID(uid)
	return res
}
