package articleAssociatedInfo

import "mk/models"

// ArticlesAssociatedColumns 文章关联的专栏
type ArticlesAssociatedColumns struct {
	models.BaseModel
	ArticleId int32 `gorm:"type:int;not null;comment '文章id'" json:"articleId"`
	ColumnId  int32 `gorm:"type:int;not null;comment '专栏id'" json:"columnId"`
}

// ArticlesRelatedTags 文章关联的标签
type ArticlesRelatedTags struct {
	models.BaseModel
	ArticleId int32  `gorm:"type:int;not null;comment '文章id'" json:"articleId"`
	TagId     string `gorm:"type:varbinary(100);;not null;comment '文章标签id'" json:"tagId"`
}
