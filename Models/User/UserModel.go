package User

type LoginModelHandler struct {
	// Username description
	Username string `json:"username"`
	// Password description
	Password string `json:"password"`
}

type SignModelHandler struct {
	// UserId 用户ID
	UserId int64 `json:"userid"`
	// Username 用户名称
	Username string `json:"username"`
	// Password 用户密码
	Password string `json:"password"`
	// Email 用户email
	Email string `json:"email"`
	// Gender 用户性别
	Gender string `json:"gender"`
	// RePassword 二次密码验证
	RePassword string `json:"repassword"`
}

type SignDelModelHandler struct {
	// Id 用户主键ID
	Id int64 `json:"id"`
}

type GetoneUserModelHandler struct {
	// Id 用户主键ID
	Id int64 `json:"id"`
}
