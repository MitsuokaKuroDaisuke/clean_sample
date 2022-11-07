package database

import "src/domain/entity"

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
