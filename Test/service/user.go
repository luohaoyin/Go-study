package service

import (
	"Test/db"
	"Test/model"
	"encoding/json"
	"fmt"
)
func ChangePassword(password string,newPassword string) string {
	var u model.User
	sqlStr := "select name,password from user where name = ?"
	rowOBJ:= db.DB.QueryRow(sqlStr,freshtoken.UserName)
	err := rowOBJ.Scan(&u.UserName,&u.Password)
	if err!=nil{
		fmt.Println("账号或密码错误1")
		return ""
	}
	if u.Password == password {
		sql := "UPDATE user SET password = ? WHERE name = %s"
		str := fmt.Sprintf(sql, freshtoken.UserName)
		_, err := db.DB.Exec(str, newPassword)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("修改成功")
	}
	rs , err := json.Marshal(model.Data{
		Status: 10000,
		Info: "success",
	})
	freshtoken.UserName = u.UserName
	freshtoken.Password = newPassword
	return string(rs)
}
func QueryInformation(username string) string {
	var u model.User
	sqlstr := "select introduction,emil,phone,qq,gender,birth from user where name = ?"
	rowOBJ:= db.DB.QueryRow(sqlstr,freshtoken.UserName)
	err := rowOBJ.Scan(&u.Introduction,&u.Emil,&u.Phone,&u.QQ,&u.Gender,&u.Birth)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.User{
		UserName: freshtoken.UserName,
		Introduction: u.Introduction,
		Emil: u.Emil,
		Phone: u.Phone,
		QQ: u.QQ,
		Gender: u.Gender,
		Birth: u.Birth,
	})
	return string(rs)
}
func ChangeInformation(name string,emil string,introduction string,phone int,qq int,
	gender string,birth string) string {
	sqlStr:="UPDATE user SET " +
		"name = ?,introduction= ?,emil= ?," +
		"phone= ?,qq= ?,gender= ?,birth= ? WHERE name = ?"
	_,err :=db.DB.Exec(sqlStr,name,introduction,emil,phone,qq,gender,birth,freshtoken.UserName)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.User{
		UserName: name,
		Introduction: introduction,
		Emil: emil,
		Phone: phone,
		QQ: qq,
		Gender: gender,
		Birth: birth,
	})
	return string(rs)
}

// CollectTopic 收藏话题
func CollectTopic(id string) string {
	sql := `insert into %s(topicName)values(?)`
	str := fmt.Sprintf(sql,freshtoken.UserName)
	_,err :=db.DB.Exec(str,id)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}

// FollowUser 关注用户
func FollowUser(id string) string {
	sql := `insert into %s(name)values(?)`
	str := fmt.Sprintf(sql,freshtoken.UserName)
	_,err :=db.DB.Exec(str,id)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}