package Mall

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	Malls "mygofarm/Logics/Mall"
	MallModels "mygofarm/Models/Mall"
	"mygofarm/Pkg/HttpResponse"
)

// 附近店铺数据
// @Summary  附近店铺数据接口
// @Tags 附近店铺数据相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} Response._ResponsePostList
// @Router /api/nearmall [get]
func GetNearMallHandler(c *gin.Context) {
	var result, err = Malls.GetNearMall()
	if err != nil {
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)
		return
	}
	//3.返回响应
	HttpResponse.ResponseSuccess(c, result)
}
// 获取单个商户信息接口
// @Summary 获取单个商户信息接口
// @Description 商户主键ID
// @Tags 获取单个用户信息接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param req body Mall.GetOneMallModelHandler true "req"
// @  type: Mall.GetOneMallModelHandler
// @Security ApiKeyAuth
// @Success 200 {object} Response._ResponsePostList
// @Router /api/getmallone [post]
func GetoneMallHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(Malls.GetoneMallModelHandler)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("Getone with invalid param", zap.Error(err))
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeFormatError, "请求参数格式错误")
		return
	}
	Malls := &Malls.GetoneMallModelHandler{MallModels.GetOneMallModelHandler{MId: p.MId}}
	var result, err = Malls.GetOneNearMall()
	if err != nil {
		if errors.Is(err, ErrorMallNotExit) {
			HttpResponse.ResponseError(c, HttpResponse.CodeMallNotExist)
			return
		}
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)
		return
	}
	if  result.MId  == 0 {
		HttpResponse.ResponseError(c, HttpResponse.CodeMallNotExist)
		return
	}
	//3.返回响应
	HttpResponse.ResponseSuccess(c, result)

}
