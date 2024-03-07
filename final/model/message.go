package model

import "time"

type Message struct {
	ID         int64     `json:"ID"`
	BookID     int64     `json:"bookID"`
	UserID     int64     `json:"userID"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
	IsDeleted  bool      `json:"isDeleted"`
}
