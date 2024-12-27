package services

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"

	"jpcpereira93/go-api-gateway/config"
)

// Get the AES-256-GCM Cipher and IV for the configured secrets.
func getCipherIv() (cipher.AEAD, []byte) {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal().AnErr("Error loading configuration: %v", err)
	}

	iv := []byte(cfg.Secret.Iv)
	key := []byte(cfg.Secret.Key)

	// Ensure key size is 32 bytes for AES-256
	if len(key) != 32 {
		panic(errors.New("key size must be 32 bytes for AES-256"))
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Create GCM (Galois/Counter Mode) instance
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// Ensure IV size matches GCM nonce size
	if len(iv) != gcm.NonceSize() {
		err := fmt.Errorf("invalid IV size: must be %d bytes", gcm.NonceSize())
		panic(err);
	}

	return gcm, iv
}

// Cipher data using AES-256-GCM with the configured key and IV.
func Cipher(data []byte) string {
	// Get the cipher and IV with the configured secret key and IV
	gcm, iv := getCipherIv()

	// Encrypt and append the authentication tag
	ciphertext := gcm.Seal(nil, iv, data, nil)

	// Return base64-encoded ciphertext and IV for safe transport
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// Decipher data using AES-256-GCM with the configured  key and IV.
func Decipher(data string) (string, error) {
	// Get the cipher and IV with the configured secret key and IV
	gcm, iv := getCipherIv()

	// Decode the base64-encoded data
	decodedData, err := base64.URLEncoding.DecodeString(data)

	if err != nil {
		return "", err
	}

	// Decrypt and verify the authentication tag
	plaintext, err := gcm.Open(nil, iv, decodedData, nil)
	
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}