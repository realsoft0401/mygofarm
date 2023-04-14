package Mall

import (
	"mygofarm/Infrastructures/Rom"
	"mygofarm/Models/Mall"
)

/*
获取附近门店信息
*/
func GetNearMall() ( []Mall.NearMall , error) {
	var nearmall[] Mall.NearMall
	Rom.Db.Find(&nearmall)
	return nearmall, nil
}

/*
获取一家门店信息
*/
type GetoneMallModelHandler struct {
	Mall.GetOneMallModelHandler
}

func (mallModel *GetoneMallModelHandler) GetOneNearMall() (Mall.NearMall, error) {
	var mallResult Mall.NearMall
	Rom.Db.Find(&mallResult, mallModel.MId)
	return mallResult, nil
}
