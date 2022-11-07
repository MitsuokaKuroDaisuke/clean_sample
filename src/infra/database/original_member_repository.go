package database

import (
	"src/domain"
)

// MemberRepository メンバーリポジトリの実装
type OriginalMemberRepository struct {
	SQLHandler
}

func (rep *OriginalMemberRepository) OriginalDataByID(id int) domain.OriginalMember {
	m := domain.OriginalMember{}
	rep.db.Raw(`SELECT members.id, 
					name, 
					email,
					member_matches.data1 as match_data1,
					member_details.data1 as detail_data1
				FROM members 
					INNER JOIN member_matches on members.id = member_matches.member_id 
					INNER JOIN member_details on members.id = member_details.member_id 
				WHERE members.id = ?`, id).Scan(&m)
	return m
}
