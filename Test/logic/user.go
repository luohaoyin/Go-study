package logic

import (
    "Test/dao"
    "Test/jwt"
    "Test/model"
    "Test/service"
    "errors"
)


func SignUp(name string,password string,emil string,introduction string,phone int,qq int,
    gender string,birth string)(error error){
   //判断用户是否存在
   err := dao.CheckUserExist(name)
   if err != nil {
       //数据库查询出错
      return err
   }
    return
}

func Login(p *model.LoginForm)(atoken,rtoken string,error error) {
    user := &model.User{
        UserName: p.UserName,
        Password: p.Password,
    }
    if err := service.Login(user.UserName, user.Password); err != "" {
        return "", "",errors.New("")
    }
    //生成JET
    return jwt.GenToken(user.UserName)
}