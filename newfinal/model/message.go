package model

import "time"

type Message struct {
	ID           int64     `json:"ID"`
	BookID       int64     `json:"bookID"`
	BookName     string    `json:"bookName"`
	PushUserID   int64     `json:"pushUserID"`
	PushUserName string    `json:"pushUserName"`
	Content      string    `json:"content"`
	CreateTime   time.Time `json:"createTime"`
	Avatar       string    `json:"avatar"`
	IsDeleted    bool      `json:"isDeleted"`
	IsPraise     bool      `json:"isPraise"`
	PraiseCount  int       `json:"praiseCount"`
	IsFocus      bool      `json:"isFocus"`
}
