package user

import "mk/models"

type User struct {
	models.BaseModel
	UserId                string `gorm:"type:varbinary(60);unique;not null;comment '用户id'" json:"userId"`
	NickName              string `gorm:"type:varbinary(40);unique;not null;comment '昵称'" json:"nickName"`
	UserAvatar            string `gorm:"type:varbinary(300);not null;comment '用户头像'" json:"userAvatar"`
	UserName              string `gorm:"type:varbinary(50);unique;comment '用户名'" json:"userName"`
	UserAccount           string `gorm:"type:varbinary(50);unique;not null;comment '用户账号'" json:"userAccount"`
	Password              string `gorm:"type:varbinary(300);not null;comment '密码'" json:"password"`
	Github                string `gorm:"type:varbinary(100);comment 'github账户'" json:"github"`
	Posts                 string `gorm:"type:varbinary(200);comment '职位'" json:"posts"`
	Role                  int    `gorm:"column:role;default:1;type:int;comment '1普通用户 2管理员'" json:"role"`
	Company               string `gorm:"type:varbinary(200);comment '所在公司'" json:"company"`
	Homepage              string `gorm:"type:varbinary(300);comment '个人主页'" json:"homepage"`
	PersonalProfile       string `gorm:"type:varbinary(300);comment '个人简介'" json:"personalProfile"`
	UserCreationPoints    string `gorm:"type:int;not null;default:0;comment '用户创作积分'" json:"userCreationPoints"`
	UserInteractionPoints string `gorm:"type:int;not null;default:0;comment '用户互动积分'" json:"userInteractionPoints"`
}

// VerificationCodeForm 发送验证码
type VerificationCodeForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

// RegisterForm 注册
type RegisterForm struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}

// LoginForm 登陆
type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,email" example:"mk@163.com"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20" example:"123456"`
}

// ListForm 获取用户列表查询参数
type ListForm struct {
	models.PaginationParameters
	Email    string `json:"email" form:"email" `
	NickName string `json:"nickName" form:"nickName" `
}
