package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptAESGCM(plaintext, key []byte) (ciphertext, nonce []byte, err error) {
	// 创建 AES-256 加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	// 使用 GCM 模式包装 AES 块
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	// 创建一个唯一的 nonce（随机数），长度为 GCM 推荐的 12 字节
	nonce = make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	// 使用 AES-GCM 加密数据
	ciphertext = aesGCM.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// AES-256-GCM 解密
func DecryptAESGCM(ciphertext, key, nonce []byte) (plaintext []byte, err error) {
	// 创建 AES-256 加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 使用 GCM 模式包装 AES 块
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 使用 AES-GCM 解密数据
	plaintext, err = aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
