package control

import (
	"errors"
	"github.com/gin-gonic/gin"
	"newfinal/dao"
	"newfinal/model"
	"newfinal/util"
	"strconv"
)

func GetMessage(c *gin.Context) {
	bookID := c.Param("bookID")
	id, err := strconv.ParseInt(bookID, 10, 64)

	messages, err := dao.GetMessage(id)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	if len(messages) == 0 {
		util.RespNormErr(c, 200, "this book has not any comment yet")
		return
	}
	util.RespOKWithData(c, messages)
}

func AddMessages(c *gin.Context) {
	var message model.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	id, err := dao.AddMessage(message)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOKWithData(c, id)
}

func DeleteMessages(c *gin.Context) {
	commentID := c.Param("id")
	id, err := strconv.ParseInt(commentID, 10, 64)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	err = dao.DeleteMessage(id)
	if err != nil {
		if errors.Is(err, errors.New("message does not exist")) {
			util.RespNormErr(c, 401, errors.New("message does not exist").Error())
			return
		}

		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
}

func UpdateMessage(c *gin.Context) {
	var request struct {
		ID      int64  `json:"ID"`
		Content string `json:"Content"`
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	err = dao.UpdateMessage(request.Content, request.ID)
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}
