package utils

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var logger = Logger

func GetPublicKey(privKey string) string {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
}

func EvmPrivateKeyFromString(privKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
		return nil, err
	}
	return privateKey, nil
}

func hashMessage(message string) []byte {
	messageHash := crypto.Keccak256Hash([]byte(message))
	_bytes := []byte{0x19}
	_bytes = append(_bytes, []byte("Ethereum Signed Message:\n32")...)
	_bytes = append(_bytes, messageHash.Bytes()...)
	return crypto.Keccak256Hash(_bytes).Bytes()
}

func Hash(message string) []byte {
	return crypto.Keccak256Hash([]byte(message)).Bytes()
}

func StringToBigInt(message string) string {
	// n := new(big.Int).SetBytes(Hash(message))
	// return new(big.Int).Mod(n, big.NewInt(10))
	return new(big.Int).SetBytes(Hash(message)).String()
}

func Sign(message string, privKey string) ([]byte, string) {

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
	}

	hash := hashMessage(message)

	// logger.WithFields(logrus.Fields{"action": "crypto.Sign", "message": message}).Infof("Message hash: %s", hash.Hex())

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		logger.Fatal(err)
	}
	// signer, err := crypto.Ecrecover(hash.Bytes(), signature[:len(signature)-1])

	return signature, hexutil.Encode(signature)
}

func GetSigner(message string, signature string) (string, error) {
	decoded, err := hexutil.Decode(signature)
	if err != nil {
		logger.Debug(err)
		return "", err
	}
	hash := hashMessage(message)
	if decoded[crypto.RecoveryIDOffset] == 27 || decoded[crypto.RecoveryIDOffset] == 28 {
		decoded[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	signer, err := crypto.SigToPub(hash, decoded)
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(*signer).Hex(), nil
}

func VerifySignature(signer string, message string, signature string) bool {
	// logger.Info("message:::", message)
	decodedSigner, err := GetSigner(message, signature)
	if err != nil {
		return false
	}
	logger.Infof("signer decoded signer %s %s : %T", decodedSigner, signer, (decodedSigner == signer))
	return strings.EqualFold(decodedSigner, signer)
}

func GetNodeHeight() int {
	return 10
}
