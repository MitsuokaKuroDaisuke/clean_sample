package usecase

import (
	"strconv"
	"www/domain/entity"
)

// UserInteractor Userビジネスロジック
type UserInteractor struct {
	UserRepository entity.UserRepository
}

// GetAllUser 取得
func (interactor *UserInteractor) GetAllUser() []entity.User {
	users := interactor.UserRepository.SelectAll()
	return users
}

// GetInfo 取得
func (interactor *UserInteractor) GetInfo(id string) (entity.User, error) {
	uid, err := strconv.Atoi(id)
	res := interactor.UserRepository.FindByID(uid)
	return res, err
}
