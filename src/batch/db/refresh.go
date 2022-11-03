package main

import (
	"fmt"
	"os"
	"www/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sTokyo", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"), "%2F")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	db.Migrator().DropTable(&domain.User{})
	db.AutoMigrate(&domain.User{})

	users := []domain.User{
		{Email: "d.mitsuoka@newbees.jp", Name: "mitsu"},
		{Email: "t.tt@newbees.jp", Name: "nog"},
		{Email: "y.ooo@newbees.jp", Name: "peco"},
	}

	db.Create(&users)

}
