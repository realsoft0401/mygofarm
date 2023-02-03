package User

import (
	"mygofarm/Infrastructures/Rom"
	"mygofarm/Models/User"
)
type UserParams struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"user_name"`
	Password string `db:"password"`
}


func GetUserFindAll() ([]User.User,error){
	var users []User.User
	Rom.Db.Find(&users)
	return users,nil
}

func AddUser(user *User.User) (int64, error){
	result := Rom.Db.Omit("userid","username","gender").Create(user)
	if result.RowsAffected !=1 {
		return 0,result.Error
	}
	return result.RowsAffected,nil
}

func UpUser(user *User.User) (int64, error)  {
	result :=Rom.Db.Model(user).Update("username")
	//result := db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	if result.RowsAffected !=1 {
		return 0,result.Error
	}
	return result.RowsAffected,nil
}

func DelUser(user *User.User) (int64, error) {
	result :=Rom.Db.Where("id",user.Id).Delete(user)
	if result.RowsAffected !=1 {
		return 0,result.Error
	}
	return result.RowsAffected,nil
}