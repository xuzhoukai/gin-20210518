package dao_mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var MysqlDB *sql.DB

func init() {
	var err error
	dsn := "gin01:gin01@tcp(139.224.0.98:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	MysqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		logrus.Errorf("fail to open mysql err:%v", err)
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = MysqlDB.Ping()
	if err != nil {
		return
	}
	MysqlDB.SetMaxOpenConns(100)
	MysqlDB.SetMaxIdleConns(30)
}
