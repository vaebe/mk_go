package models

import (
	"time"
)

// BaseModel 基础数据
type BaseModel struct {
	ID        int32     `gorm:"primaryKey;comment '主键'"`
	CreatedAt time.Time `gorm:"column=add_time;comment '创建时间'"`
	UpdatedAt time.Time `gorm:"column=update_time;comment '更新时间'"`
	IsDeleted bool      `gorm:"column=is_deleted;comment '是否删除'"`
}

type UserTest struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varbinary(11);not null;comment '手机号'"`
	Password string     `gorm:"type:varbinary(100);not null;comment '密码'"`
	NickName string     `gorm:"type:datetime;comment '昵称'"`
	Birthday *time.Time `gorm:"type:varbinary(20);comment '生日'"`
	Gender   string     `gorm:"column:gender;default:male;type:varbinary(6);comment 'female 女 male 男'"`
	Role     int        `gorm:"column:role;default:1;type:int;comment '1普通用户 2管理员'"`
	Desc     string     `gorm:"column:desc;type:text;comment '描述'"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}

type VerificationCodeForm struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

type User struct {
	BaseModel
	NickName              string `gorm:"type:varbinary(20);unique;not null;comment '昵称'" json:"nickName"`
	UserAvatar            string `gorm:"type:varbinary(300);not null;comment '用户头像'" json:"userAvatar"`
	UserName              string `gorm:"type:varbinary(50);unique;not null;comment '用户名'" json:"userName"`
	UserAccount           string `gorm:"type:varbinary(50);unique;not null;comment '用户账号'" json:"userAccount"`
	Password              string `gorm:"type:varbinary(100);not null;comment '密码'" json:"password"`
	Github                string `gorm:"type:varbinary(100);not null;comment 'github账户'" json:"github"`
	Posts                 string `gorm:"type:varbinary(200);not null;comment '职位'" json:"posts"`
	Company               string `gorm:"type:varbinary(200);not null;comment '所在公司'" json:"company"`
	Homepage              string `gorm:"type:varbinary(300);not null;comment '个人主页'" json:"homepage"`
	PersonalProfile       string `gorm:"type:varbinary(300);not null;comment '个人简介'" json:"personalProfile"`
	UserCreationPoints    string `gorm:"type:int;not null;default:0;comment '用户创作积分'" json:"userCreationPoints"`
	UserInteractionPoints string `gorm:"type:int;not null;default:0;comment '用户互动积分'" json:"userInteractionPoints"`
}
