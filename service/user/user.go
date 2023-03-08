package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models"
	"mk/utils"
)

//

// GetUserList
// @Summary     获取user用户列表
// @Description  获取user用户列表
// @Tags         user用户
// @Accept       json
// @Produce      json
// @Param 			param body    models.UserListForm  true  "请求对象"
// @Success      200  {object}  utils.ResponseResultInfo
// @Failure      500  {object}  utils.EmptyInfo
// @Router       /user/getUserList [post]
func GetUserList(ctx *gin.Context) {
	userListForm := models.UserListForm{}
	if err := ctx.ShouldBind(&userListForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	var users []models.User
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
