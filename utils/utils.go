package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/leyle/go-crud-starter/configandcontext"
	"path"
	"runtime"
	"time"
)

// usage:
// func SampleMethod(ctx *APIContext) {
// 	t0 := PrintFuncStartLog(ctx)
// 	defer PrintFuncEndLog(ctx, t0)
// }

// called by function as the first line code

func PrintFuncStartLog(ctx *configandcontext.APIContext) time.Time {
	// if it gets function name failed, simply omit it
	t := time.Now()
	funcName := "noFuncName"
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		details := runtime.FuncForPC(pc)
		if details != nil {
			funcName = path.Base(details.Name())
		}
	}
	ctx.Logger.Debug().Msgf("start executing %s()", funcName)
	return t
}

// defer call this method in function

func PrintFuncEndLog(ctx *configandcontext.APIContext, startT time.Time) {
	funcName := "noFuncName"
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		details := runtime.FuncForPC(pc)
		if details != nil {
			funcName = path.Base(details.Name())
		}
	}

	ctx.Logger.Debug().Msgf("exit executing %s(), elapsed[%v]", funcName, time.Since(startT))
}

func CalcMD5(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])
}
