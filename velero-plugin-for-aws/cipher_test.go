package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_newTestCipher(t *testing.T) {

}

func Test_testCipher_decrypt_encrypt(t1 *testing.T) {
	t1.Run("", func(t *testing.T) {
		// given
		cipher, err := newTestCipher()
		require.NoError(t, err)

		// when
		encrypt := cipher.encrypt([]byte("this is a testee"))
		decrypt := cipher.decrypt(encrypt)

		assert.Equal(t, "this is a testee", string(decrypt))
	})
}

func Test_testCipher_encrypt(t1 *testing.T) {

}
