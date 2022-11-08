package entity

import (
	"context"
	"time"
)

type MemberMatch struct {
	ID        int       `json:"id"`
	MemberID  int       `json:"member_id"`
	Data1     string    `json:"data1"`
	Data2     string    `json:"data2"`
	Data3     string    `json:"data3"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type MemberMatchRepository interface {
	Create(ctx context.Context)
}
