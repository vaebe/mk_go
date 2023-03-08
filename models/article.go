package models

type article struct {
	BaseModel
	Title            string `gorm:"type:varbinary(100);unique;not null;comment '文章标题'" json:"title"`
	Content          string `gorm:"type:text;not null;comment '文章内容'" json:"content"`
	Classify         string `gorm:"type:varbinary(200);not null;comment '文章分类'" json:"classify"`
	Tag              string `gorm:"type:varbinary(300);not null;comment '文章标签'" json:"tag"`
	CoverImg         string `gorm:"type:varbinary(200);not null;comment '文章封面'" json:"coverImg"`
	CollectionColumn string `gorm:"type:varbinary(300);not null;comment '收录专栏'" json:"collectionColumn"`
	Summary          string `gorm:"type:varbinary(300);not null;comment '摘要'" json:"summary"`
	Views            int    `gorm:"type:int;default:1;comment '阅读数'" json:"views"`
	Likes            int    `gorm:"type:int;default:1;comment '点赞数'" json:"likes"`
	Favorites        int    `gorm:"type:int;default:1;comment '收藏数'" json:"favorites"`
	Status           string `gorm:"type:varbinary(6);not null;comment '状态 发布、审核未通过、草稿、已删除'" json:"status"`
}
