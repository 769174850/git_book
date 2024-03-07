package dao

import (
	"database/sql"
	"errors"
	"log"
	"newfinal/model"
	"strings"
)

type User struct {
	Username string
	Password string
}

type userDB struct {
	ID       int64  `json:"ID"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

type Users struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
}

func GetUserByUsername(username string) (User, error) {
	var user User

	rows := DB.QueryRow("SELECT username, password FROM user WHERE username = ?", username)
	err := rows.Scan(&user.Username, &user.Password)
	if err != nil {
		log.Println(err)
		return user, errors.New("user not found")
	} else if err != nil {
		log.Println(err)
		return user, errors.New("internal server error")
	}
	return user, nil
}

func VerifyUserAndGetID(username, password string) (int64, error) {
	var userID int64
	s := "SELECT ID FROM user WHERE Username=? AND Password=?"
	err := DB.QueryRow(s, username, password).Scan(&userID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return userID, nil
}

func GetUser() ([]model.User, error) {
	var users []model.User //定义切片存储用户
	s := "SELECT ID, Username, Avatar, Gender, Password FROM user "
	rows, err := DB.Query(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Username, &u.Avatar, &u.Gender, &u.Password) //查询数据
		if err != nil {
			log.Println(err)
			return nil, err
		}

		users = append(users, u) //导入切片中
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func AddUser(u model.User) error {
	s := "INSERT INTO user (username, password) VALUES (?, ?)"
	result, err := DB.Exec(s, u.Username, u.Password)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return err

	}

	u.ID = id
	return nil
}

func ChangeUserPassword(newPassword, username string) error {
	s := "UPDATE user set password = ? where username = ?"
	result, err := DB.Exec(s, newPassword, username)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(err)
		return err
	}

	return nil
}

func GetUsersInfo(id string) (model.User, error) {
	var user model.User
	s := "SELECT ID, Username, Avatar, Gender, Password FROM user WHERE ID = ?"
	var userDB userDB
	err := DB.QueryRow(s, id).Scan(&userDB.ID, &userDB.Username, &userDB.Avatar, &userDB.Gender, &userDB.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}

	user.ID = userDB.ID
	user.Username = userDB.Username
	user.Gender = userDB.Gender
	user.Avatar = userDB.Avatar
	user.Password = userDB.Password

	return user, nil
}

func ChangeUser(id int64, request Users) error {
	updated := map[string]interface{}{
		"username": request.Username,
		"gender":   request.Gender,
		"avatar":   request.Avatar,
	}
	s := "UPDATE user SET "
	var values []interface{}
	var setClauses []string

	for field, value := range updated {
		if value != nil {
			setClauses = append(setClauses, field+" = ?")
			values = append(values, value)
		}
	}

	s += strings.Join(setClauses, ",")
	s += " WHERE ID = ?"
	values = append(values, id)

	result, err := DB.Exec(s, values...)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		log.Println(err)
		return err
	}

	return nil
}
