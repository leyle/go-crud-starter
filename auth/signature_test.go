package auth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/leyle/crud-log/pkg/crudlog"
	"github.com/leyle/go-crud-starter/configandcontext"
	"github.com/rs/zerolog"
	"testing"
	"time"
)

const (
	PrivateKeyHex = "ac178d1ef86b0b32dfc3e3dce5a386e23d407deeadb6fe3014ec0e336712f3dd"
	// address = "0xCb34c649dAb213A395164474e943C21B0E2126Af"
	// publicKeyHex = "049c83fc52d46fbe20ca91d6c2ed28a41bb3dfc0cf8915e9cc50b491cf92cc78426c20f37d19be82225c2f0279b1f1411ff3b85c8f36f5c14783f0751dea033d19"
)

var PlainMsg = []byte("ios|1686811276")

func TestCheckSignatureToken(t *testing.T) {
	logger := crudlog.NewConsoleLogger(zerolog.TraceLevel)
	ctx := &configandcontext.APIContext{
		Logger: &logger,
	}

	// generate signature
	msg := fmt.Sprintf("ios|%d", time.Now().Unix())
	signature := createSignature(t, []byte(msg))

	// generate token
	kycT := &KYCToken{
		Msg:       msg,
		Signature: signature,
	}

	data, err := json.Marshal(&kycT)
	if err != nil {
		t.Fatal(err)
	}

	token := hex.EncodeToString(data)

	result, err := CheckSignatureToken(ctx, token)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("address", result.Address)
	t.Log("public key", result.PublicKeyHex)
	t.Log("raw msg", result.RawMsg)
	t.Log("token", token)
	t.Log("data", string(data))
}

func createSignature(t *testing.T, msg []byte) string {
	privateKey, err := crypto.HexToECDSA(PrivateKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	// create a hash of the msg
	hash := crypto.Keccak256Hash(msg)

	// sign the hash
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		t.Fatal(err)
	}

	// print hex signature
	signHex := hex.EncodeToString(signature)
	t.Log("signature", signHex)
	return signHex
}

func TestCheckSignature(t *testing.T) {
	// convert private key hex to ECDSA private key
	privateKey, err := crypto.HexToECDSA(PrivateKeyHex)
	if err != nil {
		t.Fatal(err)
	}

	// create a hash of the msg
	hash := crypto.Keccak256Hash(PlainMsg)

	// sign the hash
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		t.Fatal(err)
	}

	// print hex signature
	signHex := hex.EncodeToString(signature)
	t.Log("signature", signHex)

	// then using the message and signature, to verify the signature and get back address
	vHash := crypto.Keccak256Hash(PlainMsg)
	vSignature := common.FromHex(signHex)

	// vPublicKeyBytes, err := crypto.Ecrecover(vHash.Bytes(), vSignature)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	vPublickeyECDSA, err := crypto.SigToPub(vHash.Bytes(), vSignature)
	if err != nil {
		t.Fatal(err)
	}

	vAddress := crypto.PubkeyToAddress(*vPublickeyECDSA).Hex()
	t.Log("recovery address:", vAddress)

	t.Log("v signature", hex.EncodeToString(vSignature))

	// result := crypto.VerifySignature(vPublicKeyBytes, vHash.Bytes(), vSignature)
	result := crypto.VerifySignature(crypto.CompressPubkey(vPublickeyECDSA), vHash.Bytes(), vSignature[:64])
	t.Log("result:", result)
}

func TestCreateSignature(t *testing.T) {
	// create key pair and signature
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	// get associated public key
	publicKey := privateKey.Public()

	// convert public key to its eth address format
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("invalid public key format")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	t.Log("address", address)

	msg := []byte("hello|world|ios")

	// create msg hash
	msgHash := crypto.Keccak256Hash(msg)

	// create msg signature
	signature, err := crypto.Sign(msgHash.Bytes(), privateKey)
	if err != nil {
		t.Fatal("create signature failed")
	}

	t.Log("signature", signature)
	t.Log("signature hex", hex.EncodeToString(signature))
}

func TestCreateKeyPair(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	// Get the corresponding public key
	publicKey := privateKey.Public()

	// Convert the public key to its Ethereum address format
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("invalid public key type")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// Print the generated key pair and Ethereum address
	t.Log("address: ", address)

	// print private key
	pd := crypto.FromECDSA(privateKey)
	ph := hex.EncodeToString(pd)
	t.Log("private key", ph)

	// public key
	pubk := crypto.FromECDSAPub(publicKeyECDSA)
	pubkh := hex.EncodeToString(pubk)
	t.Log("public key", pubkh)
}

func TestKeyPairUser_AesKey(t *testing.T) {
	addr := "0xCb34c649dAb213A395164474e943C21B0E2126Af"

	last := addr[len(addr)-20:]
	t.Log(last)
}
