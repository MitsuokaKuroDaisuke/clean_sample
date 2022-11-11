package domain

// LoginUser ログインユーザ
type LoginUser struct {
	ID       int    `json:"id" form:"id"`
	Email    string `json:"email" form:"email" gorm:"type:varchar(255);not null"`
	Name     string `json:"name" form:"name" gorm:"type:varchar(255);not null"`
	Password string `json:"password" form:"password" gorm:"type:varchar(255);not null"`
}

// LoginUserRepository ログインユーザリポジトリ
type LoginUserRepository interface {
	SetSession(lu LoginUser) LoginUser
}
