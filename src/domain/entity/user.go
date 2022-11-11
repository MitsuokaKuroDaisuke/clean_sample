/*
	package entity
	レイヤーはdomain層で、DBのテーブルとORMで紐づくモデルを定義。
	Repositoryによるデータ永続化処理のインターフェースの他、モデルに対するドメインロジックを記載。
	他のレイヤーに依存する処理を行うのは禁止。
*/
package entity

import (
	"time"
)

// User 管理ユーザ
type User struct {
	ID        int       `json:"id" form:"id"`
	Email     string    `json:"email" form:"email" gorm:"type:varchar(255);not null"`
	Name      string    `json:"name" form:"name" gorm:"type:varchar(255);not null"`
	Password  string    `json:"password" form:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

// UserRepository ユーザリポジトリ
type UserRepository interface {
	SelectAll() []User
	FindByID(id int) User
	DoAuth(prm User) (User, error)
}
