package Mall

import (
	"github.com/gin-gonic/gin"
	Malls "mygofarm/Logics/Mall"
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
