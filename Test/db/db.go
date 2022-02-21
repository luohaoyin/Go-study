package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)
var (
	DB *sqlx.DB
)
func InitDB()error {
	var err error
	DB, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/users")
	if err != nil {
		fmt.Printf("open %s faild, err:%v\n",err)
		return err
	}
	//查看是否连接成功
	err = DB.Ping()
	if err != nil{
		return err
	}
	fmt.Println("连接成功！")
	return nil
}