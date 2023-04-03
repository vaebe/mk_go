package articleAssociatedInfo

import "mk/models"

// ArticlesAssociatedColumns 文章关联的专栏
type ArticlesAssociatedColumns struct {
	models.BaseModel
	UserId    int32 `gorm:"type:int;not null;comment '用户id'" json:"userId"`
	ArticleId int32 `gorm:"type:int;not null;comment '文章id'" json:"articleId"`
	ColumnId  int32 `gorm:"type:int;not null;comment '专栏id'" json:"columnId"`
}

// ArticlesAssociatedColumnsForm 文章关联专栏表单
type ArticlesAssociatedColumnsForm struct {
	ArticleId int32 `json:"articleId" form:"articleId" binding:"required"`
	ColumnId  int32 `json:"columnId" form:"columnId" binding:"required"`
}

// ArticlesRelatedTags 文章关联的标签
type ArticlesRelatedTags struct {
	models.BaseModel
	UserId    int32  `gorm:"type:int;not null;comment '用户id'" json:"userId"`
	ArticleId int32  `gorm:"type:int;not null;comment '文章id'" json:"articleId"`
	TagId     string `gorm:"type:varbinary(100);;not null;comment '文章标签id'" json:"tagId"`
}
