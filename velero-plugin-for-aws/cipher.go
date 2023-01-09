package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

const testKey = "test"

type testCipher struct {
	block cipher.Block
}

func newTestCipher() *testCipher {
	block, err := aes.NewCipher([]byte(testKey))
	if err != nil {
		fmt.Printf("Error reading key: %s\n", err.Error())
		os.Exit(1)
	}

	return &testCipher{
		block: block,
	}
}

func (t *testCipher) encrypt(plaintext []byte) []byte {
	ciphertext := []byte{}
	t.block.Encrypt(ciphertext, plaintext)
	return ciphertext
}

func (t *testCipher) decrypt(ciphertext []byte) []byte {
	plaintext := []byte{}
	t.block.Decrypt(plaintext, ciphertext)
	return plaintext
}
