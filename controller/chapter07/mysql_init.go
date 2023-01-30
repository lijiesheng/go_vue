package chapter07

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"

var db *gorm.DB

func init() {
	mysql_db, db_err := gorm.Open("mysql", "root:ljs024816@tcp(101.42.97.221:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local")
	if db_err != nil {
		panic(db_err)
	}
	db = mysql_db
	//defer db.Close() // 关闭连接
}
