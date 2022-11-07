package main

import (
	"fmt"
	"os"
	"src/domain/entity"

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
		{Email: "d.mmm@example.jp", Name: "mitsu"},
		{Email: "t.tt@example.jp", Name: "test2"},
		{Email: "y.ooo@example.jp", Name: "test3"},
	}

	db.Create(&users)

	db.Migrator().DropTable(&entity.Member{})
	db.AutoMigrate(&entity.Member{})

	members := []entity.Member{
		{Email: "d.mmm@example.jp", Name: "mitsu"},
		{Email: "t.tt@example.jp", Name: "membertest2"},
		{Email: "y.ooo@example.jp", Name: "test3"},
	}

	db.Create(&members)

	db.Migrator().DropTable(&entity.MemberMatch{})
	db.AutoMigrate(&entity.MemberMatch{})

	memberMatchs := []entity.MemberMatch{
		{MemberID: 1, Data1: "data1あ", Data2: "data2い", Data3: "data3う"},
		{MemberID: 3, Data1: "data1か", Data2: "data2き", Data3: "data3く"},
		{MemberID: 2, Data1: "data1さ", Data2: "data2し", Data3: "data3す"},
	}

	db.Create(&memberMatchs)

	db.Migrator().DropTable(&entity.MemberDetail{})
	db.AutoMigrate(&entity.MemberDetail{})

	memberDetails := []entity.MemberDetail{
		{MemberID: 1, Flg: true, Data1: "data1詳細た"},
		{MemberID: 2, Flg: true, Data1: "data1詳細な"},
		{MemberID: 3, Flg: false, Data1: "data1詳細は"},
	}

	db.Create(&memberDetails)

}
