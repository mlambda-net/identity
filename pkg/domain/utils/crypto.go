package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"os"
)

func Encrypt(data string) (string, error) {
	secret := os.Getenv("SECRET_KEY")
	key, _ := hex.DecodeString(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return "", nil
	}
	nonce := make([]byte, gcm.NonceSize())

	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	plain := []byte(data)
	ciphertext := gcm.Seal(nonce, nonce,plain, nil)

	return hex.EncodeToString(ciphertext), nil

}

func Decrypt(data string) (string, error) {
  secret := os.Getenv("SECRET_KEY")
  key, _ := hex.DecodeString(secret)
  block, err := aes.NewCipher(key)
  if err != nil {
    return "", err
  }

  gcm, err := cipher.NewGCM(block)
  if err != nil {
    return "", err
  }

  plain,_ := hex.DecodeString(data)
  nonceSize := gcm.NonceSize()

  nonce, ciphertext := plain[:nonceSize], plain[nonceSize:]
  plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
  if err != nil {
    return "nil", err
  }

  return string(plaintext), nil
}
