package configandcontext

import (
	"fmt"
	"github.com/leyle/crud-log/pkg/crudlog"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"strings"
)

const (
	logFormatDefault = "line"
	logFormatJson    = "json"
	logFormatLine    = "line"
)

const (
	DefaultHTTPRequestTimeout = 15 // seconds
	DefaultRedisAcquireTime   = 30 // seconds
	DefaultRedisLockTime      = 30 // seconds
)

var (
	Version  string
	CommitId string
	Branch   string
)

var (
	MgoMinPoolSize uint64 = 20
	MgoMaxPoolSize uint64 = 200
)

type APIConfig struct {
	Server        *ServerConf        `yaml:"server"`
	Auth          *AuthConf          `yaml:"auth"`
	Log           *LogConf           `yaml:"log"`
	Redis         *RedisConf         `yaml:"redis"`
	Mongodb       *MongodbConf       `yaml:"mongodb"`
	SMS           *SMSConf           `yaml:"sms"`
	SmartContract *SmartContractConf `yaml:"smartContract"`
}

type ServerConf struct {
	Debug bool   `yaml:"debug"`
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
}

func (s *ServerConf) GetServerListeningAddr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type AuthConf struct {
	NoAuthPaths []string           `yaml:"noAuthPaths"`
	Signature   *AuthSignatureConf `yaml:"signature"`
	JWT         *AuthJWTConf       `yaml:"jwt"`
}

type AuthSignatureConf struct {
	ExpiryIn int64 `yaml:"expiryIn"`
}

type AuthJWTConf struct {
	Host      string `yaml:"host"`
	VerifyAPI string `yaml:"verifyAPI"`
}

func (jwt *AuthJWTConf) GetVerifyURL() string {
	return fmt.Sprintf("%s%s", jwt.Host, jwt.VerifyAPI)
}

type LogConf struct {
	Level              string   `yaml:"level"`
	Format             string   `yaml:"format"`
	IgnoreReqBody      []string `yaml:"ignoreReqBody"`
	IgnoreResponseBody []string `yaml:"ignoreResponseBody"`
}

func (l *LogConf) GetLevel() zerolog.Level {
	level, err := zerolog.ParseLevel(l.Level)
	if err != nil {
		level = zerolog.DebugLevel
		fmt.Printf("parse log level string failed, %v, using default log level: DEBUG\n", err)
	}
	return level
}

func (l *LogConf) GetLogger() zerolog.Logger {
	switch l.Format {
	case logFormatJson:
		return crudlog.NewJsonLogger(l.GetLevel())
	case logFormatLine:
		return crudlog.NewConsoleLogger(l.GetLevel())
	default:
		fmt.Printf("invalid log format[%s], only [%s] or [%s] is valid, now we are using default log format[%s]\n", l.Format, logFormatJson, logFormatLine, logFormatDefault)
		return crudlog.NewConsoleLogger(l.GetLevel())
	}
}

type RedisConf struct {
	HostPort    string `yaml:"hostPort"`
	Password    string `yaml:"password"`
	DbNum       int    `yaml:"dbNum"`
	Service     string `yaml:"service"`
	AcquireTime int64  `yaml:"acquireTime"`
	LockTime    int64  `yaml:"lockTime"`
}

func (r *RedisConf) GenerateRedisKey(moduleName, userKey string) string {
	// SERVICE:MODULE:USER_KEY
	// module name shouldn't have ":"
	moduleName = strings.ReplaceAll(moduleName, ":", "")
	moduleName = strings.ToUpper(moduleName)
	userKey = strings.ToUpper(userKey)
	return fmt.Sprintf("%s:%s:%s", r.Service, moduleName, userKey)
}

type MongodbConf struct {
	Replica      bool     `yaml:"replica"`
	ReplicaSet   string   `yaml:"replicaSet"`
	HostPorts    []string `yaml:"hostPorts"`
	Username     string   `yaml:"username"`
	Password     string   `yaml:"password"`
	Database     string   `yaml:"database"`
	ConnOption   string   `yaml:"connOption"`
	WriteTimeout int      `yaml:"writeTimeout"`
	ReadTimeout  int      `yaml:"readTimeout"`
	TLS          struct {
		Enabled bool   `yaml:"enabled"`
		PEM     string `yaml:"pem"`
	} `yaml:"tls"`
}

type SMSConf struct {
	Debug       bool     `yaml:"debug"`
	Rate        int64    `yaml:"rate"`
	ExpiryIn    int64    `yaml:"expiryIn"`
	CodeLength  int      `yaml:"codeLength"`
	Supported   []string `yaml:"supported"`
	MsgFormat   string   `yaml:"msgFormat"`
	TwilioURL   string   `yaml:"twilioURL"`
	TwilioSID   string   `yaml:"twilioSID"`
	TwilioToken string   `yaml:"twilioToken"`
}

type SmartContractConf struct {
	Host string                `yaml:"host"`
	API  *SmartContractAPIConf `yaml:"api"`
}

type SmartContractAPIConf struct {
	CreateProduct string `yaml:"createProduct"`
}

func (sc *SmartContractConf) GetAPI(uri string) string {
	return fmt.Sprintf("%s%s", sc.Host, uri)
}

func LoadConfig(path string, v interface{}) error {
	var err error
	if err = CheckPathExist(path, 4); err != nil {
		return err
	}

	viper.SetConfigFile(path)
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(v)
	if err != nil {
		return err
	}

	return nil
}
