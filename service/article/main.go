package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models"
	"mk/models/article"
	"mk/utils"
)

// Save
//
//	@Summary		保存文章
//	@Description	保存文章
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.SaveForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/save [post]
func Save(ctx *gin.Context) {
	saveForm := article.SaveForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		zap.S().Info("文章保存信息:", &saveForm)
		utils.HandleValidatorError(ctx, err)
		return
	}

	global.DB.Create(&article.Article{
		UserId:           saveForm.UserId,
		Classify:         saveForm.Classify,
		CollectionColumn: saveForm.CollectionColumn,
		Content:          saveForm.Content,
		CoverImg:         saveForm.CoverImg,
		Tags:             saveForm.Tags,
		Title:            saveForm.Title,
		Summary:          saveForm.Summary,
		Status:           "2", // 文章保存的信息必须进行审核
	})

	utils.ResponseResultsSuccess(ctx, "保存成功！")
}

// GetArticleList
//
//	@Summary		获取文章列表
//	@Description	获取文章列表
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.AllListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Router			/article/getArticleList [post]
func GetArticleList(ctx *gin.Context) {
	listForm := article.AllListForm{}

	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	var articles []article.Article
	res := global.DB.Where("title LIKE ? AND tags LIKE ? AND classify LIKE ?",
		"%"+listForm.Title+"%", "%"+listForm.Tag+"%", "%"+listForm.Classify+"%").Find(&articles)

	// 存在错误
	if res.Error != nil {
		zap.S().Info(res.Error)
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 获取总数
	total := int32(res.RowsAffected)

	// 分页
	res.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&articles)

	// todo 全部列表不需要返回文章详情
	// 增加用户名称 、comments  评论数
	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     articles,
	})
}

// GetUserArticleList
//
//	@Summary		获取用户文章列表
//	@Description	获取用户文章列表
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.UserArticleListForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/getUserArticleList [post]
func GetUserArticleList(ctx *gin.Context) {
	listForm := article.UserArticleListForm{}
	if err := ctx.ShouldBind(&listForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	var articles []article.Article
	res := global.DB.Where("user_id", userId).Find(&articles)

	// 存在错误
	if res.Error != nil {
		zap.S().Info(res.Error)
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	// 获取总数
	total := int32(res.RowsAffected)

	// 分页
	res.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&articles)

	utils.ResponseResultsSuccess(ctx, &models.PagingData{
		PageSize: listForm.PageSize,
		PageNo:   listForm.PageNo,
		Total:    total,
		List:     articles,
	})
}

// Details
//
//	@Summary		获取文章详情
//	@Description	获取文章详情
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"文章id"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/getArticleDetails [get]
func Details(ctx *gin.Context) {
	articleId := ctx.Query("id")

	if articleId == "" {
		utils.ResponseResultsError(ctx, "文章id不能为空！")
		return
	}

	userId, _ := ctx.Get("userId")
	var articles []article.Article
	res := global.DB.Where("id = ?, user_id = ?", articleId, userId).First(&articles)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, "文章不存在！")
		return
	}
	utils.ResponseResultsSuccess(ctx, articles[0])
}
