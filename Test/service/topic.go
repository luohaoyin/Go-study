package service

import (
	db2 "Test/db"
	"encoding/json"
	"fmt"
	"test/model"
	"time"
)

// InsertTopic 插入一个话题
func InsertTopic(content string,topicName string)bool{
	sqlstr := `insert into topic(content,topicName,author)values(?,?,?)`
	_,err := db2.DB.Exec(sqlstr,content,topicName,freshtoken.UserName)
	if err != nil{
		fmt.Println(err)
		return false
	}
	start := time.Now()
	rs, err := json.Marshal(model.Topic{
		Content: content,
		Name: topicName,
		Num: start.Format(time.RFC850),
		Author: freshtoken.UserName,
	})
	if err != nil{
		fmt.Println(err)
		return false
	}
	fmt.Println(string(rs))
	str := "create table %s (username varchar(255),message varchar(255),num varchar(255),like int(255))"
	sql:=fmt.Sprintf(str,topicName)
	_,err = db2.DB.Exec(sql)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	//创建一个本地文本保存话题名
	rs, err = json.Marshal(model.Topic{
		Content: content,
		Name: topicName,
		Num: start.Format(time.RFC850),
		Author: freshtoken.UserName,
	})
	K:=string(rs[:])
	model.WriteWithIoutil("topicName.txt",K)
	fmt.Print("加入成功")
	return true
}

// ShowTopic 展示一个话题
func ShowTopic(topicName string) string {
	var u model.Topic
	var k model.Message
	var N []byte
	sqlStr := "select topicName,content from topic where topicName = ?"
	rowOBJ:= db2.DB.QueryRow(sqlStr,topicName)
	err := rowOBJ.Scan(&u.Name,&u.Content)
	if err!=nil {
		fmt.Println(err)
		return ""
	}
	start := time.Now()
	rs, err := json.Marshal(model.Topic{
		Content: u.Content,
		Name: topicName,
		Num: start.Format(time.RFC850),
		Author: freshtoken.UserName,
	})
	str := "select * from %s"
	sql:=fmt.Sprintf(str,topicName)
	rows,err:= db2.DB.Query(sql)
	for rows.Next(){
		err := rows.Scan(&k.Username,&k.Message,&u.Num)
		if err!=nil{
			fmt.Println(err)
			return ""
		}
		rs , err := json.Marshal(model.Message{
			Username: k.Username,
			Message: k.Message,
			Num: u.Num,
		})
		N = model.BytesCombine(N,rs)
	}
	rows.Close()
	rs = model.BytesCombine(N,rs)
	return string(rs[:])
}

// InsertComment 评论一个话题
func InsertComment(topicName string,message string) string {
	var u model.Message
	var N []byte
	str := "select * from %s"
	sql:=fmt.Sprintf(str,topicName)
	rows,err:= db2.DB.Query(sql)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	start := time.Now()
	T := start.Format(time.RFC850)
	str = "insert into %s(username,message,num)values(?,?,?)"
	sql = fmt.Sprintf(str,topicName)
	_,err = db2.DB.Exec(sql,freshtoken.UserName,message,T)
	if err != nil{
		fmt.Println(err)
		return ""
	}
	for rows.Next(){
		err := rows.Scan(&u.Username,&u.Message,&u.Num,&u.Like)
		if err!=nil{
			fmt.Println(err)
			return ""
		}
		rs , err := json.Marshal(model.Message{
			Username: u.Username,
			Message: u.Message,
			Num: u.Num,
			Like: u.Like,
		})
		N = model.BytesCombine(N,rs)
	}
	rows.Close()
	return string(N[:])
}

// ChangeTopicName 修改话题
func ChangeTopicName(name string,content string,title string) string {
	sql:="UPDATE topic SET content = ? , name = ? WHERE name = %s"
	str := fmt.Sprintf(sql,name)
	_,err := db2.DB.Exec(str,content,content,title)
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

// DeleteTopic 删除话题及其评论
func DeleteTopic(id string) string {
	sql := "drop table %s"
	str := fmt.Sprintf(sql,id)
	_,err := db2.DB.Exec(str)
	if err!=nil {
		return ""
	}
	sql = "DELETE FORM topic where topicName = %s"
	str = fmt.Sprintf(sql,id)
	_,err = db2.DB.Exec(str)
	if err!=nil {
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}

// GetLike 点赞
func GetLike(topic string,id string) string {
	var k int
	sql := "select like from %s where topicName = ?"
	str:=fmt.Sprintf(sql,topic)
	rowOBJ:= db2.DB.QueryRow(str,id)
	err := rowOBJ.Scan(k)
	if err!=nil {
		fmt.Println(err)
		return ""
	}
	k += 1
	sql = "UPDATE %s SET content = ? WHERE name = %s"
	str = fmt.Sprintf(sql,topic,id)
	_,err = db2.DB.Exec(str,k)
	if err!=nil {
		return ""
	}
	rs , err := json.Marshal(model.Data{
		Info: "success",
		Status: 10000,
	})
	return string(rs)
}