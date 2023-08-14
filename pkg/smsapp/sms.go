package smsapp

import (
	"fmt"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/leyle/go-crud-starter/utils"
)

// send sms

func SendSMS(ctx *configandcontext.APIContext, info *SMSRequestForm) (string, error) {
	t0 := utils.PrintFuncStartLog(ctx)
	defer utils.PrintFuncEndLog(ctx, t0)

	userKey := utils.CalcMD5(info.Receiver)
	info.redisKey = ctx.Cfg.Redis.GenerateRedisKey(info.Module, userKey)

	if !IsSupportedCountry(ctx, info.Receiver) {
		return "", ErrNotSupportedCountry
	}

	if IsRequestTooFast(ctx, info.redisKey) {
		return "", ErrRequestTooFast
	}

	info.code = GenerateRandomCode(info.CodeLength)
	dbCode := getCodeFromRedis(ctx, info.redisKey)
	if dbCode != "" {
		info.code = dbCode
	}

	msg := fmt.Sprintf(ctx.Cfg.SMS.MsgFormat, info.code)
	info.msg = msg

	err := sendTwilioSMS(ctx, info)
	if err != nil {
		return "", ErrSendSMSFailed
	}

	// save code into redis
	err = SaveCodeIntoRedis(ctx, info)

	if ctx.Cfg.SMS.Debug {
		ctx.Logger.Debug().Msgf("receiver[%s], code[%s]", info.Receiver, info.code)
	}

	return info.code, nil
}

// sms code verification

func VerifySMS(ctx *configandcontext.APIContext, info *SMSVerifyForm) error {
	t0 := utils.PrintFuncStartLog(ctx)
	defer utils.PrintFuncEndLog(ctx, t0)

	userKey := utils.CalcMD5(info.Receiver)
	info.redisKey = ctx.Cfg.Redis.GenerateRedisKey(info.Module, userKey)
	err := CheckCodeFromRedis(ctx, info)

	// log, todo

	return err
}
