package Product

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mygofarm/Logics/Product"
	ProductModels "mygofarm/Models/Product"
	"mygofarm/Pkg/HttpResponse"
)

// 获取商户商品信息接口
// @Summary 获取单个商户信息接口
// @Description 商户主键ID 商品类型主键
// @Tags 获取商户商品信息接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param req body Product.GetMallProductModelHandler true "req"
// @  type: Product.GetMallProductModelHandler
// @Security ApiKeyAuth
// @Success 200 {object} Response._ResponsePostList
// @Router /api/getmallprod [post]
func GetMallProdHandler(c *gin.Context) {
	var result[] ProductModels.Product
	var err error
	//1.获取参数和参数校验
	p := new(Product.GetMallProductModelHandler)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("Getone with invalid param", zap.Error(err))
		HttpResponse.ResponseErrorWithMsg(c, HttpResponse.CodeFormatError, "请求参数格式错误")
		return
	}
	Product := &Product.GetMallProductModelHandler{ProductModels.GetMallProductModelHandler{Mid: p.Mid,ProductClass: p.ProductClass} }
	if string(p.ProductClass) == "all" {
		 result, err = Product.GetAllProduct()
	}else {
		 result, err = Product.GetProdClassProduct()
	}
	if err != nil {
		HttpResponse.ResponseError(c, HttpResponse.CodeServerBusy)
		return
	}
	//3.返回响应
	HttpResponse.ResponseSuccess(c, result)
}