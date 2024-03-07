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

type AuthResponse struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   AccessToken `json:"data"`
}

type AccessToken struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
}

var TokenOK = AuthResponse{
	Status: 10000,
	Info:   "success",
	Data:   AccessToken{},
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
	Info:   "Username already exists",
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

func RespErr(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Unauthorized)
}

func RespAlreadyReported(c *gin.Context) {
	c.JSON(http.StatusBadRequest, AlreadyReported)
}

func RespOKWithToken(c *gin.Context, refreshToken, token string) {
	response := TokenOK
	response.Data.RefreshToken = refreshToken
	response.Data.Token = token
	c.JSON(http.StatusOK, response)
}
