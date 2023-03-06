package User

import (
	"errors"
	Logicuser "mygofarm/Logics/User"
	"mygofarm/Models/User"
	"mygofarm/Pkg/HttpResponse"
	"mygofarm/Pkg/Jwt"
	"mygofarm/Services/Public"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 用户登录接口
// @Summary 用户登录接口
// @Description 用户根据用户名、密码登录
// @Tags 用户登录相关接口
// @Accept application/json
// @Produce application/json
// @Param req body User.LoginModelHandler true "req"
// @	type: User.LoginModelHandler
// @Success 200 {object} Response._ResponsePostList
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(User.LoginModelHandler)
	if err := c.ShouldBindJSON(&p); err != nil {
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeInvalidParams, "参数格式错误")
		return
	}
	if len(p.Username) == 0 || len(p.Password) == 0 {
		//请求参数有误
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeInvalidParams, "参数不能为空")
		return
	}
	PW := Public.EncryptPassword(p.Password)
	LogicUser := &Logicuser.LoginModelHandler{User.LoginModelHandler{
		Username: p.Username,
		Password: PW,
	}}
	var result, err = LogicUser.UserLoginUser()
	if err != nil {
		if errors.Is(err, ErrorUserExit) {
			HttpResponse.ResponseError(c, HttpResponse.CodeUserExist)
			return
		}
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)
		return
	}
	if result.Id == 0 {
		//请求参数有误
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeInvalidParams, "帐号或者密码错误")
		return
	}
	var tokenString, _ = Jwt.GenToken(result.UserId, result.UserName)
	//3.返回响应
	HttpResponse.ResponseSuccess(c, gin.H{"token": tokenString})

}

// 用户注册接口
// @Summary 用户注册接口
// @Description 用户ID、名称、密码、性别、Emaill
// @Tags 用户注册接口
// @Accept application/json
// @Produce application/json
// @Param req body User.SignModelHandler true "req"
// @  type: User.SignModelHandler
// @Success 200 {object} Response._ResponsePostList
// @Router /signadd [post]
func SignHandler(c *gin.Context) {
	p := new(User.User)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("Sign with invalid param", zap.Error(err))
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeFormatError, "请求参数格式错误")
		return
	}
	//手动的业务请求参数的规则校验
	if len(p.UserName) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeInvalidParams, "参数不能为空或者密码与校验密码不相同")
		return
	}
	///
	///生成雪花ID生成user_id
	///
	node, snowflakeerr := snowflake.NewNode(1)
	if snowflakeerr != nil {
		HttpResponse.ResponseError(c, HttpResponse.SnowFlakeError)
	}
	// Generate a snowflake ID.
	UserId := node.Generate().Int64()

	password := Public.EncryptPassword(p.Password)
	LogicUser := &Logicuser.SignModelHandler{User.SignModelHandler{
		UserId:     UserId,
		Username:   p.UserName,
		Password:   password,
		Email:      p.Email,
		Gender:     p.Gender,
		RePassword: password,
	}}
	var result, err = LogicUser.UserSignUser()
	if err != nil {
		if errors.Is(err, ErrorUserExit) {
			HttpResponse.ResponseError(c, HttpResponse.CodeUserExist)
			return
		}
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)
		return
	}
	//3.返回响应
	HttpResponse.ResponseSuccess(c, result)
}

// 用户删除接口
// @Summary 用户删除接口
// @Description 用户主键ID
// @Tags 用户删除接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param req body User.SignDelModelHandler true "req"
// @  type: User.SignDelModelHandler
// @Security ApiKeyAuth
// @Success 200 {object} Response._ResponsePostList
// @Router /signdel [post]
func SignDelHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(User.SignDelModelHandler)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("SignDel with invalid param", zap.Error(err))
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeFormatError, "请求参数格式错误")
		return
	}
	LogicUser := &Logicuser.UserSignDelUser{User.SignDelModelHandler{Id: p.Id}}
	var result, err = LogicUser.UserSignDelUser()
	if err != nil {
		if errors.Is(err, ErrorUserExit) {
			HttpResponse.ResponseError(c, HttpResponse.CodeUserExist)
			return
		}
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)
		return
	}
	//3.返回响应
	HttpResponse.ResponseSuccess(c, result)
}

// 获取单个用户信息接口
// @Summary 获取单个用户信息接口
// @Description 用户主键ID
// @Tags 获取单个用户信息接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param req body User.GetoneUserModelHandler true "req"
// @  type: User.GetoneUserModelHandler
// @Security ApiKeyAuth
// @Success 200 {object} Response._ResponsePostList
// @Router /signone [post]
func GetoneUserHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(User.SignDelModelHandler)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("Getone with invalid param", zap.Error(err))
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeFormatError, "请求参数格式错误")
		return
	}
	LogicUser := &Logicuser.GetoneUserModelHandler{User.GetoneUserModelHandler{Id: p.Id}}
	var result, err = LogicUser.UserGetOneUser()
	if err != nil {
		if errors.Is(err, ErrorUserExit) {
			HttpResponse.ResponseError(c, HttpResponse.CodeUserExist)
			return
		}
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)
		return
	}
	//3.返回响应
	HttpResponse.ResponseSuccess(c, result)

}
