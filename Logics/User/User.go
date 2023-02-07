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
	Rom.Db.Find(&userResult, userModel.UserName,userModel.Password)
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
	result := Rom.Db.Model(userModel).Update("username", "gender")
	//result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
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
