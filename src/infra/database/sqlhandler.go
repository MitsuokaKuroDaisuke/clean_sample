package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLHandler ハンドラ
type SQLHandler struct {
	db *gorm.DB
}

// NewSQLHandler ハンドラ初期化
func NewSQLHandler() SQLHandler {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sTokyo", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"), "%2F")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := SQLHandler{}
	sqlHandler.db = db

	return sqlHandler
}

func NewTestSQLHandler() SQLHandler {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sTokyo", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), "test-db", os.Getenv("DB_DATABASE"), "%2F")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := SQLHandler{}
	sqlHandler.db = db

	return sqlHandler
}
