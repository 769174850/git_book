package router

import (
	"becourse7/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 创建gin引擎，使用"github.com/gin-gonic/gin"包提供的方法，例如：
	// 创建一个gin引擎
	r := gin.Default()

	r.LoadHTMLGlob("html/*")

	// 定义路由
	r.GET("/page", controllers.RenderPage)
	r.GET("/message", controllers.GetMessage)           // 获取所有留言
	r.POST("/message", controllers.AddMessage)          // 添加留言
	r.DELETE("/message/:id", controllers.DeleteMessage) // 删除留言

	return r
}
