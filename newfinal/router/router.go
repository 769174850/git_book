package router

import (
	"github.com/gin-gonic/gin"
	"newfinal/control"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", control.Register)               //注册一个新用户
	r.POST("/user/token", control.Login)                //登录用户
	r.POST("/user/token/refresh", control.RefreshToken) //刷新token维持登录状态
	r.PUT("/user/password", control.ChangePassword)     //更改用户密码
	r.GET("/user/info/:id", control.GetUserInfo)        //获取用户信息
	r.PUT("/user/info/:id", control.ChangeUser)         //更改用户信息

	r.GET("/book/list", control.GetBookList) //获取网站内的所有书籍
	r.GET("book/search", control.SearchBook) //根据书本名字查找书本
	r.PUT("book/star", control.StarBook)     //搜藏书本
	r.GET("book/label", control.LabelBook)   //根据书籍标签查找书本

	r.GET("comment/:bookID", control.GetMessage)    //获取书籍书评
	r.POST("comment/:bookID", control.AddMessages)  //书写一个书评
	r.DELETE("comment/:id", control.DeleteMessages) //删除书评
	r.PUT("comment/:id", control.UpdateMessage)     //更新书评

	r.PUT("operate/praise", control.PraiseMessage)     //点赞书评
	r.GET("operate/collect/list", control.GetUserBook) //获取用书收藏的书籍

	return r
}
