package articleColumn

import (
	"mk/models"
)

// ArticleColumn 文章专栏
type ArticleColumn struct {
	models.BaseModel
	Name         string `gorm:"type:varbinary(100);unique;not null;comment '专栏名称'" json:"name"`
	Introduction string `gorm:"type:varbinary(300);unique;not null;comment '专栏简介'" json:"introduction"`
	CoverImg     string `gorm:"type:varbinary(200);not null;comment '专栏封面'" json:"coverImg"`
	Status       string `gorm:"type:varbinary(6);not null;comment '状态 1审核中、2未通过、3已发布'" json:"status"`
	Articles     int    `gorm:"type:int;default:0;comment '文章数量'" json:"articles"`
	Subscribers  int    `gorm:"type:int;default:0;comment '订阅人数'" json:"subscribers"`
}

// SaveForm 文章专栏信息
type SaveForm struct {
	ID           int32  `json:"id" form:"id"`
	Name         string `json:"name" form:"name" binding:"required"`
	Introduction string `json:"introduction" form:"introduction" binding:"required"`
	CoverImg     string `json:"coverImg" form:"coverImg"`
}

// ListForm 获取列表
type ListForm struct {
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
}

// ReviewForm 文章专栏审核表单
type ReviewForm struct {
	ID          int32  `json:"id" form:"id"`
	Status      string `json:"status" form:"status" example:"2"` // 2未通过、3已发布
	Description string `json:"description" form:"description"`   // 审核意见
}
