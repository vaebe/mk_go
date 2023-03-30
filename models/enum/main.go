package enum

import "mk/models"

// Enum 枚举表
type Enum struct {
	models.BaseModel
	Name     string `gorm:"type:varbinary(100); not null; comment '枚举名称'" json:"name"`
	Value    string `gorm:"type:varbinary(100); not null; comment '枚举值'" json:"value"`
	TypeName string `gorm:"type:varbinary(100); not null; comment '枚举分类名称'" json:"typeName"`
	TypeCode string `gorm:"type:varbinary(100); not null; comment '枚举分类值'" json:"typeCode"`
	ParentId string `gorm:"type:varbinary(100); comment '枚举上级id'" json:"parentId"`
}

// EnumsForm 枚举保存表单
type EnumsForm struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Value    string `form:"value" json:"value" binding:"required"`
	TypeName string `form:"typeName" json:"typeName" binding:"required"`
	TypeCode string `form:"typeCode" json:"typeCode" binding:"required"`
	ParentId string `form:"parentId" json:"parentId"`
}

// ListForm 分页查询枚举参数
type ListForm struct {
	models.PaginationParameters
	Name     string `json:"name" form:"name"`
	TypeName string `json:"typeName" form:"typeName"`
}
