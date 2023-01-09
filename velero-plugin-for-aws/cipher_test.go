package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
	"time"
)

const testKey = "aler,amz3daps.f9hgandkal4dsxk3d0"

func Test_testCipher_decrypt_encrypt(t1 *testing.T) {
	t1.Run("encrypt_decrypt", func(t *testing.T) {
		// given
		cipher, err := newTestCipher(testKey)
		require.NoError(t, err)

		// when
		encrypt, err := cipher.Encrypt([]byte("this is a testee"))
		require.NoError(t, err)
		decrypt, err := cipher.Decrypt(encrypt)
		require.NoError(t, err)

		// then
		assert.Equal(t, string(decrypt), "this is a testee")
	})
	t1.Run("encrypt_decrypt_gzip", func(t *testing.T) {
		// given
		to_encrypt := []byte("data to encrypt")
		var b bytes.Buffer
		gz := gzip.NewWriter(&b)
		gz.Name = "a-new-hope.txt"
		gz.Comment = "an epic space opera by George Lucas"
		date := time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)
		gz.ModTime = date.UTC()

		cipher, err := newTestCipher(testKey)
		require.NoError(t, err)

		// when
		_, err = gz.Write(to_encrypt)
		require.NoError(t, err)

		err = gz.Close()
		require.NoError(t, err)

		gzipped, err := io.ReadAll(&b)
		require.NoError(t, err)

		encrypted, err := cipher.Encrypt(gzipped)
		require.NoError(t, err)
		decrypt, err := cipher.Decrypt(encrypted)
		require.NoError(t, err)

		gzReader, err := gzip.NewReader(bytes.NewReader(decrypt))
		require.NoError(t, err)

		assert.Equal(t, "a-new-hope.txt", gzReader.Name)
		assert.Equal(t, "an epic space opera by George Lucas", gzReader.Comment)
		assert.Equal(t, date.UTC(), gzReader.ModTime.UTC())

		unzipped, err := io.ReadAll(gzReader)
		require.NoError(t, err)

		// then
		assert.Equal(t, "data to encrypt", string(unzipped))
	})
}
