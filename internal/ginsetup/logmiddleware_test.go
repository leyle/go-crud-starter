package ginsetup

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

const (
	path1 = "/api/def/xyz"
	path2 = "/api/def/xyz/123abcdef"
	path3 = "/api/def/xyz/123abcdef/98a"
	path4 = "/api/def/abc"
)

func TestIgnoreBodyPath(t *testing.T) {
	pattern := "/api/def/xyz/\\w+$"

	re := regexp.MustCompile(pattern)
	t.Log(re.String())

	assert.Equal(t, false, re.MatchString(path1))
	assert.Equal(t, true, re.MatchString(path2))
	assert.Equal(t, false, re.MatchString(path3))
	assert.Equal(t, false, re.MatchString(path4))

}

func TestIgnoreBodyPath2(t *testing.T) {
	src := "/api/def/xyz/*"

	pattern := strings.Replace(src, "*", `\w+$`, 1)
	re := regexp.MustCompile(pattern)
	t.Log(re.String())

	assert.Equal(t, false, re.MatchString(path1))
	assert.Equal(t, true, re.MatchString(path2))
	assert.Equal(t, false, re.MatchString(path3))
	assert.Equal(t, false, re.MatchString(path4))
}
