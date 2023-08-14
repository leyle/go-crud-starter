package ginsetup

import (
	"errors"
	"net/http"
)

type StopForm struct {
	Err            error
	HTTPStatusCode int
	APICode        int
	Data           interface{}
}

func StopExec(err error, httpStatusCode, apiCode int, data interface{}) {
	if err == nil {
		return
	}

	sf := &StopForm{
		Err:            err,
		HTTPStatusCode: httpStatusCode,
		APICode:        apiCode,
		Data:           data,
	}
	panic(sf)
}

func StopExec400(err error, code int, data interface{}) {
	if err == nil {
		return
	}
	StopExec(err, http.StatusBadRequest, code, data)
}

func StopExec401(err error, code int, data interface{}) {
	if err == nil {
		return
	}
	StopExec(err, http.StatusUnauthorized, code, data)
}

func StopExec403(err error, code int, data interface{}) {
	if err == nil {
		return
	}
	StopExec(err, http.StatusForbidden, code, data)
}

func StopExec404(err error, code int, data interface{}) {
	if err == nil {
		return
	}
	StopExec(err, http.StatusNotFound, code, data)
}

func StopExec500(err error, code int, data interface{}) {
	if err == nil {
		return
	}
	StopExec(err, http.StatusInternalServerError, code, data)
}

func StopExec501() {
	err := errors.New("not implemented")
	StopExec(err, http.StatusInternalServerError, 501, "")
}
