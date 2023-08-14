package ginsetup

import (
	"github.com/gin-gonic/gin"
	"github.com/leyle/go-crud-starter/auth"
	"github.com/leyle/go-crud-starter/configandcontext"
)

var noAuthPaths []string

const (
	JWTTokenHeaderKey = "Authorization"
	KYCTokenHeaderKey = "X-KYC-TOKEN"
)

func AddNoAuthPaths(paths ...string) []string {
	noAuthPaths = append(noAuthPaths, paths...)
	return noAuthPaths
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := configandcontext.GetAPIContextFromGinContext(c)
		ctx.Logger.Trace().Msg("execute AuthMiddleware()")

		needAuth := true
		uri := c.Request.URL.Path
		for _, path := range noAuthPaths {
			if uri == path {
				ctx.Logger.Info().Str("uri", uri).Msg("current path won't be executed in Auth() method")
				needAuth = false
				break
			}
		}

		if needAuth {
			Auth(ctx, c)
		}

		c.Next()
	}
}

func Auth(ctx *configandcontext.APIContext, c *gin.Context) {
	ctx.Logger.Trace().Msg("execute Auth()")
	kycToken := c.Request.Header.Get(KYCTokenHeaderKey)
	jwtToken := c.Request.Header.Get(JWTTokenHeaderKey)

	if kycToken == "" && jwtToken == "" {
		ctx.Logger.Error().Msg("no token header has been passed in request headers")
		Return401Json(c, 401, "no token header in request headers")
		return
	}

	if kycToken != "" {
		ctx.UserType = configandcontext.UserTypeKYC
		if !isValidKYCToken(ctx, kycToken) {
			Return401Json(c, 401, "invalid token")
			return
		}
	} else if jwtToken != "" {
		ctx.UserType = configandcontext.UserTypeJWT
		if !isValidJWTToken(ctx, jwtToken) {
			Return401Json(c, 401, "invalid token")
			return
		}
	} else {
		ctx.Logger.Error().Msg("no token header has been passed in request headers")
		Return401Json(c, 401, "no token header in request headers")
		return
	}

	// everything is ok, token is valid
	configandcontext.SetAPIContextWithGinContext(ctx, c)

	c.Next()
}

func isValidKYCToken(ctx *configandcontext.APIContext, token string) bool {
	result, err := auth.CheckSignatureToken(ctx, token)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("invalid kyc token")
		return false
	}

	aesKey := result.AesKey()
	userInfo := &configandcontext.KYCUserInfo{
		Address:      result.Address,
		EIP55Address: result.EIP55Address,
		PublicKey:    result.PublicKeyHex,
		AesKey:       aesKey,
	}
	ctx.KYCUser = userInfo

	return true
}

func isValidJWTToken(ctx *configandcontext.APIContext, token string) bool {
	result, err := auth.CheckJWTToken(ctx, token)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("invalid jwt token")
		return false
	}

	userInfo := &configandcontext.JWTUserInfo{
		UserId:   result.UserId,
		Username: result.UserName,
		Expiry:   result.Expiry,
	}
	ctx.JWTUser = userInfo

	return true
}
