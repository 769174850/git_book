package router

import (
	"final/control"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", control.Register)
	r.POST("/user/token", control.Login)
	r.POST("/user/token/refresh", control.RefreshToken)
	r.PUT("/user/")

	return r
}
