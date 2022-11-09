/*
	package interactor
	ビジネスロジック。domain層意外に依存してはならない。
*/
package interactor

import (
	"src/domain/entity"
	"strconv"
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

// Create ※このメソッドは未完成状態！！
// 現状はあくまでもコンパイルを通すためだけの処理が記載
func (interactor *User) Create(user entity.User) entity.User {
	res := entity.User{}
	return res
}
