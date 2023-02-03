package User

import (
	"mygofarm/Models/User"
)

//建立用户实体关系接口
type EntityUserService interface {
	//添加用户信息
	//返回参数为更新记录数，报错信息
	UserSignUser(user *User.User) (result int64,err error)
	//修改用户信息
	//返回参数为删除记录数，报错信息
	UserSignUpUser(user *User.User) (result int64,err error)
	//删除用户信息
	//返回参数为删除记录数，报错信息
	UserSignDelUser(user *User.User) (result int64,err error)
}
