package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	auth := func(c *gin.Context) {
		value, err := c.Cookie("gin_cookie") //获取cookie

		if err != nil {
			c.JSON(403, gin.H{
				"message": "认证失败,无cookie",
			})
			middleware := gin.HandlerFunc(func(c *gin.Context) {})

			r.Use(middleware)
			v1 := r.Group("v1")
			v1.Use(middleware)
			c.Abort() //终止后面所有的该请求下的函数
		} else { //将获取到的cookie的值写入上下文
			c.Set("cookie", value)
			c.Next() //挂着，跳过下列，返回执行。
			v, _ := c.Get("next")
			fmt.Println(v)
		}
	}
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "123" && password == "321" {// ctrl=b,看参数
			c.SetCookie("gin_cookie", username, 3600, "/",
			        	"", false, true)

			c.JSON(200, gin.H{
				"msg": "login successfully",
			})
		} else {
			c.JSON(403, gin.H{
				"message": "认证失败,账号密码错误",
			})
		}
	})

	r.GET("/hello", auth, func(c *gin.Context) { //在中间放入鉴权中间件
		cookie, _ := c.Get("cookie")
		str := cookie.(string) //类型断言……？应该是。
		c.String(200, "hello world"+str)
		c.Set("next", "test next") //测试next函数
	})
	r.Run(":8080")
}