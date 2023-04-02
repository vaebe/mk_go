package article

import "mk/models"

// Article 文章
type Article struct {
	models.BaseModel
	UserId           int32  `gorm:"type:int;not null;comment '用户id'" json:"userId"`
	Title            string `gorm:"type:varbinary(100);unique;not null;comment '文章标题'" json:"title"`
	Content          string `gorm:"type:longtext;not null;comment '文章内容'" json:"content"`
	Classify         string `gorm:"type:varbinary(200);not null;comment '文章分类'" json:"classify"`
	Tags             string `gorm:"type:varbinary(300);not null;comment '文章标签'" json:"tags"`
	CoverImg         string `gorm:"type:varbinary(200);not null;comment '文章封面'" json:"coverImg"`
	CollectionColumn string `gorm:"type:varbinary(300);not null;comment '收录专栏'" json:"collectionColumn"`
	Summary          string `gorm:"type:varbinary(300);not null;comment '摘要'" json:"summary"`
	Views            int    `gorm:"type:int;default:1;comment '阅读数'" json:"views"`
	Likes            int    `gorm:"type:int;default:1;comment '点赞数'" json:"likes"`
	Favorites        int    `gorm:"type:int;default:1;comment '收藏数'" json:"favorites"`
	ShowNumber       int    `gorm:"type:int;default:1;comment '展现数'" json:"showNumber"`
	Status           string `gorm:"type:varbinary(6);not null;comment '状态 1草稿 2待审核 3审核未通过 4已发布 5已删除'" json:"status"`
}

// ArticlesRelatedTags 文章关联的标签
type ArticlesRelatedTags struct {
	models.BaseModel
	ArticleId int32  `gorm:"type:int;not null;comment '文章id'" json:"articleId"`
	TagId     string `gorm:"type:varbinary(100);;not null;comment '文章标签id'" json:"tagId"`
}

// ArticlesAssociatedColumns 文章关联的专栏
type ArticlesAssociatedColumns struct {
	models.BaseModel
	ArticleId int32 `gorm:"type:int;not null;comment '文章id'" json:"articleId"`
	ColumnId  int32 `gorm:"type:int;not null;comment '专栏id'" json:"columnId"`
}

// SaveForm 文章保存表单
type SaveForm struct {
	ID               int32    `json:"id" form:"id"`
	UserId           int32    `json:"userId" form:"userId" binding:"required"`
	Title            string   `json:"title" form:"title"`
	Content          string   `json:"content" form:"content" binding:"required"`
	Classify         string   `json:"classify" form:"classify"`
	Tags             []string `json:"tags" form:"tags" swaggertype:"array,string"`
	CoverImg         string   `json:"coverImg" form:"coverImg"`
	CollectionColumn []int32  `json:"collectionColumn" form:"collectionColumn" swaggertype:"array,int32"`
	Summary          string   `json:"summary" form:"summary"`
}

// AllListForm 获取全部文章列表
type AllListForm struct {
	models.PaginationParameters
	Title    string `json:"title" form:"title"`
	Classify string `json:"classify" form:"classify"`
	Tag      string `json:"tag" form:"tag"`
	Status   string `json:"status" form:"status" example:"1"` // 1草稿 2待审核 3审核未通过 4已发布 5已删除
}

// UserArticleListForm 获取文章列表
type UserArticleListForm struct {
	models.PaginationParameters
	Status string `json:"status" form:"status" example:"1"` // 1草稿 2待审核 3审核未通过 4已发布
}

// ReviewForm 文章审核表单
type ReviewForm struct {
	ID          int32  `json:"id" form:"id"`
	Status      string `json:"status" form:"status" example:"3"` // 3驳回 4通过
	Description string `json:"description" form:"description"`   // 审核意见
}
