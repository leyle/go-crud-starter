package main

import (
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/leyle/go-crud-starter/internal/ginsetup"
	"github.com/leyle/go-crud-starter/pkg/pingapp"
	"os"
)

func httpServer(ctx *configandcontext.APIContext) {
	var err error
	logger := ctx.Cfg.Log.GetLogger()

	e := ginsetup.SetupGin(logger)
	e.Use(configandcontext.NewAPIContextMiddleware(ctx))
	e.Use(ginsetup.AuthMiddleware())

	// if needed to print request headers
	if ctx.Cfg.Server.Debug {
		ginsetup.PrintHeaders = true
	}

	// set no auth path
	if len(ctx.Cfg.Auth.NoAuthPaths) > 0 {
		ginsetup.AddNoAuthPaths(ctx.Cfg.Auth.NoAuthPaths...)
	}

	// set no log req body uris
	if len(ctx.Cfg.Log.IgnoreReqBody) > 0 {
		ginsetup.AddIgnoreReadReqBodyPath(ctx.Cfg.Log.IgnoreReqBody...)
	}

	// set no log response body uris
	if len(ctx.Cfg.Log.IgnoreResponseBody) > 0 {
		ginsetup.AddIgnoreReadResponseBodyPath(ctx.Cfg.Log.IgnoreResponseBody...)
	}

	baseRouter := e.Group("/api/v1")

	// ping router
	pingapp.PingRouter(baseRouter.Group(""))

	addr := ctx.Cfg.Server.GetServerListeningAddr()
	err = e.Run(addr)
	if err != nil {
		logger.Error().Err(err).Msg("start http server failed")
		os.Exit(1)
	}
}
