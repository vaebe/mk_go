package articleServices

import (
	"errors"
	"gorm.io/gorm"
	"mk/global"
	"mk/models/article"
	"mk/models/articleAssociatedInfo"
	"mk/services/commonServices"
	"mk/utils"
)

// 保存关联专栏信息
func saveArticlesAssociatedColumns(tx *gorm.DB, saveInfo article.SaveForm) error {
	// 关联信息直接删除
	res := tx.Where("id = ?", saveInfo.ID).Unscoped().Delete(&articleAssociatedInfo.ArticlesAssociatedColumns{})
	if res.Error != nil {
		return res.Error
	}

	if len(saveInfo.CollectionColumn) == 0 {
		return nil
	}

	var articlesAssociatedColumns []articleAssociatedInfo.ArticlesAssociatedColumns
	for _, v := range saveInfo.CollectionColumn {
		articlesAssociatedColumns = append(articlesAssociatedColumns, articleAssociatedInfo.ArticlesAssociatedColumns{
			ColumnId:  v,
			ArticleId: saveInfo.ID,
			UserId:    saveInfo.UserId,
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
	res := tx.Where("id = ?", saveInfo.ID).Unscoped().Delete(&articleAssociatedInfo.ArticlesRelatedTags{})
	if res.Error != nil {
		return res.Error
	}

	if len(saveInfo.Tags) == 0 {
		return nil
	}

	var articlesRelatedTags []articleAssociatedInfo.ArticlesRelatedTags
	for _, v := range saveInfo.Tags {
		articlesRelatedTags = append(articlesRelatedTags, articleAssociatedInfo.ArticlesRelatedTags{
			TagId:     v,
			ArticleId: saveInfo.ID,
			UserId:    saveInfo.UserId,
		})
	}

	res = tx.Create(&articlesRelatedTags)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// CreateAndUpdate 创建或者更新
func CreateAndUpdate(saveForm article.SaveForm, status string, loginUserId int32) (int32, error) {
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
		db := global.DB.Create(&articleInfo)
		if db.Error != nil {
			return 0, db.Error
		}
		return articleInfo.ID, nil
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
		articleInfo.ID = saveForm.ID
		db := tx.Model(&article.Article{}).Where("id = ? AND user_id = ?", saveForm.ID, loginUserId).Updates(&articleInfo)
		if db.Error != nil {
			// 发生错误回滚事务
			tx.Rollback()
			return 0, db.Error
		}

		// 保存文章关联专栏信息
		err := saveArticlesAssociatedColumns(tx, saveForm)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		// 保存关联标签信息
		err = saveArticlesRelatedTags(tx, saveForm)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		// 提交事务
		err = tx.Commit().Error
		if err != nil {
			return 0, err
		}

		return articleInfo.ID, err
	}
}

// FormatArticleListReturnData 格式化文章列表返回数据
func FormatArticleListReturnData(articles []article.Article) ([]article.ArticleInfo, error) {
	var resultList []article.ArticleInfo
	// 文章ids
	articleIds := make([]int32, 0, len(articles))

	// 用户ids map 文章用户id去重
	userIdsMap := make(map[int32]bool)

	for i, v := range articles {
		// 将文章内容重置为 ""
		articles[i].Content = ""

		// 文章用户id去重
		userIdsMap[v.UserId] = true

		// 保存文章ids
		articleIds = append(articleIds, v.ID)
	}

	// 查用户信息
	userIds := make([]int32, 0, len(userIdsMap))
	for id := range userIdsMap {
		userIds = append(userIds, id)
	}

	userInfoMap, err := commonServices.GetUserInfoMapWithIdAskey(userIds)
	if err != nil {
		return resultList, err
	}

	// 查文章标签信息
	articleTagsMap, err := commonServices.GetArticleTagMapWithIdAskey(articleIds)
	if err != nil {
		return resultList, err
	}

	for _, v := range articles {
		// 标题为空赋值为无标题
		if v.Title == "" {
			v.Title = "无标题"
		}

		resultList = append(resultList, article.ArticleInfo{
			ArticleDetails:   v,
			UserInfo:         userInfoMap[v.UserId],
			Tags:             articleTagsMap[v.ID],
			NumberOfComments: 0,
		})
	}

	return resultList, nil
}

// GetArticleList 获取文章列表
func GetArticleList(listForm article.ListForm) ([]article.ArticleInfo, int32, error) {
	var articles []article.Article
	db := global.DB

	if listForm.UserId != 0 {
		db = db.Where("user_id", listForm.UserId)
	}

	if listForm.Title != "" {
		db = db.Where("title LIKE ?", "%"+listForm.Title+"%")
	}

	if listForm.Classify != "" {
		db = db.Where("classify LIKE ?", "%"+listForm.Classify+"%")
	}

	if listForm.Status != "" {
		db = db.Where("status = ?", listForm.Status)
	}

	db.Find(&articles)

	var resultList []article.ArticleInfo

	// 存在错误
	if db.Error != nil {
		return resultList, 0, db.Error
	}

	// 获取总数
	total := int32(db.RowsAffected)

	// 分页
	db = db.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&articles)

	if db.Error != nil {
		return resultList, 0, db.Error
	}

	resultList, err := FormatArticleListReturnData(articles)
	return resultList, total, err
}

// Details 获取文章详情
func Details(id string) (article.Details, error) {
	details := article.Article{}
	db := global.DB.Where("id = ?", id).First(&details)

	var result article.Details

	if db.Error != nil {
		return result, db.Error
	}

	// 查看记录加1
	if details.Status == "4" {
		db := global.DB.Model(&article.Article{}).Where("id = ?", id).UpdateColumn("views", gorm.Expr("views + ?", 1))
		if db.Error != nil {
			return result, db.Error
		}
	}

	result = article.Details{
		ID:         details.ID,
		UserId:     details.UserId,
		Title:      details.Title,
		Content:    details.Content,
		Classify:   details.Classify,
		CoverImg:   details.CoverImg,
		Summary:    details.Summary,
		Views:      details.Views + 1,
		Likes:      details.Likes,
		Favorites:  details.Favorites,
		ShowNumber: details.ShowNumber,
		Status:     details.Status,
		UpdatedAt:  details.UpdatedAt,
		CreatedAt:  details.CreatedAt,
	}

	// 获取文章专栏
	var columns []articleAssociatedInfo.ArticlesAssociatedColumns
	db = global.DB.Select("column_id").Where("article_id", details.ID).Find(&columns)
	if db.Error != nil {
		return result, db.Error
	}

	for _, v := range columns {
		result.CollectionColumn = append(result.CollectionColumn, v.ColumnId)
	}

	// 数据为空返回空数组
	if len(columns) == 0 {
		result.CollectionColumn = make([]int32, 0)
	}

	// 获取文章标签
	var tags []articleAssociatedInfo.ArticlesRelatedTags
	db = global.DB.Select("tag_id").Where("article_id", details.ID).Find(&tags)
	if db.Error != nil {
		return result, db.Error
	}

	if len(tags) == 0 {
		result.Tags = make([]string, 0)
	}

	for _, v := range tags {
		result.Tags = append(result.Tags, v.TagId)
	}

	return result, nil
}

// Delete 文章删除
func Delete(id string, userId int32) error {
	db := global.DB.Where("id = ? AND user_id = ?", id, userId).Delete(&article.Article{})

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在")
	}

	return nil
}

// Review 文章审核
func Review(reviewForm article.ReviewForm) error {
	db := global.DB.Where("id = ?", reviewForm.ID).Updates(article.Article{
		Status: reviewForm.Status,
	})

	return db.Error
}
