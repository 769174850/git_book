package control

import (
	"github.com/gin-gonic/gin"
	"log"
	"newfinal/dao"
	"newfinal/util"
)

func GetBookList(c *gin.Context) {
	books, err := dao.GetBook()
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	if len(books) == 0 {
		util.RespNormErr(c, 200, "this user has not collected any books yet")
		return
	}
	util.RespOKWithData(c, books)
}

func SearchBook(c *gin.Context) {
	var request struct {
		BookName string `json:"bookName"`
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		util.RespParamErr(c)
		return
	}

	books, err := dao.SearchBookByBookName(request.BookName)
	if err != nil {
		log.Println(err)
		util.RespInternalErr(c)
		return
	}
	util.RespOKWithData(c, books)
}

func StarBook(c *gin.Context) {
	var request struct {
		UserID int64 `json:"userID"`
		BookID int64 `json:"bookID"`
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	err = dao.StarBook(request.UserID, request.BookID)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func LabelBook(c *gin.Context) {
	var request struct {
		Label string `json:"label"`
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		util.RespParamErr(c)
		return
	}

	books, err := dao.LabelBook(request.Label)
	if err != nil {
		log.Println(err)
		util.RespInternalErr(c)
		return
	}
	util.RespOKWithData(c, books)
}
