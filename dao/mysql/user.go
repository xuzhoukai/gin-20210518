package dao_mysql

import (
	"fmt"
	"gin-20210518/model"

	"github.com/sirupsen/logrus"
)

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := MysqlDB.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u model.User
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := MysqlDB.QueryRow(sqlStr, 2).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	logrus.Infof("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := MysqlDB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		logrus.Infof("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := MysqlDB.Exec(sqlStr, 39, 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := MysqlDB.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	logrus.Infof("delete success, affected rows:%d\n", n)
}
