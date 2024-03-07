package model

type User struct {
	ID       int64      `json:"ID"`
	Username string     `json:"username"`
	Gender   string     `json:"gender" default:"unknown"`
	Password string     `json:"password"`
	Avatar   string     `json:"avatar"`
	Books    []UserBook `json:"books"`
}
