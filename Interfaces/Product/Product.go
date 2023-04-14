package Product

import (
	product "mygofarm/Models/Product"
)

type EntityMallService interface {
	GetMallProduct(ModelHandler *product.GetMallProductModelHandler) (ProductResult[] *product.Product, err error)
	//GetOneNearMall(ModelHandler *Mall.GetOneMallModelHandler) (MallResult[] *Mall.NearMall, err error)
}
