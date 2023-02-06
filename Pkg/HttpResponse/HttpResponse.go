package HttpResponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//定义
/*
{
	"code":10001,   //程序中的错误码
	"msg":xxx,      //提示信息
	"data": {},     //服务数据
}
*/
type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, c ResCode) {
	rd := &ResponseData{
		Code: c,
		Msg:  c.Msg(),
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code ResCode, errMsg string) {
	rd := &ResponseData{
		Code: code,
		Msg:  errMsg,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	ctx.JSON(http.StatusOK, rd)
}