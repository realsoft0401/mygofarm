package Product

import (
	"mygofarm/Infrastructures/Rom"
	"mygofarm/Models/Product"
)


/*
获取门店商品信息
*/
type GetMallProductModelHandler struct {
	Product.GetMallProductModelHandler
}

//获取所有商品信息
func (ProductModel *GetMallProductModelHandler) GetAllProduct() ([]Product.Product, error) {
	var ProdResult[] Product.Product
	Rom.Db.Where("mid=?",ProductModel.Mid).Find(&ProdResult)
	return ProdResult, nil
}

//获取某种类型的商品信息
func (ProductModel *GetMallProductModelHandler) GetProdClassProduct() ([]Product.Product, error) {
	var ProdResult[] Product.Product
	Rom.Db.Where("mid=? and product_class=?",ProductModel.Mid,ProductModel.ProductClass).Find(&ProdResult)
	return ProdResult, nil
}
