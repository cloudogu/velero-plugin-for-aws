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

func (t *testCipher) encrypt(plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	t.block.Encrypt(ciphertext, plaintext)
	return ciphertext
}

func (t *testCipher) decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	t.block.Decrypt(plaintext, ciphertext)
	return plaintext
}
