package commentInfo

import (
	"github.com/gin-gonic/gin"
	"mk/global"
	"mk/models/commentInfo"
	"mk/models/user"
	"mk/utils"
)

// Save
//
//	@Summary		保存评论信息
//	@Description	保存评论信息
//	@Tags			commentInfo评论
//	@Accept			json
//	@Produce		json
//	@Param			param	body		commentInfo.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/commentInfo/save [post]
func Save(ctx *gin.Context) {
	saveForm := commentInfo.SaveForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	saveInfo := commentInfo.CommentInfo{
		ObjId:           saveForm.ObjId,
		ParentCommentId: saveForm.ParentCommentId,
		UserId:          saveForm.UserId,
		ReplyUserId:     saveForm.ReplyUserId,
		CommentText:     saveForm.CommentText,
		ImgUrl:          saveForm.ImgUrl,
		Type:            saveForm.Type,
	}

	res := global.DB.Create(&saveInfo)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}
	utils.ResponseResultsSuccess(ctx, "保存成功！")
}

type ItemType struct {
	ID              int32
	ParentCommentId int32
	ObjId           int32
	UserId          int32
	ReplyUserId     int32
	CommentText     string
	ImgUrl          string
	Type            string
	NickName        string
	UserAvatar      string
	Posts           string
}

type TreeItem struct {
	ItemType
	Children []*TreeItem
}

type IdMapTreeType map[int32]*TreeItem

// GetListById
//
//	@Summary		根据id获取评论列表
//	@Description	根据id获取评论列表
//	@Tags			commentInfo评论
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"对象id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Router			/commentInfo/getListById [get]
func GetListById(ctx *gin.Context) {
	objId := ctx.Query("id")

	if objId == "" {
		utils.ResponseResultsError(ctx, "对象id不能为空！")
		return
	}

	var infoList []commentInfo.CommentInfo
	res := global.DB.Where("obj_id = ?", objId).Find(&infoList)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	var userInfoObj map[int32]user.User

	var tree []*TreeItem
	idMapTreeItem := make(IdMapTreeType)
	for _, item := range infoList {
		var treeItem TreeItem
		treeItem.ID = item.ID
		treeItem.ParentCommentId = item.ParentCommentId
		treeItem.ObjId = item.ObjId
		treeItem.UserId = item.UserId
		treeItem.ReplyUserId = item.ReplyUserId
		treeItem.CommentText = item.CommentText
		treeItem.ImgUrl = item.ImgUrl
		treeItem.Type = item.Type

		// 获取用户信息
		userInfo, ok := userInfoObj[item.UserId]
		if !ok {
			res := global.DB.Where("id = ?", item.UserId).First(&userInfo)

			if res.Error != nil {
				utils.ResponseResultsError(ctx, res.Error.Error())
				return
			}
		}

		// 用户信息
		treeItem.NickName = userInfo.NickName
		treeItem.UserAvatar = userInfo.UserAvatar
		treeItem.Posts = userInfo.Posts

		// 根节点收集
		if item.ParentCommentId == 0 {
			tree = append(tree, &treeItem)
		} else {
			// 子节点收集
			idMapTreeItem[item.ParentCommentId].Children = append(idMapTreeItem[item.ParentCommentId].Children, &treeItem)
		}
		// 把节点映射到map表
		idMapTreeItem[item.ID] = &treeItem
	}

	utils.ResponseResultsSuccess(ctx, tree)
}
