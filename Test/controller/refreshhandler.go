package controller

import (
	"Test/jwt"
	"Test/model"
	"Test/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func RefreshTokenHandler(c *gin.Context){

	//rt := c.QueryArray("refresh_token")
	 authHeader := c.Request.Header.Get("Authorization")

	 if authHeader == ""{
		 c.Abort()
		 return
	 }
     parts := strings.SplitN(authHeader,"",2)
	 if !(len(parts)== 2 && parts[0] == "Bearer"){
		 c.Abort()
		 return
	 }
	 _,err := jwt.ParseToken(parts[1])
	 if err!= nil{
		 c.Abort()
		 return
	 }
	 c.Next()
}

func SingUpHandler(c *gin.Context){
	//获取参数，校验数据
	var fo *model.User
	if err := c.ShouldBindJSON(&fo);err != nil{
		return
	}
	if err := service.Login(fo.UserName,fo.Password);err != ""{
		return
	}
}
