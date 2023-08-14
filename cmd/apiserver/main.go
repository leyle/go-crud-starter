package main

import (
	"flag"
	"fmt"
	"github.com/leyle/crud-log/pkg/crudlog"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

var (
	Version  string
	CommitId string
	Branch   string
)

func main() {
	configandcontext.Version = Version
	configandcontext.CommitId = CommitId
	configandcontext.Branch = Branch
	printVersion()

	initialLogger := crudlog.NewConsoleLogger(zerolog.TraceLevel)

	var cFile string
	var port int
	flag.StringVar(&cFile, "c", "", "-c /path/to/config.yaml")
	flag.IntVar(&port, "p", 0, "-p PORT")

	flag.Parse()

	// load config
	if cFile == "" {
		initialLogger.Error().Msg("no config file set in the cli args")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var cfg *configandcontext.APIConfig
	err := configandcontext.LoadConfig(cFile, &cfg)
	if err != nil {
		initialLogger.Error().Err(err).Msg("parse config file failed")
		os.Exit(1)
	}

	initialLogger.Debug().Msgf("parse config[%s] succeed", cFile)

	if port > 0 {
		cfg.Server.Port = port
		initialLogger.Debug().Msgf("user changed http server listening port to: %d", port)
	}

	apiCtx := &configandcontext.APIContext{
		Cfg: cfg,
	}

	// get mongodb client
	// if error happened, panic it
	initialLogger.Debug().Msg("trying to connect to mongodb")
	ds := getMongodbClient(cfg.Mongodb)
	defer ds.Close()
	apiCtx.Ds = ds
	initialLogger.Debug().Msg("connect to mongodb succeed")

	// insure mongodb index
	err = insureMongodbIndex(ds)
	if err != nil {
		initialLogger.Error().Err(err).Msg("create mongodb collection index key failed")
		os.Exit(1)
	}

	// get redis client, if error occurred, crash the hole program
	initialLogger.Debug().Msg("trying to connect to redis")
	rdb := getRedisClient(cfg.Redis)
	apiCtx.Redis = rdb
	defer rdb.Close()
	initialLogger.Debug().Msg("connect to redis succeed")

	// start http server
	go httpServer(apiCtx)

	// try to catch some error info
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	// do some cleaning things

	os.Exit(0)
}

func printVersion() {
	fmt.Printf("version: %s, git hash: %s\n, git branch: %s\n", configandcontext.Version, configandcontext.CommitId, configandcontext.Branch)
}
