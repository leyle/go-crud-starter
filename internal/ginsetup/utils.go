package ginsetup

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	SortAsc  = 1
	SortDesc = -1
)

var (
	MaxOnePageSize = 200
)

func GetPageSizeSkip(c *gin.Context) (page, size, skip int) {
	p := c.Query("page")
	s := c.Query("size")

	if p != "" {
		page, _ = strconv.Atoi(p)
	} else {
		page = 1
	}

	if s != "" {
		size, _ = strconv.Atoi(s)
	} else {
		size = 10
	}

	if page < 1 {
		page = 1
	}

	if size > MaxOnePageSize {
		size = MaxOnePageSize
	}

	skip = (page - 1) * size

	return page, size, skip
}

func GuessFileContentType(filePrefix512Buf []byte) string {
	ct := http.DetectContentType(filePrefix512Buf)
	return ct
}
