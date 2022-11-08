package database

import (
	"context"
	"src/domain/entity"

	"gorm.io/gorm"
)

// MemberRepository メンバーリポジトリの実装
type MemberRepository struct {
	SQLHandler
}

// AdminDataByID 取得
func (rep *MemberRepository) AdminDataByID(id int) entity.Member {
	m := entity.Member{}
	rep.db.Where("id = ?", id).Preload("MemberMatch").Preload("MemberDetail").First(&m)
	return m
}

type MemberDetailRepository struct {
	SQLHandler
}

func (rep *MemberDetailRepository) Create(ctx context.Context) {
	md := entity.MemberDetail{MemberID: 3, Flg: true, Data1: "新し"}
	tx, _ := ctx.Value(TxCtxKey).(*gorm.DB)
	tx.Create(&md)
}

type MemberMatchRepository struct {
	SQLHandler
}

func (rep *MemberMatchRepository) Create(ctx context.Context) {
	mm := entity.MemberMatch{MemberID: 1, Data1: "だた１", Data2: "だた２", Data3: "打田３３"}
	tx, _ := ctx.Value(TxCtxKey).(*gorm.DB)
	tx.Create(&mm)
}
