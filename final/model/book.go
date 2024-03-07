package model

import "time"

type Book struct {
	ID          int64     `json:"ID"`
	BookName    string    `json:"bookName"`
	Author      string    `json:"author"`
	Table       string    `json:"table"`
	Description string    `json:"description"`
	CoverImage  string    `json:"coverImage"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	IsDeleted   bool      `json:"isDeleted"`
	Comment     Message   `json:"comment"`
}
