package commentInfo

import (
	"github.com/gin-gonic/gin"
	"mk/models/commentInfo"
	"mk/models/user"
	"mk/services/commentInfoServices"
	"mk/services/commonServices"
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

	userId, _ := ctx.Get("userId")
	err := commentInfoServices.Save(saveForm, userId.(int32))

	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}
	utils.ResponseResultsSuccess(ctx, "保存成功！")
}

// 根据评论信息数据获取以用户id做key用户信息对象
func getUserInfoMap(infoList []commentInfo.CommentInfo) (map[int32]user.User, error) {
	userIdsMap := make(map[int32]bool)
	for _, v := range infoList {
		userIdsMap[v.UserId] = true

		if v.ReplyUserId != 0 {
			userIdsMap[v.ReplyUserId] = true
		}
	}

	// 获取userIds list 数据
	userIds := make([]int32, 0, len(userIdsMap))
	for id := range userIdsMap {
		userIds = append(userIds, id)
	}

	// 获取以用户 id 做 key 的用户信息对象
	return commonServices.GetUserInfoMapWithIdAskey(userIds)
}

type ItemType struct {
	commentInfo.CommentInfo
	ReplyInfoText string    `json:"replyInfoText"`
	Level         int32     `json:"level"`
	UserInfo      user.User `json:"userInfo"`
	ReplyUserInfo user.User `json:"replyUserInfo"`
}

type TreeItem struct {
	ItemType
	Children []*TreeItem `json:"children"`
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

	infoList, err := commentInfoServices.GetListById(objId)
	// 没有评论直接返回空list，无需执行后续操作
	if err != nil || len(infoList) == 0 {
		utils.ResponseResultsSuccess(ctx, infoList)
		return
	}

	// 获取以用户 id 做 key 的用户信息对象
	userInfoMap, err := getUserInfoMap(infoList)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	// 以id做key的评论信息对象
	commentInfoMap := make(map[int32]commentInfo.CommentInfo)
	for _, v := range infoList {
		commentInfoMap[v.ID] = v
	}

	// list 转森林结构返回
	var tree []*TreeItem
	idMapTreeItem := make(IdMapTreeType)
	for _, item := range infoList {
		var treeItem TreeItem
		treeItem.ID = item.ID
		treeItem.ParentCommentId = item.ParentCommentId
		treeItem.ObjId = item.ObjId
		treeItem.UserId = item.UserId
		treeItem.ReplyUserId = item.ReplyUserId
		treeItem.ReplyInfoId = item.ReplyInfoId
		treeItem.CommentText = item.CommentText
		treeItem.ImgUrl = item.ImgUrl
		treeItem.Type = item.Type
		treeItem.CreatedAt = item.CreatedAt
		treeItem.UserInfo = userInfoMap[item.UserId]
		treeItem.ReplyUserInfo = userInfoMap[item.ReplyUserId]

		// 回复用户id 存在获取引用信息，引用信息不存在则表示已删除
		if item.ReplyUserId != 0 {
			if val, ok := commentInfoMap[item.ReplyInfoId]; !ok {
				treeItem.ReplyInfoText = "该评论已被删除"
			} else {
				treeItem.ReplyInfoText = val.CommentText
			}
		}

		// 根节点收集
		if item.ParentCommentId == 0 {
			treeItem.Level = 1
			tree = append(tree, &treeItem)
		} else {
			// 子节点收集
			treeItem.Level = 2
			idMapTreeItem[item.ParentCommentId].Children = append(idMapTreeItem[item.ParentCommentId].Children, &treeItem)
		}
		// 把节点映射到map表
		idMapTreeItem[item.ID] = &treeItem
	}

	utils.ResponseResultsSuccess(ctx, tree)
}

// Delete
//
//	@Summary		根据id删除评论
//	@Description	根据id删除评论
//	@Tags			commentInfo评论
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"评论id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/commentInfo/delete [delete]
func Delete(ctx *gin.Context) {
	dataId := ctx.Query("id")

	if dataId == "" {
		utils.ResponseResultsError(ctx, "评论信息id不能为空！")
		return
	}

	userId, _ := ctx.Get("userId")
	err := commentInfoServices.Delete(dataId, userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}
