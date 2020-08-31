package user

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
	Email           string `json:"email" form:"email"`
}

// LoginCredentials 默认登录方式-邮箱
type LoginCredentials struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// PhoneLoginCredentials 手机登录
type PhoneLoginCredentials struct {
	Phone      int `json:"phone" form:"phone" binding:"required" example:"13010002000"`
	VerifyCode int `json:"verify_code" form:"verify_code" binding:"required" example:"120110"`
}


// UpdateRequest 更新请求
type UpdateRequest struct {
	Avatar string `json:"avatar" form:"avatar"`
	Sex    int    `json:"sex" form:"sex"`
}