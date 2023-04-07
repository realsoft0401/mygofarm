package Mall

import (
	Mall "mygofarm/Models/Mall"
)

type EntityMallService interface {
	GetNearMall() (MallResult *Mall.NearMall, err error)
}