package entity

import (
	"context"
	"time"
)

type MemberDetail struct {
	ID        int       `json:"id"`
	MemberID  int       `json:"member_id"`
	Flg       bool      `json:"flg"`
	Data1     string    `json:"data1"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type MemberDetailRepository interface {
	Create(ctx context.Context)
}
