package User

import (
	"mygofarm/Infrastructures/Rom"
	"mygofarm/Models/User"
)

/*
获取一个用户信息
*/
type GetoneUserModelHandler struct {
	User.GetoneUserModelHandler
}

func (userModel *GetoneUserModelHandler) UserGetOneUser() (User.User, error) {
	var userResult User.User
	Rom.Db.Find(&userResult, userModel.Id)
	return userResult, nil
}

/*
获取用户登录
*/
type LoginModelHandler struct {
	User.LoginModelHandler
}

func (userModel *LoginModelHandler) UserLoginUser() (User.User, error) {
	var userResult User.User
	//fmt.Printf("username:%v\n", userModel.UserName)
	//fmt.Printf("password:%v\n", userModel.Password)
	Rom.Db.Find(&userResult, map[string]interface{}{"username": userModel.Username, "password": userModel.Password})
	return userResult, nil
}

/*
添加用户信息
*/
type SignModelHandler struct {
	User.SignModelHandler
}

func (userModel *SignModelHandler) UserSignUser() (int64, error) {
	var user User.User
	//result := Rom.Db.Select("userid", "username", "password", "email", "gender").Create(userModel)
	result := Rom.Db.Model(user).Create(map[string]interface{}{"userid": userModel.UserId,"username": userModel.Username,"password":userModel.Password ,"email": userModel.Email, "gender": userModel.Gender})
	if result.RowsAffected != 1 {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

/*
修改一个用户信息
*/
type SignUpUserHandler struct {
	User.SignModelHandler
}

func (userModel *SignUpUserHandler) UserSignUpUser() (int64, error) {
	var user User.User
	result := Rom.Db.Model(user).Updates(map[string]interface{}{"username": userModel.Username, "email": userModel.Email, "gender": userModel.Gender})
	if result.RowsAffected != 1 {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

/*
删除用户信息
*/
type UserSignDelUser struct {
	User.SignDelModelHandler
}

func (userModel *UserSignDelUser) UserSignDelUser() (int64, error) {
	var user User.User
	result := Rom.Db.Where("id", userModel.Id).Delete(user)
	//fmt.Printf("result:%v\n", result)
	if result.RowsAffected != 1 {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
