package util

import (
	"github.com/rayyu03/gin-blog/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(context *gin.Context)  int {
	result := 0
	page, _ := com.StrTo(context.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}