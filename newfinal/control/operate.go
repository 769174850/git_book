package control

import (
	"github.com/gin-gonic/gin"
	"newfinal/dao"
	"newfinal/util"
)

func PraiseMessage(c *gin.Context) {
	var request struct {
		ID int64 `json:"ID"`
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	err = dao.PraiseMessage(request.ID)
	if err != nil {
		util.RespInternalErr(c)
		return
	}

	util.RespOK(c)
}

func GetUserBook(c *gin.Context) {
	var request struct {
		UserID int64 `json:"UserID"`
	}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.RespParamErr(c)
		return
	}

	userBook, err := dao.GetUserBook(request.UserID)

	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOKWithData(c, userBook)
}
