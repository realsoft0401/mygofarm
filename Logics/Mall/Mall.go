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