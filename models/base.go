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

// PaginationParameters 分页参数
type PaginationParameters struct {
	PageSize int `json:"pageSize" form:"pageSize" example:"10" binding:"required,min=0"`
	PageNo   int `json:"pageNo" form:"pageNo" example:"1" binding:"required,min=0"`
}

// PagingData 分页数据对象
type PagingData struct {
	List     any   `json:"list"`
	PageSize int   `json:"pageSize" `
	PageNo   int   `json:"pageNo" `
	Total    int32 `json:"total"`
}
