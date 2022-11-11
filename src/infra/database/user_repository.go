/*
	package database
	レイヤーはinfra層で、データの永続化を行う。
*/
package database

import (
	"src/domain/entity"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository ユーザリポジトリの実装
type UserRepository struct {
	SQLHandler
}

// SelectAll 取得
func (rep *UserRepository) SelectAll() []entity.User {
	users := []entity.User{}
	rep.db.Order("id desc").Find(&users)
	return users
}

// FindByID 取得
func (rep *UserRepository) FindByID(id int) entity.User {
	user := entity.User{}
	rep.db.Where("id = ?", id).First(&user)
	return user
}

func (rep *UserRepository) DoAuth(prm entity.User) (entity.User, error) {
	user := entity.User{}
	rep.db.Where("email = ?", prm.Email).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(prm.Password))
	return user, err
}
