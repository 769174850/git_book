package model

import "time"

type UserBook struct {
	ID          int64     `json:"ID"`
	BookID      int64     `json:"bookID"`
	UserID      int64     `json:"userID"`
	BookName    string    `json:"bookName"`
	Author      string    `json:"author"`
	IsStar      bool      `json:"isStar"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	CoverImage  string    `json:"coverImage"`
	CreateTime  time.Time `json:"createTime"`
	AddTime     time.Time `json:"addTime"`
	IsDeleted   bool      `json:"isDeleted"`
	Comment     []Message `json:"comment"`
}
