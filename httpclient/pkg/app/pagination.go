package app

import (
	"httpclient/global"
	"httpclient/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int64 {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return int64(page)
}

func GetPageSize(c *gin.Context) int64 {
	pageSize := int64(convert.StrTo(c.Query("page_size")).MustInt())
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int64) int64 {
	result := int64(0)
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
