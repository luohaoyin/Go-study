package dao

import (
	"Test/db"
	"Test/model"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/jmoiron/sqlx"
)

const secret = "888888"

//检查用户名的用户是否存在
func CheckUserExist(name string)(error error){
	sqlStr := `select name from user where username ?`
	var c bool
	c = (sqlStr == name)
	if c != false{
		return errors.New("用户已经存在~")
	}
	return
}

func Register(name string,password string,emil string,introduction string,phone int,qq int,
	gender string,birth string)string{
	sqlStr := `insert into user(name,password,introduction,emil,phone,qq,gender,birth)" +
		"values(?,?,?,?,?,?,?,?)`
	  _,err := db.DB.Exec(sqlStr,name,password,introduction,emil,phone,qq,gender,birth)
	  if err!=nil{
		  fmt.Println(err)
	  }
	  //顺势建立一张表去保存用户的相关收藏以及关注话题
	str := "create table %s (name varchar(255),topicName varchar(255))"
	sql :=fmt.Sprintf(str,model.FreshToken{}.UserName)
	_,err =db.DB.Exec(sql)
	if err!=nil {
		fmt.Println(err)
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}