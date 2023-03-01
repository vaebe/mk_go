package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	data := map[string]interface{}{
		"code": "0",
		"msg":  "登陆",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	c.AsciiJSON(http.StatusOK, data)
}
