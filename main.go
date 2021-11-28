package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 设置内存下限8 MiB(默认好像是32 MiB)
	r.Static("/", "./public")
	r.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		// 通过c.FormFile来获取文件
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {//SaveUploadedFile()函数，存入值。
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.",
			file.Filename, name, email))
	})
}