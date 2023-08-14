package pingapp

import (
	"github.com/gin-gonic/gin"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/leyle/go-crud-starter/internal/ginsetup"
	"github.com/leyle/go-crud-starter/utils"
)

func PingHandler(c *gin.Context) {
	ctx := configandcontext.GetAPIContextFromGinContext(c)
	t0 := utils.PrintFuncStartLog(ctx)
	defer utils.PrintFuncEndLog(ctx, t0)

	// return server info
	method := c.Request.Method
	userAgent := c.Request.Header.Get("User-Agent")

	pong := &PongInfo{
		HttpMethod: method,
		UserAgent:  userAgent,
		ReqId:      ctx.ReqId(),
		Version:    configandcontext.Version,
		CommitId:   configandcontext.CommitId,
	}

	ginsetup.ReturnOKJson(c, pong)
	return
}
