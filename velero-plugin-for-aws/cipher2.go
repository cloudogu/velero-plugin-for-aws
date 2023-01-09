package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type testCipher2 struct {
	gcm cipher.AEAD
	key []byte
}

func newTestCipher2(key string) (*testCipher2, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	return &testCipher2{
		gcm: aesGCM,
		key: keyBytes,
	}, nil
}

func (t *testCipher2) Encrypt(plainBytes []byte) ([]byte, error) {
	nonce := make([]byte, t.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherBytes := t.gcm.Seal(nonce, nonce, plainBytes, nil)
	return cipherBytes, nil
}

func (t *testCipher2) Decrypt(cipherBytes []byte) ([]byte, error) {
	nonceSize := t.gcm.NonceSize()
	nonce, cipherBytesWithoutNounce := cipherBytes[:nonceSize], cipherBytes[nonceSize:]

	plainBytes, err := t.gcm.Open(nil, nonce, cipherBytesWithoutNounce, nil)
	if err != nil {
		return nil, err
	}

	return plainBytes, nil
}
