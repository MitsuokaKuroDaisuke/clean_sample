package interactor

import (
	"strconv"
	"www/domain/entity"
)

// User Userビジネスロジック
type User struct {
	UserRepository entity.UserRepository
}

// GetAllUser 取得
func (interactor *User) GetAllUser() []entity.User {
	users := interactor.UserRepository.SelectAll()
	return users
}

// GetInfo 取得
func (interactor *User) GetInfo(id string) (entity.User, error) {
	uid, err := strconv.Atoi(id)
	res := interactor.UserRepository.FindByID(uid)
	return res, err
}
