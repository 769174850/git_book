package main

import (
	"final/dao"
	"final/router"
)

func main() {
	dao.InitDB()

	r := router.InitRouter()
	r.Run(":8080")
}
