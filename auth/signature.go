package auth

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/leyle/go-crud-starter/configandcontext"
	"strconv"
	"strings"
	"time"
)

// signature verification

func CheckSignatureToken(ctx *configandcontext.APIContext, token string) (*KeyPairUser, error) {
	// token format: hex json string
	// {
	//    "msg": "ios|1686811276",
	//    "signature": "8992ef577466b7465458cfa913f953504a895c50063fa41aa8e8b4a87c6fc55d5462b815bde49c1308687f5b0f5fec9373877054960915bbdae62a73da3b176d00"
	// }

	// 1. hex to bytes
	ctx.Logger.Debug().Msg("try to valid signature")
	result, err := hex.DecodeString(token)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("decode hex token failed, maybe invalid format")
		return nil, err
	}

	// 2. bytes to json
	var kycToken KYCToken
	err = json.Unmarshal(result, &kycToken)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("unmarshal kyc token to struct failed")
		return nil, err
	}

	// check msg format and created time
	msg := strings.Split(kycToken.Msg, "|")
	if len(msg) != 2 {
		ctx.Logger.Error().Str("msg", kycToken.Msg).Msg("invalid kyc token msg format")
		return nil, errors.New("invalid msg format")
	}

	// check platform, todo

	// check timestamp
	t, err := strconv.ParseInt(msg[1], 10, 64)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("parse timestamp to int64 failed")
		return nil, err
	}

	curT := time.Now().Unix()

	if curT-t > ctx.Cfg.Auth.Signature.ExpiryIn {
		ctx.Logger.Error().Msg("signature has expired")
		return nil, errors.New("signature has expired")
	}

	// check signature, get ethereum address
	kpu, err := checkSignature(ctx, kycToken.Msg, kycToken.Signature)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("validate signature failed")
		return nil, err
	}

	ctx.Logger.Debug().Str("address", kpu.Address).Msg("kyc token is valid")
	return kpu, nil
}

func checkSignature(ctx *configandcontext.APIContext, msg, signature string) (*KeyPairUser, error) {
	// 1. calc msg hash
	hMsg := crypto.Keccak256Hash([]byte(msg))
	signatureBytes := common.FromHex(signature)

	// 2. get public key
	publicKey, err := crypto.SigToPub(hMsg.Bytes(), signatureBytes)
	if err != nil {
		ctx.Logger.Error().Err(err).Msg("recovery public key from signature failed")
		return nil, err
	}

	// 3. get address
	address := crypto.PubkeyToAddress(*publicKey).Hex()
	eip55Address := address
	address = strings.ToLower(address)

	// 4. verify signature
	// notes: signature last byte is recovery ID
	valid := crypto.VerifySignature(crypto.CompressPubkey(publicKey), hMsg.Bytes(), signatureBytes[:64])
	if !valid {
		ctx.Logger.Error().Msg("invalid signature, cannot be verified")
		return nil, errors.New("invalid signature")
	}

	publicKeyHex := hex.EncodeToString(crypto.CompressPubkey(publicKey))

	kpu := &KeyPairUser{
		Valid:        true,
		Address:      address,
		EIP55Address: eip55Address,
		PublicKeyHex: publicKeyHex,
		RawMsg:       msg,
		Signature:    signature,
	}

	return kpu, nil
}
