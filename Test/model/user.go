package model

import (
	"encoding/json"
	"errors"
)

type User struct {
	UserName string `db:user_name`
	Password string `db:password`
	Emil string `db:emil`
	Introduction string `db:introduction`
	Phone int `db:phone`
	QQ int `db:qq`
	Gender string `db:gender`
	Birth string `db:birth`
}
type UserInfo struct {
	UserName string `db:user_name`
	Topic string `db:topic`
}
type LoginForm struct {
	UserName string    `json:"username"binding:"required"`
 	Password  string   `json:"password"binding:"required"`

}
   //验证一下是否有填密码、以及密码是否一致
func (r *User) UnmarshalJson(data []byte)(err error){
	required := struct {
		UserName    string  `json:"username"binding:"required"`
		Password    string  `json:"password"bindig:"required"`
		ConfirmPassword  string   `json:"confirm_password"`
	}{}
    err = json.Unmarshal(data, &required)
	if err!= nil{} else if len(required.UserName) == 0{
		err = errors.New("必须写用户名！")
	}else if len(required.Password) == 0{
		err = errors.New("密码也得写哎！")
	}else if required.ConfirmPassword != required.Password {
		err = errors.New("两次密码不一致诶~")
		return
	}else{
		r.UserName = required.UserName
		r.Password = required.Password
	}
    return
}
