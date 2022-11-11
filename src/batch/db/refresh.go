package main

import (
	"fmt"
	"os"
	"src/domain/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sTokyo", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"), "%2F")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}

	db.Migrator().DropTable(&entity.User{})
	db.AutoMigrate(&entity.User{})

	users := []entity.User{
		{Email: "d@example.jp", Name: "mitsu", Password: passwordEncrypt("password")},
		{Email: "t.tt@example.jp", Name: "test2", Password: passwordEncrypt("password02")},
		{Email: "y.ooo@example.jp", Name: "test3", Password: passwordEncrypt("pass")},
	}

	db.Create(&users)
}

// PasswordEncrypt 暗号(Hash)化
func passwordEncrypt(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
