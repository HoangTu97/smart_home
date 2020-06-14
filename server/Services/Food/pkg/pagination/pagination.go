package pagination

import (
	"Food/pkg/converter"

	"github.com/gin-gonic/gin"
)

// Pageable godoc
type Pageable interface {
	GetPageNumber() int
	GetPageSize() int
}

type pageable struct {
	pageNumber int
	pageSize   int
}

// GetPage get page parameters
func GetPage(c *gin.Context) Pageable {
	result := 0
	page := converter.MustInt(c.Query("page"))
	size := converter.MustInt(c.Query("size"))
	if page > 0 {
		result = (page - 1) * size
	}

	return &pageable{pageNumber: result, pageSize: size}
}

func (p *pageable) GetPageNumber() int {
	return p.pageNumber
}

func (p *pageable) GetPageSize() int {
	return p.pageSize
}
