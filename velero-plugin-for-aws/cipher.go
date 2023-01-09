package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type aesCipher struct {
	gcm cipher.AEAD
	key []byte
}

func newAesCipher(key string) (*aesCipher, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return &aesCipher{
		gcm: aesGCM,
		key: keyBytes,
	}, nil
}

func (t *aesCipher) Encrypt(plainBytes []byte) ([]byte, error) {
	nonce := make([]byte, t.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherBytes := t.gcm.Seal(nonce, nonce, plainBytes, nil)
	return cipherBytes, nil
}

func (t *aesCipher) Decrypt(cipherBytes []byte) ([]byte, error) {
	nonceSize := t.gcm.NonceSize()
	nonce, cipherBytesWithoutNounce := cipherBytes[:nonceSize], cipherBytes[nonceSize:]

	plainBytes, err := t.gcm.Open(nil, nonce, cipherBytesWithoutNounce, nil)
	if err != nil {
		return nil, err
	}

	return plainBytes, nil
}
