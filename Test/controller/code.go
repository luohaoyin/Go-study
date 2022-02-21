package controller

type  Mycode int64

const (
	CodeSuccess            Mycode = 1000
	CodeInvalidParams      Mycode = 1001
	CodeUserExist          Mycode = 1002
	CodeUserENotExist      Mycode = 1003
	CodeInvalidPassword    Mycode = 1004
	CodeServerBusy         Mycode = 1005

	CodeInvalidToken        Mycode = 1006
	CodeInvalidAuthFormat   Mycode = 1007
)
var msglags = map[Mycode]string{
	CodeSuccess:            "success",
	CodeInvalidParams:      "请求参数错误",
	CodeUserExist:          "用户名重复",
	CodeUserENotExist:      "用户不存在",
	CodeInvalidPassword:    "用户名或密码错误",
	CodeServerBusy:         "服务繁忙",

	CodeInvalidToken:        "无效的Token",
	CodeInvalidAuthFormat:   "认证格式有误",
}

func (c Mycode) Msg() string{
	msg ,ok := msglags[c]
	if ok{
		return msg
	}
	return msglags[CodeServerBusy]
}