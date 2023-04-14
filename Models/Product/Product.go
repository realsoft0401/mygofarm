package Product

type Product struct {
	Pid          int64   `gorm:"column:pid;unique"`
	Mid          int64   `gorm:"column:mid;unique"`
	ProductName  string  `gorm:"column:product_name"`
	ProductClass string  `gorm:"column:product_class"`
	ImgUrl       string  `gorm:"column:imgUrl"`
	Sales        int64   `gorm:"column:sales"`
	Price        float64 `gorm:"column:price"`
	OldPrice     float64 `gorm:"column:oldprice"`
}

func (Product) TableName() string {
	return "product"
}
