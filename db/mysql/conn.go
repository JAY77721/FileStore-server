package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(192.168.182.139:3307)/fileserver?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error open DB:", err)
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Cannot connect to MySQL:", err)
		return
	}

	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)
	fmt.Println("ping successfull")
	if err != nil {
		fmt.Println("failed to connected to mysql,err:" + err.Error())
		os.Exit(1)
	}
}

// DBConn：返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
