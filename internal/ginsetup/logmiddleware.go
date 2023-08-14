package ginsetup

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leyle/crud-objectid/pkg/objectid"
	"github.com/rs/zerolog"
	"io"
	"time"
)

const (
	ReqIdHeaderName        = "X-REQ-ID"
	ReqIdContextName       = "reqId"
	ClientReqIdContextName = "clientReqId"
)

var PrintHeaders = false

var (
	ignoreReadReqBodyPath      []string
	ignoreReadResponseBodyPath []string
)

const (
	pathTypeReq = iota
	pathTypeResponse
)

func AddIgnoreReadReqBodyPath(paths ...string) {
	ignoreReadReqBodyPath = append(ignoreReadReqBodyPath, paths...)
}

func AddIgnoreReadResponseBodyPath(paths ...string) {
	ignoreReadResponseBodyPath = append(ignoreReadResponseBodyPath, paths...)
}

func isIgnoreReadBodyPath(pathType int, reqPath string) bool {
	paths := ignoreReadReqBodyPath
	if pathType == pathTypeResponse {
		paths = ignoreReadResponseBodyPath
	}

	for _, path := range paths {
		if reqPath == path {
			return true
		}
	}
	return false
}

// rewrite Write()
type respWriter struct {
	gin.ResponseWriter
	cache *bytes.Buffer
}

// it will increase memory usage
func (r *respWriter) Write(b []byte) (int, error) {
	r.cache.Write(b)
	return r.ResponseWriter.Write(b)
}

func reqJson(c *gin.Context) []byte {
	path := c.Request.RequestURI
	method := c.Request.Method
	cType := c.Request.Header.Get("Content-Type")
	clientIp := c.ClientIP()

	data := map[string]string{
		"path":        path,
		"method":      method,
		"contentType": cType,
		"ip":          clientIp,
	}
	bData, _ := json.Marshal(data)
	return bData
}

func reqBody(c *gin.Context) []byte {
	var err error
	var body []byte
	if c.Request.ContentLength > 0 && !isIgnoreReadBodyPath(pathTypeReq, c.Request.URL.Path) {
		body, err = io.ReadAll(c.Request.Body)
		if err == nil {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
	}
	return body
}

func GinLogMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startT := time.Now()

		curReqId := objectid.GetObjectId()

		l := logger.With().Str(ReqIdContextName, curReqId).Logger()

		l.Trace().Msg("execute GinMiddleware() instance")

		// first get reqId from client
		clientReqId := c.Request.Header.Get(ReqIdHeaderName)
		if clientReqId != "" {
			l.Info().Str(ClientReqIdContextName, clientReqId).Send()
		}

		// save reqId into current context
		reqIdCtx := context.WithValue(c.Request.Context(), ReqIdContextName, curReqId)
		c.Request = c.Request.WithContext(reqIdCtx)

		zl := l.WithContext(c.Request.Context())
		c.Request = c.Request.WithContext(zl)

		// print req
		body := reqBody(c)
		reqInfo := reqJson(c)
		event := l.Debug().Str("type", "REQUEST").RawJSON("req", reqInfo).RawJSON("body", body)
		if PrintHeaders {
			headers, _ := json.Marshal(c.Request.Header)
			event.RawJSON("headers", headers)
		}
		event.Send()

		// write req id to response headers
		c.Writer.Header().Set(ReqIdHeaderName, curReqId)

		c.Writer = &respWriter{
			ResponseWriter: c.Writer,
			cache:          bytes.NewBufferString(""),
		}

		c.Next()

		// write response
		// not all apis need to do this, exclude special case, e.g. file download
		if c.Writer.Header().Get("Content-Type") == "" {
			c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		}
		statusCode := c.Writer.Status()
		rEvent := l.Info().Str("type", "RESPONSE").Int("statusCode", statusCode)
		if clientReqId != "" {
			rEvent.Str(ClientReqIdContextName, clientReqId)
		}
		rw, ok := c.Writer.(*respWriter)
		if !ok {
			// silently passed
		} else {
			if rw.cache.Len() > 0 && !isIgnoreReadBodyPath(pathTypeResponse, c.Request.URL.Path) {
				rEvent.RawJSON("body", rw.cache.Bytes())
			}
		}
		latency := time.Since(startT)
		rEvent.Str("latency", latency.String()).Msg("")
	}
}
