package database

import "src/domain/entity"

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
