package model

import "time"

type Book struct {
	ID          int64     `json:"ID"`
	BookName    string    `json:"bookName"`
	Author      string    `json:"author"`
	IsStar      bool      `json:"isStar"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	CoverImage  string    `json:"coverImage"`
	CreateTime  time.Time `json:"createTime"`
	IsDeleted   bool      `json:"isDeleted"`
	Comment     Message   `json:"comment"`
	CommentNum  int       `json:"commentNum"`
	Label       string    `json:"label"`
}
