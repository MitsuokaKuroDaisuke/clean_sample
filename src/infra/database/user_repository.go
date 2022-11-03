package database

import (
	"www/domain"
)

// UserRepository ユーザリポジトリの実装
type UserRepository struct {
	SQLHandler
}

// SelectAll 取得
func (rep *UserRepository) SelectAll() []domain.User {
	users := []domain.User{}
	rep.db.Order("id desc").Find(&users)
	return users
}

// FindByID 取得
func (rep *UserRepository) FindByID(id int) domain.User {
	user := domain.User{}
	rep.db.Where("id = ?", id).First(&user)
	return user
}
