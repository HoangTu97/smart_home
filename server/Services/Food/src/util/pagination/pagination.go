package pagination

import (
	"Food/util/converter"
	"github.com/gin-gonic/gin"
)

// GetPage get page parameters
func GetPage(c *gin.Context) (int, int) {
	result := 0
	page := converter.MustInt(c.Query("page"))
	size := converter.MustInt(c.Query("size"))
	if page > 0 {
		result = (page - 1) * size
	}

	return result, size
}