package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/leyle/dbandpubsub/mongodb"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func getMongodbClient(cfg *configandcontext.MongodbConf) *mongodb.DataSource {
	op := &mongodb.MgoOption{
		HostPorts:    cfg.HostPorts,
		Username:     cfg.Username,
		Password:     cfg.Password,
		Database:     cfg.Database,
		ConnOption:   cfg.ConnOption,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	moreOptions := &options.ClientOptions{
		MaxPoolSize: &configandcontext.MgoMaxPoolSize,
		MinPoolSize: &configandcontext.MgoMinPoolSize,
	}
	// override default replicaset name
	if cfg.Replica && cfg.ReplicaSet != "" {
		moreOptions.ReplicaSet = &cfg.ReplicaSet
	}

	if cfg.TLS.Enabled {
		// check file is valid
		err := configandcontext.CheckPathExist(cfg.TLS.PEM, 4)
		if err != nil {
			fmt.Println("invalid mongodb tls pem file: ", cfg.TLS.PEM, err.Error())
			os.Exit(1)
		}
		tlsCfg, err := setMongodbTLSCfg(cfg.TLS.PEM)
		if err != nil {
			fmt.Println("load pem file failed: ", cfg.TLS.PEM, err.Error())
			os.Exit(1)
		}

		moreOptions.SetTLSConfig(tlsCfg)
	}

	op.MoreOptions = moreOptions

	ds := mongodb.NewDataSource(op)

	return ds
}

func setMongodbTLSCfg(pem string) (*tls.Config, error) {
	// Load the PEM file
	pemBytes, err := os.ReadFile(pem)
	if err != nil {
		return nil, err
	}

	// Create a CertPool to hold the PEM file's contents
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(pemBytes); !ok {
		return nil, errors.New("load pem file failed")
	}

	// Create a new TLS Config instance with the CertPool
	tlsConfig := &tls.Config{
		RootCAs: roots,
	}

	return tlsConfig, nil
}

func getRedisClient(cfg *configandcontext.RedisConf) *redis.Client {
	opt := &redis.Options{
		Addr:     cfg.HostPort,
		Password: cfg.Password,
		DB:       cfg.DbNum,
	}
	rdb := redis.NewClient(opt)

	// test redis connection
	key := cfg.GenerateRedisKey("main", "test-redis")
	val := "test-redis-connection"

	ctx := context.Background()
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}

	dbVal, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if dbVal != val {
		fmt.Printf("val[%s] != dbVal[%s], fatal error occurred\n", val, dbVal)
		os.Exit(1)
	}
	err = rdb.Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}

	return rdb
}

func insureMongodbIndex(ds *mongodb.DataSource) error {
	return nil
}
