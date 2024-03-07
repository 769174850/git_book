package dao

import (
	"final/model"
)

func VerifyUserAndGetID(username, password string) (int64, error) {
	var userID int64
	sql := "SELECT ID FROM user WHERE Username=? AND Password=?"
	err := DB.QueryRow(sql, username, password).Scan(&userID)
	if err != nil {
		return 0, nil
	}
	return userID, nil
}

func GetUser() ([]model.User, error) {
	var users []model.User //定义切片存储用户
	sql := "SELECT ID, Username, Gender, Password FROM user "
	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Username, &u.Gender, &u.Password) //查询数据
		if err != nil {
			return nil, err
		}

		users = append(users, u) //导入切片中
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
func AddUser(u model.User) error {
	sql := "INSERT INTO user (username, password) VALUES (?, ?)"
	result, err := DB.Exec(sql, u.Username, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err

	}

	u.ID = id
	return nil
}
