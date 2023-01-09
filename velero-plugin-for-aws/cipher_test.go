package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

const testKey = "aler,amz3daps.f9hgandkal4dsxk3d0"

func Test_newTestCipher(t *testing.T) {

}

func Test_testCipher_decrypt_encrypt(t1 *testing.T) {
	t1.Run("encrypt_decrypt", func(t *testing.T) {
		// given
		cipher, err := newTestCipher(testKey)
		require.NoError(t, err)

		// when
		encrypt := cipher.encrypt([]byte("this is a testee"))
		decrypt := cipher.decrypt(encrypt)

		// then
		assert.Equal(t, string(decrypt), "this is a testee")
	})
	t1.Run("encrypt_decrypt_gzip", func(t *testing.T) {
		// given
		to_encrypt := []byte("data to encrypt")
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		defer func() {
			err := gz.Close()
			require.NoError(t, err)
		}()
		fmt.Println(b.Bytes())
		cipher, err := newTestCipher(testKey)
		require.NoError(t, err)

		// when
		_, err = gz.Write(to_encrypt)
		require.NoError(t, err)
		gzipped := b.Bytes()

		encrypted := cipher.encrypt(gzipped)
		decrypt := cipher.decrypt(encrypted)

		gzReader, err := gzip.NewReader(bytes.NewReader(decrypt))
		require.NoError(t, err)
		_, err = io.Copy(os.Stdout, gzReader)
		require.NoError(t, err)
		unzipped, err := io.ReadAll(gzReader)

		// then
		assert.Equal(t, "data to encrypt", string(unzipped))
	})
}

func Test_testCipher_encrypt(t1 *testing.T) {

}
