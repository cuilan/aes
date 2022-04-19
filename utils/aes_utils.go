package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AesEncrypt 加密函数
func AesEncrypt(data, key []byte) ([]byte, error) {
	// 创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 判断加密块的大小
	blockSize := block.BlockSize()
	// 填充
	encryptBytes := pkcs7Padding(data, blockSize)
	// 初始化加密数据接收切片
	encrypted := make([]byte, len(encryptBytes))
	// 使用CBC加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// 执行加密
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return encrypted, nil
}

// AesDecrypt 加密函数
func AesDecrypt(data, key []byte) ([]byte, error) {
	// 创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 获取加密块的大小
	blockSize := block.BlockSize()
	// 使用CBC解密模式
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	// 初始化解密数据接收切片
	decrypted := make([]byte, len(data))
	// 执行解密
	blockMode.CryptBlocks(decrypted, data)
	// 去除填充
	decrypted = pkcs7UnPadding(decrypted)
	return decrypted, nil
}

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	// 判断缺少几位长度，最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	// 补足位数，把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 去除填充
func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
