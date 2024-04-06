package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const SALT_SIZE uint8 = 255
const AES_STD_SIZE uint8 = 255

func hashUnique(timestep int64, salt string) string {
	sequence := fmt.Sprintf("%d%s", timestep, salt)
	seqhash := sha512.Sum512([]byte(sequence))
	return hex.EncodeToString(seqhash[:])
}

func salt(sizeOf uint8) string {

	nonce := make([]byte, sizeOf)
	hash := sha512.Sum512(nonce)

	return hex.EncodeToString(hash[:])
}

func encryptFile(key string, data string) (string, error) {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		main.raise(Err.cyph.Block)
	}
	cfb_bytes := make([]byte, AES_STD_SIZE)
	cfb := cipher.NewCFBEncrypter(block, cfb_bytes)
	cipherText := make([]byte, len([]byte(data)))
	cfb.XORKeyStream(cipherText, []byte(data))

	return hex.EncodeToString(cipherText), nil
}

func decryptFile(key string, data string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		main.raise(Err.cyph.Block)
	}
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dst, []byte(data))
	if err != nil {
		main.raise(Err.cyph.Decode)
	}
	dst = dst[:n]
	cfb := cipher.NewCFBDecrypter(block, dst)
	plain_text := make([]byte, n)
	cfb.XORKeyStream(plain_text, []byte(data))

	return string(plain_text), nil
}
