package models

import (
	"time"
)

// BaseModel 基础数据
type BaseModel struct {
	ID int32 `gorm:"primaryKey"`
	//CreatedAt time.Time `gorm:"column=add_time"`
	//UpdatedAt time.Time `gorm:"column=update_time"`
	//IsDeleted bool
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
