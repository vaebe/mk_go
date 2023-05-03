package articleColumnServices

import (
	"errors"
	"mk/global"
	"mk/models/articleColumn"
	"mk/utils"
)

// SaveOrUpdate 保存或更新专栏
func SaveOrUpdate(saveForm articleColumn.SaveForm, loginUserId int32) (int32, error) {
	// 新增编辑信息均需要审核
	savaInfo := articleColumn.ArticleColumn{
		Name:         saveForm.Name,
		Introduction: saveForm.Introduction,
		CoverImg:     saveForm.CoverImg,
		Status:       "1",
	}

	// 专栏默认封面
	if savaInfo.CoverImg == "" {
		savaInfo.CoverImg = "https://cdn.qiniu.vaebe.top/mk/default/default_article_column.jpg"
	}

	// id 不存在新增
	if saveForm.ID == 0 {
		if loginUserId == 0 {
			return 0, errors.New("新增专栏用户 id 不能为空！")
		}

		savaInfo.UserId = loginUserId

		db := global.DB.Create(&savaInfo)
		if db.Error != nil {
			return 0, db.Error
		}
		return savaInfo.ID, nil
	} else {
		db := global.DB.Model(&articleColumn.ArticleColumn{}).Where("id = ? AND user_id = ?", saveForm.ID, loginUserId).Updates(&savaInfo)
		if db.Error != nil {
			return 0, db.Error
		}
		return saveForm.ID, nil
	}
}

// Delete 根据id删除专栏
func Delete(id string, loginUserId int32) error {
	db := global.DB.Where("id = ? AND user_id = ?", id, loginUserId).Delete(&articleColumn.ArticleColumn{})

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在！")
	}

	return nil
}

// Details 根据 id 获取专栏详情
func Details(id string) (articleColumn.ArticleColumn, error) {
	columnInfo := articleColumn.ArticleColumn{}
	db := global.DB.Model(&articleColumn.ArticleColumn{}).Where("id = ?", id).First(&columnInfo)

	if db.Error != nil {
		return columnInfo, db.Error
	}

	if db.RowsAffected == 0 {
		return columnInfo, errors.New("专栏不存在！")
	}

	return columnInfo, nil
}

// List 获取专栏列表
func List(listForm articleColumn.ListForm) ([]articleColumn.ArticleColumn, int32, error) {
	var articleColumnList []articleColumn.ArticleColumn

	db := global.DB
	if listForm.UserId != 0 {
		db = db.Where("user_id = ? ", listForm.UserId)
	}

	if listForm.Name != "" {
		db = db.Where("name LIKE ?", "%"+listForm.Name+"%")
	}

	if listForm.Status != "" {
		db = db.Where("status = ?", listForm.Status)
	}

	db.Find(&articleColumnList)

	// 存在错误
	if db.Error != nil {
		return articleColumnList, 0, db.Error
	}

	// 获取总数
	total := int32(db.RowsAffected)

	// 分页
	db.Scopes(utils.Paginate(listForm.PageNo, listForm.PageSize)).Find(&articleColumnList)

	return articleColumnList, total, nil
}

// Review 文章专栏审核
func Review(reviewForm articleColumn.ReviewForm) error {
	db := global.DB.Where("id = ?", reviewForm.ID).Updates(articleColumn.ArticleColumn{
		Status: reviewForm.Status,
	})

	if db.Error != nil {
		return db.Error
	}
	return nil
}
