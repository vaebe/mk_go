package articleComment

import (
	"mk/models"
)

// ArticleComment 文章评论
type ArticleComment struct {
	models.BaseModel
	ParentCommentId int32  `gorm:"type:int;unique; not null; comment '上级评论id'" json:"parentCommentId"`
	ArticleId       string `gorm:"type:int;unique; not null; comment '文章id'" json:"articleId"`
	UserId          int32  `gorm:"type:int;unique; not null; comment '用户id'" json:"userId"`
	CommentText     string `gorm:"type:varbinary(500); not null; comment '评论内容'" json:"commentText"`
	ImgUrl          string `gorm:"type:varbinary(200); comment '评论图片'" json:"imgUrl"`
}

// SaveForm 保存评论
type SaveForm struct {
	ParentCommentId int32  `json:"parentCommentId" form:"parentCommentId" binding:"required"`
	ArticleId       int32  `json:"articleId" form:"articleId" binding:"required"`
	UserId          int32  `json:"userId" form:"userId" binding:"required"`
	CommentText     string `json:"commentText" form:"commentText" binding:"required"`
	ImgUrl          string `json:"imgUrl" form:"imgUrl"`
}

// DetailForm 获取文章评论详情
type DetailForm struct {
	ArticleId int32  `json:"articleId" form:"articleId" binding:"required"`
	Status    string `json:"status" form:"status" binding:"required"`
}
