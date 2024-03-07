package dao

import (
	"errors"
	"log"
	"newfinal/model"
	"time"
)

func GetMessage(bookID int64) ([]model.Message, error) {
	var messages []model.Message //定义切片储存
	s := "SELECT message.ID, message.BookID, book.BookName, message.PushUserID, user.UserName, message.Content, message.CreateTime, user.Avatar, message.IsDeleted FROM message LEFT JOIN book AS book ON message.BookID = book.ID LEFT JOIN user AS user ON message.PushUserID = user.ID WHERE message.BookID = ?"
	rows, err := DB.Query(s, bookID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message model.Message
		err := rows.Scan(&message.ID, &message.BookID, &message.BookName, &message.PushUserID, &message.PushUserName, &message.Content, &message.CreateTime, &message.Avatar, &message.IsDeleted)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		messages = append(messages, message)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return messages, nil
}

func AddMessage(message model.Message) (int64, error) {
	s := "INSERT INTO message (content,bookID,pushUserID) VALUES (?, ?, ?)"
	result, err := DB.Exec(s, message.Content, message.BookID, message.PushUserID)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	message.ID = id
	return message.ID, nil
}

func DeleteMessage(id int64) error {
	s := "UPDATE message SET isDeleted = ? where id = ?"

	result, err := DB.Exec(s, true, id)
	if err != nil {
		log.Println(err)
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	if rows == 0 {
		return errors.New("messages does not exist")
	}

	return nil
}

func UpdateMessage(content string, id int64) error {
	s := "UPDATE message SET content = ? ,createTime = ? WHERE id = ?"

	result, err := DB.Exec(s, content, time.Now(), id)
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
