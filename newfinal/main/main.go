package main

import (
	"newfinal/dao"
	"newfinal/router"
)

func main() {
	dao.InitDB()

	r := router.InitRouter()
	r.Run(":8080")
}
