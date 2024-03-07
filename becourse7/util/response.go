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

func RespOKWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data":   data,
	})
}

func RespOK(c *gin.Context) {
	c.JSON(http.StatusOK, OK)
}

func RespParamErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ParamError)
}

func RespInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalError)
}

func RespNormErr(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
		"data":   "",
	})
}
