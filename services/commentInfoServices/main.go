package commentInfoServices

import (
	"errors"
	"mk/global"
	"mk/models/commentInfo"
)

// Save 保存评论信息
func Save(saveForm commentInfo.SaveForm, loginUserId int32) error {
	saveInfo := commentInfo.CommentInfo{
		ObjId:           saveForm.ObjId,
		ParentCommentId: saveForm.ParentCommentId,
		ReplyInfoId:     saveForm.ReplyInfoId,
		UserId:          loginUserId,
		ReplyUserId:     saveForm.ReplyUserId,
		CommentText:     saveForm.CommentText,
		ImgUrl:          saveForm.ImgUrl,
		Type:            saveForm.Type,
	}

	db := global.DB.Create(&saveInfo)
	if db.Error != nil {

		return db.Error
	}
	return nil
}

// GetListById 根据id获取评论列表
func GetListById(objId string) ([]commentInfo.CommentInfo, error) {
	var infoList []commentInfo.CommentInfo
	db := global.DB.Where("obj_id = ?", objId).Find(&infoList)
	return infoList, db.Error
}

// Delete 根据 id 删除评论信息
func Delete(id string, loginUserId int32) error {
	db := global.DB.Where("id = ? AND user_id = ?", id, loginUserId).Delete(&commentInfo.CommentInfo{})

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("需要删除的数据不存在！")
	}

	return nil
}
