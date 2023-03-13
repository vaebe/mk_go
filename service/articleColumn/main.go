package articleColumn

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models/articleColumn"
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

	global.DB.Create(&articleColumn.ArticleColumn{
		Name:         saveForm.Name,
		Introduction: saveForm.Introduction,
		CoverImg:     saveForm.CoverImg,
		Status:       "1", // 新建专栏需要审核
	})

	utils.ResponseResultsSuccess(ctx, "保存成功！")
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
	enumsId := ctx.Query("id")

	if enumsId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	res := global.DB.Where("id = ?", enumsId).Delete(&articleColumn.ArticleColumn{})

	total := res.RowsAffected
	if total == 0 {
		utils.ResponseResultsError(ctx, "需要删除的数据不存在！")
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
	enumsId := ctx.Query("id")

	if enumsId == "" {
		utils.ResponseResultsError(ctx, "专栏id不能为空！")
		return
	}

	columnInfo := articleColumn.ArticleColumn{}
	res := global.DB.Model(&articleColumn.ArticleColumn{}).Where("id = ?", enumsId).First(&columnInfo)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, columnInfo)
}

// AllArticleColumnList
//
//	@Summary		获取全部专栏
//	@Description	获取全部专栏
//	@Tags			articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/getAllArticleColumnList [get]
func AllArticleColumnList(ctx *gin.Context) {
	var articleColumnList []articleColumn.ArticleColumn
	res := global.DB.Model(&articleColumn.ArticleColumn{}).Find(&articleColumnList)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 按状态分类的对象
	columns := make(map[string][]articleColumn.ArticleColumn)
	for _, v := range articleColumnList {
		columns[v.Status] = append(columns[v.Status], v)
	}

	utils.ResponseResultsSuccess(ctx, columns)
}
