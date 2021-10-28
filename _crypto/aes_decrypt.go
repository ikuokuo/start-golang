package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {
	// aes key: either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256
	key, _ := hex.DecodeString("40295a286556544f695d59306544246d395072567857273541514e6972242435")
	ciphertext, _ := hex.DecodeString("a4ed3611a2a60a8b3ef25a0cd2aed7b7ad6359253ae58173c96bd016daabd2e2fd6e955386f0e690d039")

	// fmt.Printf("\033[1;35mKEY\033[0m\n%s\n", key)
	fmt.Printf("\033[1;35mCIPHER\033[0m\n%x\n", ciphertext)

	plaintext, err := aesDecrypt(key, ciphertext)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\033[1;35mPLAIN\033[0m\n%s\n", plaintext)
}

// NewGCM (Decrypt)
//  https://pkg.go.dev/crypto/cipher#example-NewGCM-Decrypt
func aesDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	textSize := len(ciphertext)
	nonceSize := aesgcm.NonceSize()
	if textSize < nonceSize {
		return nil, errors.New(fmt.Sprintf("ciphertext size %d < nonce size %d",
			textSize, nonceSize))
	}
	n := textSize - nonceSize
	nonce, ciphertext := ciphertext[n:], ciphertext[:n]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
