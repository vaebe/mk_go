package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mk/global"
	"mk/models"
	"net/http"
)

func Login(c *gin.Context) {
	var users []models.UserTest
	res := global.DB.Find(&users)

	// 存在错误
	if res.Error != nil {
		return
	}

	zap.S().Info("数据", users)

	data := map[string]interface{}{
		"code": "0",
		"msg":  "登陆",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}
