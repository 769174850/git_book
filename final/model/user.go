package model

type User struct {
	ID       int64  `json:"ID"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
	Books    Book   `json:"books"`
}
