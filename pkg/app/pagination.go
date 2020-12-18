package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shea11012/go_blog/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()

	if page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return 1
	}

	return pageSize
}

func GetPageOffset(page,pageSize int) int {
	result := 0
	if page > 1 {
		result = (page - 1) * pageSize
	}

	return result
}
