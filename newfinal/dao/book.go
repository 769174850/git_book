package dao

import (
	"log"
	"newfinal/model"
	"time"
)

func GetBook() ([]model.Book, error) {
	var books []model.Book //定义切片储存
	s := "SELECT * FROM book"
	rows, err := DB.Query(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.BookName, &book.Author, &book.IsStar, &book.Link, &book.Description, &book.CoverImage, &book.CreateTime, &book.IsDeleted, &book.CommentNum, &book.Label)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		books = append(books, book)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(books) == 0 {
		return nil, nil
	}

	return books, nil
}

func SearchBookByBookName(BookName string) ([]model.Book, error) {
	var books []model.Book
	s := "SELECT ID, BookName, Author, IsStar, Link, Description, CoverImage, CreateTime, IsDeleted, CommentNum, Label FROM book WHERE BookName = ?"
	rows, err := DB.Query(s, BookName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.BookName, &book.Author, &book.IsStar, &book.Link, &book.Description, &book.CoverImage, &book.CreateTime, &book.IsDeleted, &book.CommentNum, &book.Label)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		books = append(books, book)
	}

	return books, err
}

func StarBook(userID, bookID int64) error {
	var book model.Book

	s1 := "SELECT ID, BookName, Author, IsStar, Link, Description, CoverImage, CreateTime, IsDeleted, CommentNum, Label FROM book WHERE ID = ?"
	//查询数据
	err := DB.QueryRow(s1, bookID).Scan(&book.ID, &book.BookName, &book.Author, &book.IsStar, &book.Link, &book.Description, &book.CoverImage, &book.CreateTime, &book.IsDeleted, &book.CommentNum, &book.Label)
	if err != nil {
		log.Println(err)
		return err
	}

	s2 := "INSERT INTO user_book(BookID, UserID, BookName, Author, IsStar, Link, Description, CoverImage, CreateTime, AddTIme) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	//将查询到的数据复制到userBook中实现搜藏
	_, err = DB.Exec(s2, book.ID, userID, book.BookName, book.Author, 1, book.Link, book.Description, book.CoverImage, book.CreateTime, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
	//s := "UPDATE book SET IsStar = ? WHERE bookID = ?"
	//result, err := DB.Exec(s, IsStar, bookID)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//
	//if rowsAffected == 0 {
	//	log.Println(err)
	//	return err
	//}
	//
	//return nil
}

func LabelBook(label string) ([]model.Book, error) {
	var books []model.Book
	s := "SELECT ID, BookName, Author, IsStar, Link, Description, CoverImage, CreateTime, IsDeleted, CommentNum, Label FROM book WHERE Label = ?"
	rows, err := DB.Query(s, label)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.BookName, &book.Author, &book.IsStar, &book.Link, &book.Description, &book.CoverImage, &book.CreateTime, &book.IsDeleted, &book.CommentNum, &book.Label)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		books = append(books, book)
	}

	return books, err
}
