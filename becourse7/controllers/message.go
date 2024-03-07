package controllers

import (
	"becourse7/dao"
	"becourse7/models"
	"becourse7/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetMessage 处理GET请求，用于用户查看留言
func GetMessage(c *gin.Context) {
	// 调用dao.GetMessages函数，获取所有留言
	messages, err := dao.GetMessages()
	if err != nil {
		// 如果发生错误，返回500状态码和错误信息
		util.RespInternalErr(c)
		return
	}
	// 如果成功，返回200状态码和留言列表
	util.RespOKWithData(c, messages)
}

// AddMessage 处理POST请求，用于用户发布留言
func AddMessage(c *gin.Context) {
	// 定义一个Message变量，用于接收请求体中的参数
	var m models.Message
	// 使用c.ShouldBindJSON方法，将请求体中的JSON数据绑定到Message变量中
	err := c.ShouldBindJSON(&m)
	if err != nil {
		//如果发生错误，返回400状态码和错误信息
		util.RespParamErr(c)
		return
	}

	// 调用dao.AddMessage函数，添加一条留言
	err = dao.AddMessage(m)
	if err != nil {
		// 如果发生错误，返回500状态码和错误信息
		util.RespInternalErr(c)
		return
	}
	// 如果成功，返回200状态码和留言信息
	util.RespOK(c)
}

// DeleteMessage 处理DELETE请求，用于用户删除留言
func DeleteMessage(c *gin.Context) {
	// 使用c.Param(方法，获取路由参数中的id
	id := c.Param("id")
	fmt.Println(id)
	intid, err := strconv.Atoi(id)
	if err != nil {
		util.RespParamErr(c)
		return
	}
	// 调用dao.DeleteMessage函数，删除一条留言
	err = dao.DeleteMessage(int64(intid))
	if err != nil {
		// 如果要删除的记录不存在
		if err == util.NoRecordExistError {
			util.RespNormErr(c, 401, util.NoRecordExistError.Error())
			return
		}
		// 如果发生错误，返回500状态码和错误信息
		util.RespInternalErr(c)
		return
	}
	// 如果成功，返回200状态码和删除成功的信息
	util.RespOK(c)
}

func RenderPage(c *gin.Context) {
	c.HTML(http.StatusOK, "testmessage.html", gin.H{})
}
