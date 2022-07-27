package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func TouchSql() {
	// DSN:Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	// 创建数据库连接,返回一个*sql.DB类型的数据库连接,这个连接可以用于执行SQL语句
	//指定数据库mysql,指定数据源dsn
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
