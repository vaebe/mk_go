package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models"
	"mk/models/user"
	"mk/utils"
)

// GetUserList
//
//	@Summary		获取user用户列表
//	@Description	获取user用户列表
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/getUserList [post]
func GetUserList(ctx *gin.Context) {
	listForm := user.ListForm{}
	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	var users []user.User
	db := global.DB
	if listForm.Email != "" {
		db = db.Where("user_account LIKE ?", "%"+listForm.Email+"%")
	}

	if listForm.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+listForm.NickName+"%")
	}

	res := db.Find(&users)

	// 存在错误
	if res.Error != nil {
		zap.S().Info(res.Error)
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 获取总数
	total := int32(res.RowsAffected)

	// 分页
	res.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&users)

	for i := range users {
		users[i].Password = ""
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     users,
	})
}

// Details
//
//	@Summary		获取用户详情
//	@Description	获取用户详情
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"用户id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/details [get]
func Details(ctx *gin.Context) {
	userId := ctx.Query("id")

	if userId == "" {
		utils.ResponseResultsError(ctx, "用户id不能为空！")
		return
	}

	userInfo := user.User{}
	res := global.DB.Where("id", userId).First(&userInfo)
	userInfo.Password = ""

	if res.Error != nil {
		utils.ResponseResultsError(ctx, "用户不存在！")
		return
	}
	utils.ResponseResultsSuccess(ctx, userInfo)
}

// Edit
//
//	@Summary		编辑用户信息
//	@Description	编辑用户信息
//	@Tags			user用户
//	@Accept			json
//	@Produce		json
//	@Param			param	body		user.EditForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/user/edit [post]
func Edit(ctx *gin.Context) {
	editForm := user.EditForm{}
	if err := ctx.ShouldBind(&editForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	if editForm.ID != userId {
		utils.ResponseResultsError(ctx, "非本人不可修改用户信息！")
		return
	}

	res := global.DB.Where("id = ?", userId).Updates(&user.User{
		NickName:        editForm.NickName,
		Posts:           editForm.Posts,
		Homepage:        editForm.Homepage,
		PersonalProfile: editForm.PersonalProfile,
		Github:          editForm.Github,
		UserAvatar:      editForm.UserAvatar,
		Company:         editForm.Company,
	})

	if res.RowsAffected == 0 {
		utils.ResponseResultsError(ctx, "需要更新的数据不存在！")
		return
	}

	utils.ResponseResultsSuccess(ctx, "更新用户信息成功！")
}
