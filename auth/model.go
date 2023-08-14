package auth

import (
	"fmt"
	"github.com/leyle/go-crud-starter/utils"
	"strings"
)

const (
	aesKeyPrefix = "EMALICBDC"
)

type KYCToken struct {
	Msg       string `json:"msg"`
	Signature string `json:"signature"`
}

type KeyPairUser struct {
	Valid        bool
	Address      string // ethereum address, lower case
	EIP55Address string
	PublicKeyHex string
	RawMsg       string
	Signature    string
}

func (kpu *KeyPairUser) AesKey() string {
	addr := strings.ToLower(kpu.Address[len(kpu.Address)-20:])
	msg := fmt.Sprintf("%s%s", aesKeyPrefix, addr)
	hash := utils.CalcMD5(msg)

	aesKey := strings.ToLower(hash)
	return aesKey
}

type JWTTokenResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data *JWTTokenUser `json:"data"`
}

type JWTTokenUser struct {
	// only extract partial fields
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Expiry   int64  `json:"exp"`
	RawToken string `json:"-"`
}
