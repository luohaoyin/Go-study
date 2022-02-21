package model

import (
"fmt"
"io/ioutil"
"strings"
)

// WriteWithIoutil 文件写入
func WriteWithIoutil(name string,content string) {
	data :=  []byte(content)
	if ioutil.WriteFile(name,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
	}
}

// Ioutil 文件读取
func Ioutil(name string)string{
	var result string
	if contents,err := ioutil.ReadFile(name);err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result = strings.Replace(string(contents),"\n","",1)
	}
	return result
}