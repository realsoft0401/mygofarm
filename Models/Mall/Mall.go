package Mall

type NearMall struct {
	MId      int64   `gorm:"column:mid;unique"`
	MallName string  `gorm:"column:mallname"`
	ImgUrl   string  `gorm:"column:imgUrl"`
	Sales    int64   `gorm:"column:sales"`
	Limt     int64   `gorm:"column:limt"`
	Price    float64 `gorm:"column:price"`
	Slogan   string  `gorm:"column:slogan"`
}

func (NearMall) TableName() string {
	return "nearmall"
}
