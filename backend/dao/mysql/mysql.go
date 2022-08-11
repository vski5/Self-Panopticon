package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		/*
		   此处不该是返回nil，
		   但是不这么写的话，在测试里面会报错，应该只写一个return
		   正式做出来东西的时候要记得改
		*/
		//zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		//return

		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return

	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}

//对外封装一个close方法
func Close() {
	_ = db.Close()
}
