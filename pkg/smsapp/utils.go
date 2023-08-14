package smsapp

import (
	"context"
	"crypto/rand"
	"github.com/leyle/go-crud-starter/configandcontext"
	"strings"
	"time"
)

const (
	minCodeLength     = 4
	maxCodeLength     = 8
	defaultCodeLength = 6
)

func GenerateRandomCode(length int) string {
	if length < minCodeLength {
		length = defaultCodeLength
	}
	if length > maxCodeLength {
		length = defaultCodeLength
	}

	base := "0123456789"
	baseLen := len(base)

	data := make([]byte, length)
	rand.Read(data)
	for i := 0; i < length; i++ {
		data[i] = base[int(data[i])%baseLen]
	}
	return string(data)
}

func IsSupportedCountry(ctx *configandcontext.APIContext, phone string) bool {
	for _, prefix := range ctx.Cfg.SMS.Supported {
		if strings.HasPrefix(phone, prefix) {
			return true
		}
	}
	return false
}

func IsRequestTooFast(ctx *configandcontext.APIContext, key string) bool {
	// https://redis.io/commands/ttl/
	// ttl check
	t, err := ctx.Redis.TTL(context.Background(), key).Result()
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("get redis key TTL failed")
		return true
	}

	remainT := int64(t / time.Second)
	if remainT < 0 {
		// key doesn't exist or doesn't have expiry time
		return false
	}

	baseT := ctx.Cfg.SMS.Rate
	maxT := ctx.Cfg.SMS.ExpiryIn // in case to send the same code within a period time
	if maxT-remainT < baseT {
		// too fast
		ctx.Logger.Debug().Int64("maxT", maxT).Int64("remainT", remainT).Int64("baseT", baseT).Int64("sub", maxT-remainT).Send()
		return true
	}

	return false
}

func SaveCodeIntoRedis(ctx *configandcontext.APIContext, info *SMSRequestForm) error {
	key := info.redisKey
	val := info.code
	expiryIn := time.Duration(ctx.Cfg.SMS.ExpiryIn) * time.Second

	err := ctx.Redis.Set(context.Background(), key, val, expiryIn).Err()
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("save verification code into redis failed")
		return ErrInternalError
	}
	return nil
}

func CheckCodeFromRedis(ctx *configandcontext.APIContext, info *SMSVerifyForm) error {
	key := info.redisKey
	code, err := ctx.Redis.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			// code has expired
			return err
		}
		return err
	}

	if code == info.Code {
		return nil
	}

	ctx.Logger.Error().Msgf("invalid user input sms code")
	return ErrInvalidCode
}

func getCodeFromRedis(ctx *configandcontext.APIContext, key string) string {
	code, err := ctx.Redis.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return ""
		}
		// other error, also return empty string
		return ""
	}

	return code
}
