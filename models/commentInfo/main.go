package commentInfo

import (
	"mk/models"
)

// CommentInfo 评论信息
type CommentInfo struct {
	models.BaseModel
	ObjId           int32  `gorm:"type:int; not null; comment '评论对象id'" json:"ObjId"`
	ParentCommentId int32  `gorm:"type:int; comment '上级评论id'" json:"parentCommentId"`
	UserId          int32  `gorm:"type:int; not null; comment '用户id'" json:"userId"`
	ReplyUserId     int32  `gorm:"type:int; not null; comment '回复用户id'" json:"replyUserId"`
	CommentText     string `gorm:"type:varbinary(500); not null; comment '评论内容'" json:"commentText"`
	ImgUrl          string `gorm:"type:varbinary(200); comment '评论图片'" json:"imgUrl"`
	Type            string `gorm:"type:varbinary(10); comment '评论类型 1 文章 2沸点'" json:"type"`
}

// SaveForm 保存评论
type SaveForm struct {
	ParentCommentId int32  `json:"parentCommentId" form:"parentCommentId"`
	ObjId           int32  `json:"objId" form:"objId" binding:"required"`
	UserId          int32  `json:"userId" form:"userId" binding:"required"`
	ReplyUserId     int32  `json:"replyUserId" form:"replyUserId"`
	CommentText     string `json:"commentText" form:"commentText" binding:"required"`
	ImgUrl          string `json:"imgUrl" form:"imgUrl"`
	Type            string `json:"type" form:"type" example:"1" binding:"required"` // 评论类型 1 文章 2沸点
}
