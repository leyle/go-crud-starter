package redisclient

import (
	"context"
	"github.com/leyle/crud-objectid/pkg/objectid"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	defaultAcquireTimeout = 30 * time.Second
	defaultLockKeyTimout  = 30 * time.Second
	defaultRetryDuration  = 10 * time.Millisecond
)

const (
	moduleName = "LOCK"
)

func AcquireLock(ctx *configandcontext.APIContext, resource string, acquireTimeout, lockTimeout time.Duration) (string, bool) {
	if acquireTimeout < 0 {
		acquireTimeout = defaultAcquireTimeout
	}
	if lockTimeout < 0 {
		lockTimeout = defaultLockKeyTimout
	}

	val := objectid.GetObjectId()
	endTime := time.Now().Add(acquireTimeout)

	redisKey := ctx.Cfg.Redis.GenerateRedisKey(moduleName, resource)

	for time.Now().UnixMilli() < endTime.UnixMilli() {
		ok, err := ctx.Redis.SetNX(context.Background(), redisKey, val, lockTimeout).Result()
		if err != nil {
			ctx.Logger.Error().Err(err).Str("resource", resource).Msg("try to set redis lock(SetNX) failed")
			return "", false
		}
		if ok {
			ctx.Logger.Info().Str("resource", resource).Msgf("set redis lock succeeded, lock val is:[%s]", val)
			return val, true
		} else {
			// retry with 10 millisecond
			time.Sleep(defaultRetryDuration)
			continue
		}
	}

	ctx.Logger.Error().Str("resource", resource).Msgf("with [%s] time period, try to get lock failed", acquireTimeout.String())
	return "", false
}

func ReleaseLock(ctx *configandcontext.APIContext, resource, val string) bool {
	redisKey := ctx.Cfg.Redis.GenerateRedisKey(moduleName, resource)

	v, err := ctx.Redis.Get(context.Background(), redisKey).Result()
	if err != nil && err != redis.Nil {
		ctx.Logger.Error().Err(err).Str("resource", resource).Str("val", val).Msg("release redis lock failed")
		return false
	}

	if err == redis.Nil {
		ctx.Logger.Debug().Str("resource", resource).Str("val", val).Msg("lock key has expired, release lock succeed")
		return true
	}

	if v == val {
		ctx.Redis.Del(context.Background(), redisKey)
		ctx.Logger.Debug().Str("resource", resource).Str("val", val).Msg("delete lock key, release lock succeed")
		return true
	} else {
		ctx.Logger.Warn().Str("resource", resource).Str("val", val).Msg("when try to release lock, but the lock has locked by others, we think this situation is ok ")
		return true
	}
}
