package auth

import (
	"encoding/json"
	"errors"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/leyle/go-crud-starter/internal/httpclient"
	"net/http"
	"strings"
)

func CheckJWTToken(ctx *configandcontext.APIContext, token string) (*JWTTokenUser, error) {
	// token format: Bearer + space + jwt token
	prefix := "Bearer"
	if !strings.HasPrefix(token, prefix) {
		ctx.Logger.Error().Str("token", token).Msg("invalid jwt token format")
		return nil, errors.New("invalid jwt token format")
	}

	tt := strings.Split(token, " ")
	if len(tt) != 2 {
		ctx.Logger.Error().Str("token", token).Msg("invalid jwt token format")
		return nil, errors.New("invalid jwt token format")
	}

	jwtToken := strings.TrimSpace(tt[1])

	// call sso api to check token
	type ReqBody struct {
		Token string `json:"token"`
	}

	reqBody := &ReqBody{Token: jwtToken}
	reqData, err := json.Marshal(&reqBody)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("marshal token request body to bytes failed")
		return nil, err
	}
	ctx.Logger.Debug().Str("reqBody", string(reqData)).Send()

	url := ctx.Cfg.Auth.JWT.GetVerifyURL()
	client := httpclient.NewClientRequest(http.MethodPost, url)
	client.SetBody(reqData)

	resp := client.Do(ctx)

	if resp.Code != http.StatusOK {
		ctx.Logger.Error().Msg("verify token failed")
		return nil, errors.New(string(resp.Body))
	}

	// decode response body
	var jwtResp *JWTTokenResponse
	err = json.Unmarshal(resp.Body, &jwtResp)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("decode jwt verify response failed")
		return nil, err
	}

	ctx.Logger.Debug().
		Str("userId", jwtResp.Data.UserId).
		Str("username", jwtResp.Data.UserName).
		Msg("jwt token is valid")

	return jwtResp.Data, nil
}
