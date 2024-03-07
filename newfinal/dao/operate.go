package dao

import (
	"database/sql"
	"errors"
	"log"
	"newfinal/model"
)

func PraiseMessage(id int64) error {
	s := "UPDATE message SET isPraise = ? where id = ?"

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

func GetUserBook(id int64) (model.UserBook, error) {
	var userBook model.UserBook
	s := "SELECT ID, BookID, UserID, BookName, Author, IsStar, Link, Description, CoverImage, CreateTime, AddTIme , ISDeleted FROM user_book WHERE UserID = ?"
	err := DB.QueryRow(s, id).Scan(&userBook.ID, &userBook.BookID, &userBook.UserID, &userBook.BookName, &userBook.Author, &userBook.IsStar, &userBook.Link, &userBook.Description, &userBook.CoverImage, &userBook.CreateTime, &userBook.AddTime, &userBook.IsDeleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return model.UserBook{}, errors.New("user not found")
		}
		log.Println(err)
		return model.UserBook{}, err
	}

	return userBook, nil
}
