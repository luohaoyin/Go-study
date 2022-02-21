package service

import (
	"Test/db"
	"encoding/json"
	"fmt"
	"test/model"
)

var freshtoken model.FreshToken

func Login(name string,password string) string {
	var u model.User
	sqlStr := "select name,password from user where name = ?"
	rowOBJ:= db.DB.QueryRow(sqlStr,name)
	err := rowOBJ.Scan(&name,&u.Password)
	if err!=nil{
		return ""
	}
	if u.Password == password {
		fmt.Println("登录成功")
	} else{
		fmt.Println("密码错误")
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Status: 10000,
		Info: "success",
		RefreshToken: "refreshToken",
		Token: "token",
	})
	freshtoken.UserName = u.UserName
	freshtoken.Password = u.Password
	return string(rs)
}