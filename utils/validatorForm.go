package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mk/global"
	"net/http"
	"strings"
)

func removeTopStruct(filedMap map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range filedMap {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ResponseResultsError(c, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"data": nil,
		"msg":  removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}
