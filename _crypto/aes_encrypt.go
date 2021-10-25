package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// aes key: either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256
	key, _ := hex.DecodeString("40295a286556544f695d59306544246d395072567857273541514e6972242435")
	plaintext := []byte("I love golang!")

	// fmt.Printf("\033[1;35mKEY\033[0m\n%s\n", key)
	fmt.Printf("\033[1;35mPLAIN\033[0m\n%s\n", plaintext)

	ciphertext, err := aesEncrypt(key, plaintext)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\033[1;35mCIPHER\033[0m\n%x\n", ciphertext)
}

// NewGCM (Encrypt)
//  https://pkg.go.dev/crypto/cipher@go1.17.2#example-NewGCM-Encrypt
func aesEncrypt(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(ciphertext, nonce...)
	return ciphertext, nil
}
