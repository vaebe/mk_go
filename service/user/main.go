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
	userListForm := user.ListForm{}
	if err := ctx.ShouldBind(&userListForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	var users []user.User
	res := global.DB.Where("user_account LIKE ? AND nick_name LIKE ?", "%"+userListForm.Email+"%", "%"+userListForm.NickName+"%").Find(&users)

	// 存在错误
	if res.Error != nil {
		zap.S().Info(res.Error)
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 获取总数
	total := int32(res.RowsAffected)

	// 分页
	res.Scopes(utils.Paginate(userListForm.PageNo, userListForm.PageSize)).Find(&users)

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: userListForm.PageSize,
		PageNo:   userListForm.PageNo,
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
//	@Router			/user/getUserDetails [get]
func Details(ctx *gin.Context) {
	articleId := ctx.Query("id")

	if articleId == "" {
		utils.ResponseResultsError(ctx, "用户id不能为空！")
		return
	}

	var users []user.User
	res := global.DB.Where("id", articleId).First(&users)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, "用户不存在！")
		return
	}
	utils.ResponseResultsSuccess(ctx, users[0])
}