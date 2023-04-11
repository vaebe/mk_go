package common

import (
	"mk/global"
	"mk/models/articleAssociatedInfo"
	"mk/models/user"
)

// GetUserInfoMapWithIdAskey 获取以 id 做 key 的用户信息对象
func GetUserInfoMapWithIdAskey(userIds []int32) (map[int32]user.User, error) {
	// 查询用户信息数据
	var userList []user.User
	res := global.DB.Where("id in (?)", userIds).Find(&userList)

	if res.Error != nil {
		return nil, res.Error
	}

	userInfoMap := make(map[int32]user.User)
	for _, v := range userList {
		v.Password = ""
		userInfoMap[v.ID] = v
	}

	return userInfoMap, nil
}

// GetArticleTagMapWithIdAskey 获取以 id 做 key 的文章标签对象
func GetArticleTagMapWithIdAskey(articleIds []int32) (map[int32][]string, error) {
	// 查询文章标签信息
	var articlesRelatedTag []articleAssociatedInfo.ArticlesRelatedTags
	res := global.DB.Where("article_id in (?)", articleIds).Find(&articlesRelatedTag)

	if res.Error != nil {
		return nil, res.Error
	}

	articlesRelatedTagsMap := make(map[int32][]string)
	for _, v := range articlesRelatedTag {
		articlesRelatedTagsMap[v.ArticleId] = append(articlesRelatedTagsMap[v.ArticleId], v.TagId)
	}

	return articlesRelatedTagsMap, nil
}
