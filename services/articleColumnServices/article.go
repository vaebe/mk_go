package articleColumnServices

import (
	"errors"
	"mk/global"
	"mk/models/article"
	"mk/models/articleAssociatedInfo"
	"mk/models/articleColumn"
	"mk/services/articleServices"
)

// GetAssociatedArticlesList 获取专栏关联文章列表
func GetAssociatedArticlesList(id string) ([]article.ArticleInfo, error) {
	var linkedData []articleAssociatedInfo.ArticlesAssociatedColumns
	db := global.DB.Select("article_id").Where("column_id = ?", id).Find(&linkedData)

	var resultList []article.ArticleInfo

	if db.Error != nil {
		return resultList, db.Error
	}

	// 获取文章id数组
	articleIdsMap := make(map[int32]bool)
	for _, v := range linkedData {
		articleIdsMap[v.ArticleId] = true
	}

	articleIds := make([]int32, 0, len(articleIdsMap))
	for id := range articleIdsMap {
		articleIds = append(articleIds, id)
	}

	// 根据文章id 查询返回
	var articles []article.Article
	db = global.DB.Where("id IN ?", articleIds).Find(&articles)

	if db.Error != nil {
		return resultList, db.Error
	}

	resultList, err := articleServices.FormatArticleListReturnData(articles)
	return resultList, err
}

// ListArticlesThatCanBeIncluded 获取可以被收录的文章列表
func ListArticlesThatCanBeIncluded(id string) ([]article.Article, error) {
	// 获取关联次数大于2 或者是当前专栏关联的文章 = 不可以被再次关联
	var linkedData []articleAssociatedInfo.ArticlesAssociatedColumns
	db := global.DB.Where("article_id IN (?) OR id != ?",
		global.DB.Table("articles_associated_columns").Select("article_id").Group("article_id").Having("COUNT(*) > 2"), id).
		Find(&linkedData)

	var articles []article.Article

	if db.Error != nil {
		return articles, db.Error
	}

	// 获取文章id数组
	articleIdsMap := make(map[int32]bool)
	for _, v := range linkedData {
		articleIdsMap[v.ArticleId] = true
	}

	articleIds := make([]int32, 0, len(articleIdsMap))
	for id := range articleIdsMap {
		articleIds = append(articleIds, id)
	}

	db = global.DB.Where("NOT id IN (?) AND status = ?", articleIds, "4").Find(&articles)
	return articles, db.Error
}

// DeleteAssociatedArticle 删除关联文章
func DeleteAssociatedArticle(delForm articleAssociatedInfo.ArticlesAssociatedColumnsForm, loginUserId int32) error {
	db := global.DB.Where("column_id = ? AND article_id = ? AND user_id = ?",
		delForm.ColumnId, delForm.ArticleId, loginUserId).Unscoped().Delete(&articleAssociatedInfo.ArticlesAssociatedColumns{})

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在！")
	}
	return nil
}

// AddAssociatedArticle
//
//	@Summary		添加关联文章
//	@Description	添加关联文章
//	@Tags				articleColumn专栏
//	@Accept			json
//	@Produce		json
//	@Param			param	body		articleAssociatedInfo.ArticlesAssociatedColumnsForm	true	"请求对象"
//	@Success		200	{object}	utils.ResponseResultInfo
//	@Failure		500	{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/articleColumn/addAssociatedArticle [post]
func AddAssociatedArticle(saveForm articleAssociatedInfo.ArticlesAssociatedColumnsForm, loginUserId int32) error {
	// 校验是否是当前用户的专栏
	columnInfo := articleColumn.ArticleColumn{}
	db := global.DB.Where("id = ?", saveForm.ColumnId).First(&columnInfo)
	if db.Error != nil {
		return db.Error
	}

	if columnInfo.UserId != loginUserId {
		return errors.New("非本用户的专栏无法关联！")
	}

	if columnInfo.Status != "3" {
		return errors.New("不是发布状态的专栏无法关联！")
	}

	// 校验是否是当前用户的文章
	articleInfo := article.Article{}
	db = global.DB.Where("id = ?", saveForm.ArticleId).First(&articleInfo)
	if db.Error != nil {
		return db.Error
	}

	if articleInfo.UserId != loginUserId {
		return errors.New("非本用户的文章无法关联！")
	}

	if articleInfo.Status != "4" {
		return errors.New("不是发布状态的文章不能关联！")
	}

	// 验证通过创建关联信息
	db = global.DB.Create(&articleAssociatedInfo.ArticlesAssociatedColumns{
		UserId:    articleInfo.UserId,
		ArticleId: saveForm.ArticleId,
		ColumnId:  saveForm.ColumnId,
	})

	return db.Error
}
