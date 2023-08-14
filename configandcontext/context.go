package configandcontext

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/leyle/dbandpubsub/mongodb"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

const (
	APICtxKey = "apiCtx"
)

type UserType string

const (
	UserTypeKYC UserType = "KYC"
	UserTypeJWT UserType = "JWT"
)

type APIContext struct {
	Cfg      *APIConfig
	C        *gin.Context
	Ds       *mongodb.DataSource
	Redis    *redis.Client
	Logger   *zerolog.Logger
	UserType UserType
	KYCUser  *KYCUserInfo
	JWTUser  *JWTUserInfo
}

type KYCUserInfo struct {
	Address      string // lower case
	EIP55Address string
	PublicKey    string
	AesKey       string
}

type JWTUserInfo struct {
	UserId   string
	Username string
	Expiry   int64
}

func (apiCtx *APIContext) New(c *gin.Context) *APIContext {
	logger := zerolog.Ctx(c.Request.Context())

	logger.Trace().Msg("create new api context")

	newCtx := &APIContext{
		Cfg:    apiCtx.Cfg,
		C:      c,
		Ds:     apiCtx.Ds,
		Redis:  apiCtx.Redis,
		Logger: logger,
	}

	return newCtx
}

func NewAPIContextMiddleware(apiCtx *APIContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		newCtx := apiCtx.New(c)
		newCtx.Logger.Trace().Msg("create new api context middleware for gin")

		SetAPIContextWithGinContext(newCtx, c)

		c.Next()
	}
}

func SetAPIContextWithGinContext(apiCtx *APIContext, c *gin.Context) {
	withCtx := context.WithValue(c.Request.Context(), APICtxKey, apiCtx)
	c.Request = c.Request.WithContext(withCtx)
}

func GetAPIContextFromGinContext(c *gin.Context) *APIContext {
	ctx, ok := c.Request.Context().Value(APICtxKey).(*APIContext)
	if !ok {
		panic("runtime error, get api context from from request context failed")
	}
	return ctx
}

func (apiCtx *APIContext) ReqId() string {
	if apiCtx.C == nil {
		return ""
	}

	reqId := apiCtx.C.Request.Context().Value("reqId").(string)
	return reqId
}
