package ginsetup

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	okMsg = "OK"
)

type APIReturnForm struct {
	// ReqId string      `json:"reqId"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func generateReturnData(code int, msg string, data interface{}) *APIReturnForm {
	info := &APIReturnForm{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return info
}

func ReturnJson(c *gin.Context, statusCode, code int, msg string, data interface{}) {
	text := http.StatusText(statusCode)
	if text == "" {
		panic(errors.New("runtime error: invalid http status code"))
	}

	ret := generateReturnData(code, msg, data)
	// reqId := c.Request.Context().Value(ReqIdContextName).(string)
	// ret.ReqId = reqId

	if statusCode != http.StatusOK {
		c.AbortWithStatusJSON(statusCode, ret)
	} else {
		c.JSON(statusCode, ret)
	}
}

func ReturnOKJson(c *gin.Context, data interface{}) {
	ReturnJson(c, http.StatusOK, 200, okMsg, data)
}

func Return400Json(c *gin.Context, code int, msg string) {
	ReturnJson(c, http.StatusBadRequest, code, msg, "")
}

func Return401Json(c *gin.Context, code int, msg string) {
	ReturnJson(c, http.StatusUnauthorized, code, msg, "")
}

func Return403Json(c *gin.Context, code int, msg string) {
	ReturnJson(c, http.StatusForbidden, code, msg, "")
}

func Return500Json(c *gin.Context, code int, msg string) {
	ReturnJson(c, http.StatusInternalServerError, code, msg, "")
}
