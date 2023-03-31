package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models"
	"mk/models/article"
	"mk/utils"
)

// SaveArticle 保存文章
func SaveArticle(ctx *gin.Context, status string) {
	saveDraftForm := article.SaveDraftForm{}

	if err := ctx.ShouldBind(&saveDraftForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	articleInfo := article.Article{
		UserId:           saveDraftForm.UserId,
		Classify:         saveDraftForm.Classify,
		CollectionColumn: saveDraftForm.CollectionColumn,
		Content:          saveDraftForm.Content,
		CoverImg:         saveDraftForm.CoverImg,
		Tags:             saveDraftForm.Tags,
		Title:            saveDraftForm.Title,
		Summary:          saveDraftForm.Summary,
		Status:           status,
	}

	// id 不存在新增
	if saveDraftForm.ID == 0 {
		global.DB.Create(&articleInfo)
		utils.ResponseResultsSuccess(ctx, map[string]any{"id": articleInfo.ID})
	} else {
		userId, _ := ctx.Get("userId")
		// 增加 userId 防止非对应用户更改信息
		res := global.DB.Model(&article.Article{}).Where("id = ? AND user_id = ?", saveDraftForm.ID, userId).Updates(&articleInfo)
		if res.Error != nil {
			utils.ResponseResultsError(ctx, res.Error.Error())
			return
		}

		utils.ResponseResultsSuccess(ctx, map[string]any{"id": saveDraftForm.ID})
	}
}

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
	SaveArticle(ctx, "2")
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
	res := global.DB.Where("title LIKE ? AND tags LIKE ? AND classify LIKE ? AND status LIKE ?",
		"%"+listForm.Title+"%", "%"+listForm.Tag+"%", "%"+listForm.Classify+"%", "%"+listForm.Status+"%").Find(&articles)

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

	for i := range articles {
		articles[i].Content = ""
	}

	// todo 增加用户名称 、comments  评论数
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
	res := global.DB.Where("user_id = ?", userId)

	if listForm.Status == "" {
		res.Not("status = ?", "1").Find(&articles)
	} else {
		res.Where("status = ?", listForm.Status).Find(&articles)
	}

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
	res := global.DB.Where("id = ? AND user_id = ?", articleId, userId).First(&articles)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, "文章不存在！")
		return
	}
	utils.ResponseResultsSuccess(ctx, articles[0])
}

// Review
//
//	@Summary		文章审核
//	@Description	文章审核
//	@Tags			article文章
//	@Accept			json
//	@Produce		json
//	@Param			param	body		article.ReviewForm	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/article/review [post]
func Review(ctx *gin.Context) {
	reviewForm := article.ReviewForm{}
	if err := ctx.ShouldBind(&reviewForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	authorityId, _ := ctx.Get("authorityId")
	if authorityId == 1 {
		utils.ResponseResultsError(ctx, "您没有审核权限！")
		return
	}

	res := global.DB.Where("id = ?", reviewForm.ID).Updates(article.Article{
		Status: reviewForm.Status,
	})

	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
	} else {
		utils.ResponseResultsSuccess(ctx, "审核成功！")
	}
}
