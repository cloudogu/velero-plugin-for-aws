package main

import (
	"crypto/aes"
	"crypto/cipher"
)

type testCipher struct {
	block cipher.Block
}

func newTestCipher(key string) (*testCipher, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	return &testCipher{
		block: block,
	}, nil
}

func (t *testCipher) Encrypt(plaintext []byte) ([]byte, error) {
	ciphertext := make([]byte, len(plaintext))
	t.block.Encrypt(ciphertext, plaintext)
	return ciphertext, nil
}

func (t *testCipher) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext := make([]byte, len(ciphertext))
	t.block.Decrypt(plaintext, ciphertext)
	return plaintext, nil
}
