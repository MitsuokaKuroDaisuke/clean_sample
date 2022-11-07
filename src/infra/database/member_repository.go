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

func (rep *MemberRepository) OriginalDataByID(id int) entity.Member {
	m := entity.Member{}
	rep.db.Raw(`SELECT members.id, 
					name, 
					email,
					member_matches.data1 
				FROM members 
					INNER JOIN member_matches on members.id = member_matches.member_id 
				WHERE members.id = ?`, id).Scan(&m)
	return m
}
