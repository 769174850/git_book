package models

import (
	"time"
)

// Message 结构体，对应message表
type Message struct {
	ID       int64     `json:"id"`
	Detail   string    `json:"detail"`
	CreateAt time.Time `json:"createAt"`
	UpDateAt time.Time `json:"upDateAt"`
}
