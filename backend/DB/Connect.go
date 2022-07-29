package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	// 创建数据库连接,返回一个*sql.DB类型的数据库连接,这个连接可以用于执行SQL语句
	//指定数据库mysql,指定数据源dsn
	//Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
	//defer db.Close() // 注意这行代码要写在上面err判断的下面 (第三方库的原因)
}

/*
func (db *DB) SetMaxOpenConns(n int)
设置与数据库建立连接的最大数目。
如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。
如果n<=0，不会限制最大开启连接数，默认为0（无限制）。
*/

/*
func (db *DB) SetMaxIdleConns(n int)
设置连接池中的最大闲置连接数。
如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。
如果n<=0，不会保留闲置连接。
*/
