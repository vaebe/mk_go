package article

import "mk/models"

// Column 文章专栏
type Column struct {
	models.BaseModel
	Name         string `gorm:"type:varbinary(50);unique;not null;comment '专栏名称'" json:"name"`
	Introduction string `gorm:"type:varbinary(300);unique;not null;comment '专栏简介'" json:"introduction"`
	CoverImg     string `gorm:"type:varbinary(200);not null;comment '专栏封面'" json:"coverImg"`
	Status       string `gorm:"type:varbinary(6);not null;comment '状态 已发布、审核中、未通过'" json:"status"`
	Articles     int    `gorm:"type:int;default:1;comment '文章数量'" json:"articles"`
	Subscribers  int    `gorm:"type:int;default:1;comment '订阅人数'" json:"subscribers"`
}
