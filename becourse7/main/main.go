package main

import (
	"becourse7/dao"
	"becourse7/router"
)

func main() {
	// 连接数据库
	dao.InitDB()

	// 创建 Gin 引擎
	r := router.InitRouter()

	// 服务，监听8080端口，处理请求
	r.Run(":8080")
}
