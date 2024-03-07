package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type respTemplate struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

type AccessToken struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
}

var OK = respTemplate{
	Status: 200,
	Info:   "success",
	Data:   "",
}

var ParamError = respTemplate{
	Status: 300,
	Info:   "params error",
	Data:   "",
}

var InternalError = respTemplate{
	Status: 500,
	Info:   "internal error",
	Data:   "",
}

var Unauthorized = respTemplate{
	Status: 401,
	Info:   "invalid username or password",
	Data:   "",
}

var AlreadyReported = respTemplate{
	Status: 400,
	Info:   "Username or password already exists",
	Data:   "",
}

var Unauthorized1 = respTemplate{
	Status: 401,
	Info:   "invalid password please try again",
	Data:   "",
}

var NotFound = respTemplate{
	Status: 404,
	Info:   "page not found",
	Data:   "",
}

func RespOKWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   data,
	})
} //带有数据的返回200

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
} //返回200

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
} //返回获取信息失败400

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
} //返回内部错误500

func RespNormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
		"data":   "",
	})
}

func RespErr(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Unauthorized)
} //返回账号不存在错误401

func RespUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Unauthorized1)
} //返回修改密码原密码错误401

func RespAlreadyReported(c *gin.Context) {
	c.JSON(http.StatusBadRequest, AlreadyReported)
} //返回账号已存在错误400

func RespNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, NotFound)
} //返回页面不存在错误404
