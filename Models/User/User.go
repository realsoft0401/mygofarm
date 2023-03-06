package User

type User struct {
	Id         int64  `gorm:"column:id;unique"`
	UserId     int64  `gorm:"column:userid"`
	Password   string `gorm:"column:password"`
	UserName   string `gorm:"column:username"`
	Email      string `gorm:"column:email"`
	Gender     string `gorm:"column:Gender"`
	RePassword string
}

func (User) TableName() string {
	return "user"
}
