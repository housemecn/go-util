package crypto

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
)

// @Project: go-util
// @Author: houseme
// @Description:
// @File: aes
// @Version: 1.0.0
// @Date: 2021/4/5 23:39
// @Package go_util

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, paddingText...)
}

func pKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unPadding := int(origData[length-1])
    return origData[:(length - unPadding)]
}

// GoAES 加密
type GoAES struct {
    Key []byte
}

// NewGoAES 返回GoAES
func NewGoAES(key []byte) *GoAES {
    return &GoAES{Key: key}
}

// Encrypt 加密数据
func (a *GoAES) Encrypt(origData []byte) ([]byte, error) {
    block, err := aes.NewCipher(a.Key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    origData = pKCS7Padding(origData, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, a.Key[:blockSize])
    encrypted := make([]byte, len(origData))
    blockMode.CryptBlocks(encrypted, origData)
    return encrypted, nil
}

// Decrypt 解密数据
func (a *GoAES) Decrypt(encrypted []byte) ([]byte, error) {
    block, err := aes.NewCipher(a.Key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, a.Key[:blockSize])
    origData := make([]byte, len(encrypted))
    blockMode.CryptBlocks(origData, encrypted)
    origData = pKCS7UnPadding(origData)
    return origData, nil
}
