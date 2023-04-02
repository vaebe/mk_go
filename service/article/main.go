package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mk/global"
	"mk/models"
	"mk/models/article"
	"mk/utils"
)

// 保存关联专栏信息
func saveArticlesAssociatedColumns(tx *gorm.DB, saveInfo article.SaveForm) error {
	res := tx.Where("id = ?", saveInfo.ID).Delete(&article.ArticlesAssociatedColumns{})
	if res.Error != nil {
		return res.Error
	}

	if len(saveInfo.CollectionColumn) == 0 {
		return nil
	}

	var articlesAssociatedColumns []article.ArticlesAssociatedColumns
	for _, v := range saveInfo.CollectionColumn {
		articlesAssociatedColumns = append(articlesAssociatedColumns, article.ArticlesAssociatedColumns{
			ColumnId:  v,
			ArticleId: saveInfo.ID,
		})
	}

	res = tx.Create(&articlesAssociatedColumns)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// 保存文章关联标签信息
func saveArticlesRelatedTags(tx *gorm.DB, saveInfo article.SaveForm) error {
	res := tx.Where("id = ?", saveInfo.ID).Delete(&article.ArticlesRelatedTags{})
	if res.Error != nil {
		return res.Error
	}

	if len(saveInfo.Tags) == 0 {
		return nil
	}

	var articlesRelatedTags []article.ArticlesRelatedTags
	for _, v := range saveInfo.Tags {
		articlesRelatedTags = append(articlesRelatedTags, article.ArticlesRelatedTags{
			TagId:     v,
			ArticleId: saveInfo.ID,
		})
	}

	res = tx.Create(&articlesRelatedTags)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// SaveArticle 保存文章
func SaveArticle(ctx *gin.Context, status string) {
	saveForm := article.SaveForm{}

	if err := ctx.ShouldBind(&saveForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	articleInfo := article.Article{
		UserId:   saveForm.UserId,
		Classify: saveForm.Classify,
		Content:  saveForm.Content,
		CoverImg: saveForm.CoverImg,
		Title:    saveForm.Title,
		Summary:  saveForm.Summary,
		Status:   status,
	}

	// id 不存在新增
	if saveForm.ID == 0 {
		res := global.DB.Create(&articleInfo)
		if res.Error != nil {
			utils.ResponseResultsError(ctx, res.Error.Error())
			return
		}
		utils.ResponseResultsSuccess(ctx, map[string]any{"id": articleInfo.ID})
	} else {
		// 创建事务
		tx := global.DB.Begin()

		// 发生 panic 回滚事务
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// 保存文章信息
		userId, _ := ctx.Get("userId")
		articleInfo.ID = saveForm.ID
		res := tx.Model(&article.Article{}).Where("id = ? AND user_id = ?", saveForm.ID, userId).Updates(&articleInfo)
		if res.Error != nil {
			// 发生错误回滚事务
			tx.Rollback()
			utils.ResponseResultsError(ctx, res.Error.Error())
			return
		}

		// 保存文章关联专栏信息
		err := saveArticlesAssociatedColumns(tx, saveForm)
		if err != nil {
			// 发生错误回滚事务
			tx.Rollback()
			utils.ResponseResultsError(ctx, err.Error())
			return
		}
		// 保存关联标签信息
		err = saveArticlesRelatedTags(tx, saveForm)
		if err != nil {
			// 发生错误回滚事务
			tx.Rollback()
			utils.ResponseResultsError(ctx, err.Error())
			return
		}

		// 提交事务
		err = tx.Commit().Error
		if err != nil {
			utils.ResponseResultsError(ctx, err.Error())
			return
		}

		utils.ResponseResultsSuccess(ctx, map[string]any{"id": articleInfo.ID})
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
	res := global.DB.Where("title LIKE ? AND classify LIKE ? AND status LIKE ?",
		"%"+listForm.Title+"%", "%"+listForm.Classify+"%", "%"+listForm.Status+"%").Find(&articles)

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
	details := article.Article{}
	res := global.DB.Where("id = ? AND user_id = ?", articleId, userId).First(&details)

	if res.Error != nil {
		utils.ResponseResultsError(ctx, "文章不存在！")
		return
	}

	result := article.Details{
		ID:         details.ID,
		UserId:     details.UserId,
		Title:      details.Title,
		Content:    details.Content,
		Classify:   details.Classify,
		CoverImg:   details.CoverImg,
		Summary:    details.Summary,
		Views:      details.Views,
		Likes:      details.Likes,
		Favorites:  details.Favorites,
		ShowNumber: details.ShowNumber,
		Status:     details.Status,
	}

	var columns []article.ArticlesAssociatedColumns
	res = global.DB.Select("column_id").Where("article_id", details.ID).Find(&columns)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	for _, v := range columns {
		result.CollectionColumn = append(result.CollectionColumn, v.ColumnId)
	}

	// 数据为空返回空数组
	if len(columns) == 0 {
		result.CollectionColumn = make([]int32, 0)
	}

	var tags []article.ArticlesRelatedTags
	global.DB.Select("tag_id").Where("article_id", details.ID).Find(&tags)
	if res.Error != nil {
		utils.ResponseResultsError(ctx, res.Error.Error())
		return
	}

	if len(tags) == 0 {
		result.Tags = make([]string, 0)
	}

	for _, v := range tags {
		result.Tags = append(result.Tags, v.TagId)
	}

	utils.ResponseResultsSuccess(ctx, result)
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
