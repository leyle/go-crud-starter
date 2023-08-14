package smsapp

const (
	SMSModuleAddr    = "addr"
	SMSModulePasswd  = "passwd"
	SMSModuleKeyPair = "keypair"
)

type SMSRequestForm struct {
	Module     string
	Receiver   string
	CodeLength int
	redisKey   string // Module +  md5(Receiver)
	code       string
	msg        string
}

type SMSVerifyForm struct {
	Module   string
	Receiver string
	redisKey string // md5(Receiver)
	Code     string
}
