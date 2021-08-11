package cryptobench

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

var (
	// Generate a 2048-bits key
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
)

func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	// This method requires a random number of bits.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// The public key is part of the PrivateKey struct
	return privateKey, &privateKey.PublicKey
}

func init() {
	privateKey, publicKey = generateKeyPair(2048)	
}

func EncryptOAEP(msg string) (string, error) {
	encryptMsg, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(msg), nil)
	return string(encryptMsg), err
}

func DecryptOAEP(encryptMsg string) (string, error) {
	msg, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, []byte(encryptMsg), nil)
	return string(msg), err
}

func EncryptPKCS1v15(msg string) (string, error) {
	encryptMsg, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(msg))
	return string(encryptMsg), err
}

func DecryptPKCS1v15(encryptMsg string) (string, error) {
	msg, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(encryptMsg))
	return string(msg), err
}
