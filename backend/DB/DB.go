package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func TouchSql() {
	// DSN:Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/sqlname"
	// 创建数据库连接,返回一个*sql.DB类型的数据库连接,这个连接可以用于执行SQL语句
	//指定数据库mysql,指定数据源dsn
	//Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面 (第三方库的原因)
}
