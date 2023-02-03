package User

import "time"

type User struct {
	Id         int64     `gorm:"column:id;unique"`
	Userid     string    `gorm:"column:userid"`
	UserName   string    `gorm:"column:username"`
	Gender     string    `gorm:"column:Gender"`
	Createtime time.Time `gorm:"column:create_time;type:date"`
	Updatetime time.Time `gorm:"column:update_time;type:date"`
}

func (User) TableName() string {
	return "user"
}