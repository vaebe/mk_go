package models

import "time"

// BaseModel 基础数据
type BaseModel struct {
	ID        int32     `gorm:"primaryKey;comment '主键'"`
	CreatedAt time.Time `gorm:"column=add_time;comment '创建时间'"`
	UpdatedAt time.Time `gorm:"column=update_time;comment '更新时间'"`
	IsDeleted bool      `gorm:"column=is_deleted;comment '是否删除'"`
}
