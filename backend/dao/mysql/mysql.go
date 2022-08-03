package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//这里的db不是一个全局变量，但通过对外封装方法的方式，可以让外部调用，也可以让外部调用
var db *sqlx.DB

//初始化mysql，返回一个sqlx.DB，用于后续的数据库操作
func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn) //连接数据库，返回一个sqlx.DB，用于后续的数据库操作，这里的db不是一个全局变量，但通过对外封装方法的方式，可以让外部调用，也可以让外部调用
	//检查连接是否成功，如果连接不成功，则返回错误
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}
	//设置最大打开的连接数，默认值为0表示不限制
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	//设置最大闲置的连接数，如果闲置的连接数大于SetMaxOpenConns的值，最大闲置连接数会被设置为SetMaxOpenConns的值
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	//返回nil，表示无error，连接成功
	return
}

//对外封装一个close方法
func Close() {
	_ = db.Close()
}
