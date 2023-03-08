package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func PKCS7Padding(data []byte, blocksize int) []byte {
	padding := blocksize - len(data)%blocksize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, paddingText...)
}

func PKCS7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func AesEncryptByte(data []byte, key string) string {
	k := append([]byte(key), bytes.Repeat([]byte("0"), 32)...)[:32]
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])

	data = PKCS7Padding(data, blockSize)
	encryted := make([]byte, len(data))
	blockMode.CryptBlocks(encryted, data)

	return base64.StdEncoding.EncodeToString(encryted)
}

func AesEncrypt(text string, key string) string {
	return AesEncryptByte([]byte(text), key)
}

func AesDecryptByte(encryted []byte, key string) string {
	k := append([]byte(key), bytes.Repeat([]byte("0"), 32)...)[:32]
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])

	data := make([]byte, len(encryted))
	blockMode.CryptBlocks(data, encryted)
	return string(PKCS7Unpadding(data))
}

func AesDecrypt(encryted string, key string) string {
	encrytedBytes, _ := base64.StdEncoding.DecodeString(encryted)
	return AesDecryptByte(encrytedBytes, key)
}
