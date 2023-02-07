package User

import (
	"mygofarm/Infrastructures/Rom"
	"mygofarm/Models/User"
)

type UserModel struct {
	User.User
}

func (userModel *UserModel) UserGetOneUser() (User.User, error) {
	var userResult User.User
	Rom.Db.Find(&userResult, userModel.Id)
	return userResult, nil
}

func (userModel *UserModel) UserLoginUser() (User.User, error) {
	var userResult User.User
	//fmt.Printf("username:%v\n", userModel.UserName)
	//fmt.Printf("password:%v\n", userModel.Password)
	Rom.Db.Find(&userResult, map[string]interface{}{"username": userModel.UserName, "password": userModel.Password})
	return userResult, nil
}

func (userModel *UserModel) UserSignUser() (int64, error) {
	result := Rom.Db.Select("userid", "username", "password", "email", "gender").Create(userModel)
	if result.RowsAffected != 1 {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (userModel *UserModel) UserSignUpUser() (int64, error) {
	result := Rom.Db.Model(userModel).Updates(map[string]interface{}{"username": userModel.UserName, "email": userModel.Email, "gender": userModel.Gender})
	if result.RowsAffected != 1 {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (userModel *UserModel) UserSignDelUser() (int64, error) {
	result := Rom.Db.Where("id", userModel.Id).Delete(userModel)
	if result.RowsAffected != 1 {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
