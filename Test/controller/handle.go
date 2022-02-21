package controller

import (
	"Test/logic"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/model"
	"test/service"
)

func Menu(){
	//登陆
	v := gin.New()
	v1 := v.Group("/user")
	var u *model.LoginForm
	v1.GET("/refresh_token")
	v1.GET("/user/token", func(c *gin.Context) {
		u.UserName = c.Query("name")
		u.Password = c.Query("password")

		if err := c.ShouldBindJSON(&u);err != nil{
			return
		}

		//atoken,_,err := logic.Login(u)
		//if err != nil{
			return
		//}


		L := service.Login(u.UserName,u.Password)
		if L ==""{
			c.String(http.StatusFailedDependency,"登陆失败")
		}
		if L!=""{
			c.String(http.StatusOK,L)
		}
	})

	//修改密码
	r := gin.Default()
	r.PUT("/user/password", func(c *gin.Context) {
		newPassword:=c.PostForm("newPassword")
		oldPassword:=c.PostForm("oldPassword")
		B := service.ChangePassword(oldPassword,newPassword)
		if B == ""{
			c.String(http.StatusFailedDependency,"修改失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//注册
	r.POST("/user/register", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		introduction := c.PostForm("introduction")
		emil := c.PostForm("emil")
		phone := c.GetInt("phone")
		qq := c.GetInt("qq")
		gender := c.PostForm("gender")
		birth := c.PostForm("birth")

		//请求获取参数、校验数据有效性
		var fo *model.User
		if err := c.ShouldBindJSON(&fo); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})

		logic.SignUp(name,password,emil,introduction,phone,qq,gender,birth)

	})
	//查询资料
	r.GET("/user/info/{user_id}", func(c *gin.Context) {
		name := c.Query("user_id")
		B := service.QueryInformation(name)
		if B == ""{
			c.String(http.StatusFailedDependency,"查询失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//修改资料接口
	r.PUT("/user/info", func(c *gin.Context) {
		nickname := c.PostForm("nickname")
		introduction := c.PostForm("introduction")
		emil := c.PostForm("emil")
		phone := c.GetInt("phone")
		qq := c.GetInt("qq")
		gender := c.PostForm("gender")
		birth := c.PostForm("birth")
		B := service.ChangeInformation(nickname, emil, introduction, phone, qq, gender, birth)
		if B == ""{
			c.String(http.StatusFailedDependency,"修改失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//发起话题接口
	r.POST("/post/single", func(c *gin.Context) {
		content := c.PostForm("content")
		topicName := c.PostForm("topicName")
		B := service.InsertTopic(content,topicName)
		if B==false{
			c.String(http.StatusFailedDependency,"加入话题失败")
		}
		if B==true{
			c.String(http.StatusOK,"加入话题成功")
		}
	})
	//展示所有话题
	r.GET("/topic/list", func(c *gin.Context) {
		v := model.Ioutil("topicName.txt")
		c.String(http.StatusOK,v)
	})
	//展示话题接口
	r.GET("/post/single/{post_id}", func(c *gin.Context) {
		topicName := c.Query("topicName")
		B := service.ShowTopic(topicName)
		if B == ""{
			c.String(http.StatusFailedDependency,"未找到标题")
		}
		if B!=""{
			c.String(http.StatusOK,B)
		}
	})
	//评论话题
	r.GET("/comment", func(c *gin.Context) {
		topicName := c.Query("topicName")
		message := c.Query("message")
		B := service.ShowTopic(topicName)
		if B == ""{
			c.String(http.StatusFailedDependency,"未找到标题")
		}
		if B!=""{
			c.String(http.StatusOK,B)
		}
		C := service.InsertComment(topicName,message)
		if C==""{
			c.String(http.StatusFailedDependency,"未找到标题")
		}else {
			c.String(http.StatusOK,C)
		}
	})
	//更新话题
	r.PUT("/post/single/{post_id}", func(c *gin.Context) {
		name:=c.PostForm("topicName")
		content:=c.PostForm("content")
		title := c.PostForm("title")
		C := service.ChangeTopicName(name,content,title)
		if C=="" {
			c.String(http.StatusFailedDependency,"修改失败")
		}
		if C!=""{
			c.String(http.StatusOK,C)
		}
	})
	r.DELETE("/post/single/{post_id}", func(c *gin.Context) {
		id := c.Param("post_id")
		B := service.DeleteTopic(id)
		if B == ""{
			c.String(http.StatusFailedDependency,"删除失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//点赞评论
	r.PUT("/operate/praise", func(c *gin.Context) {
		name := c.PostForm("topicName")
		id := c.PostForm("id")
		B := service.GetLike(name,id)
		if B == ""{
			c.String(http.StatusFailedDependency,"点赞失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//关注用户
	r.PUT("/operate/focus", func(c *gin.Context) {
		id := c.PostForm("user_id")
		B := service.FollowUser(id)
		if B == ""{
			c.String(http.StatusFailedDependency,"关注失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	//收藏话题
	r.PUT("/operate/collect", func(c *gin.Context) {
		id := c.PostForm("post_id")
		B := service.FollowUser(id)
		if B == ""{
			c.String(http.StatusFailedDependency,"收藏失败")
		}
		if B != ""{
			c.String(http.StatusOK,B)
		}
	})
	r.Run(":8000")
}