package entity

import (
	"time"
)

type Member struct {
	ID           int          `json:"id" form:"id"`
	Email        string       `json:"email" gorm:"type:varchar(255);not null"`
	Name         string       `json:"name" gorm:"type:varchar(255);not null"`
	MemberMatch  MemberMatch  `json:"MemberMatch" `
	MemberDetail MemberDetail `json:"MemberDetail" `
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type MemberRepository interface {
	AdminDataByID(id int) Member
}
