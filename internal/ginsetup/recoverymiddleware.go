package ginsetup

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strings"
)

const (
	runtimeErrAPICode = 5006 // same as cache app error
)

func RecoveryMiddleware(f func(*gin.Context, *StopForm)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			recoveryVal := recover()
			if recoveryVal == nil {
				c.Next()
				return
			}

			e, ok := recoveryVal.(*StopForm)
			if ok {
				// user behavior, this is an expected panic
				f(c, e)
			} else {
				// other panic situations, e.g nil pointer, etc
				// check if the msg has `runtime`
				eMsg := fmt.Sprintf("%v", recoveryVal)
				if strings.Contains(eMsg, "runtime") {
					// runtime error, means internal error
					debug.PrintStack()
				}
				sf := &StopForm{
					Err:            errors.New(eMsg),
					HTTPStatusCode: http.StatusInternalServerError,
					APICode:        runtimeErrAPICode,
					Data:           nil,
				}
				f(c, sf)
			}
		}()
		c.Next()
	}
}

func DefaultStopExecHandler(c *gin.Context, sf *StopForm) {
	ReturnJson(c, sf.HTTPStatusCode, sf.APICode, sf.Err.Error(), sf.Data)
}
