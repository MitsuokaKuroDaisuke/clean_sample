package domain

type OriginalMember struct {
	ID          int    `json:"id"`
	Email       string `json:"email" gorm:"type:varchar(255);not null"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	MatchData1  string `json:"match_data1"`
	DetailData1 string `json:"detail_data1"`
	Flg         bool   `json:"flg"`
}

type OriginalMemberRepository interface {
	OriginalDataByID(id int) OriginalMember
}
