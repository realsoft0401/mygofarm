package Mall

import (
	Mall "mygofarm/Models/Mall"
)

type EntityMallService interface {
	GetNearMall() (MallResult[] *Mall.NearMall, err error)
	GetOneNearMall(ModelHandler *Mall.GetOneMallModelHandler) (MallResult[] *Mall.NearMall, err error)
}