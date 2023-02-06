package User

import (
	"errors"
	Logicsuser "mygofarm/Logics/User"
	"mygofarm/Models/User"
	"mygofarm/Pkg/HttpResponse"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignHandler(c *gin.Context) {
	p := new(User.User)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeFormatError, "请求参数格式错误")
		return
	}
	//手动的业务请求参数的规则校验
	if len(p.UserName) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeInvalidParams, "参数不能为空或者密码与校验密码不相同")
		return
	}

	puser := &Logicsuser.UserModel{User.User{p.Id, p.Userid, p.Password, p.UserName, p.Email, p.Gender, p.RePassword}}
	if _, err := puser.UserSignUser(); err != nil {
		if errors.Is(err, ErrorUserExit) {
			HttpResponse.ResponseError(c, HttpResponse.CodeUserExist)
			return
		}
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)

		return
	}
	//3.返回响应
	HttpResponse.ResponseSuccess(c, nil)
}
