package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// DB 是一个*sql.DB对象，用于操作数据库
var DB *sql.DB

// InitDB 初始化数据库连接
func InitDB() {
	// dsn为数据库连接信息，需要根据自己的数据库配置进行修改
	dsn := "root:123456@tcp(127.0.0.1:3306)/test1?charset=utf8&parseTime=True&loc=Local"
	//dsn := "username:password@tcp(host)/db_name?charset=utf8&parseTime=True&loc=Local"
	// db为一个*sql.DB对象，用于操作数据库
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		// 处理错误
		log.Fatal(err)
	}
	// 设置数据库连接池的最大连接数和空闲连接数
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	// 测试数据库连接是否正常
	err = DB.Ping()
	if err != nil {
		// 处理错误
		log.Fatal(err)
	}
}
