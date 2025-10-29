package db

import (
	mydb "filestore-server/db/mysql"
	"fmt"
)

// UserSignup：通过用户名及密码完成user表的注册操作
func UserSignup(username string, password string) bool {
	//db := mydb.DBConn()
	//fmt.Println("DB connection:", db)

	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user (`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		fmt.Println("Failed to insert ,err" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println("Exec failed,err" + err.Error())
		return false
	}
	//if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
	//	return true
	//}
	//return false
	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected failed:", err)
		return false
	}

	fmt.Println("Rows affected:", rowsAffected)
	return true
}
