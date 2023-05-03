package articleColumn

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/models"
	"mk/models/articleColumn"
	"mk/services/articleColumnServices"
	"mk/utils"
)

// Save
//
//	@Summary		保存专栏
//	@Description	保存专栏
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleColumn.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/save [post]
func Save(ctx *gin.Context) {
	saveForm := articleColumn.SaveForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		zap.S().Info("专栏保存信息:", &saveForm)
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	infoId, err := articleColumnServices.SaveOrUpdate(saveForm, userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, infoId)
}

// Delete
//
//	@Summary		根据id删除专栏
//	@Description	根据id删除专栏
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"专栏id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/delete [delete]
func Delete(ctx *gin.Context) {
	columnId := ctx.Query("id")

	if columnId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	userId, _ := ctx.Get("userId")
	err := articleColumnServices.Delete(columnId, userId.(int32))
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "删除成功！")
}

// Details
//
//	@Summary		获取专栏详情
//	@Description	获取专栏详情
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"专栏id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/details [get]
func Details(ctx *gin.Context) {
	columnId := ctx.Query("id")

	if columnId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	columnInfo, err := articleColumnServices.Details(columnId)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, columnInfo)
}

// List
//
//	@Summary		获取专栏列表
//	@Description	获取专栏列表
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleColumn.ListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/getList [post]
func List(ctx *gin.Context) {
	listForm := articleColumn.ListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	articleColumnList, total, err := articleColumnServices.List(listForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     articleColumnList,
	})
}

// Review
//
//	@Summary		文章专栏审核
//	@Description	文章专栏审核
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleColumn.ReviewForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/review [post]
func Review(ctx *gin.Context) {
	reviewForm := articleColumn.ReviewForm{}
	if err := ctx.ShouldBind(&reviewForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	authorityId, _ := ctx.Get("authorityId")
	if authorityId == 1 {
		utils.ResponseResultsError(ctx, "您没有审核权限！")
		return
	}

	err := articleColumnServices.Review(reviewForm)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, "审核成功！")
}
