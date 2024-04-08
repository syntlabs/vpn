package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

const SALT_SIZE uint8 = 255
const AES_STD_SIZE uint8 = 255

func hashMnem(mnem []string) string {
	
	sequence := strings.Join(mnem, " ")
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
	cipherText := make([]byte, len([]byte(data)))
	
	cfb := cipher.NewCFBEncrypter(block, cfb_bytes)
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

func generateMneumonic(size int) []string {

	file, ferr := os.Open(wordsPath)
	
	if ferr != nil {
		raise(Err.val.CantOpenFile)
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var randomWords []string

	for scanner.Scan() {
		
		line := scanner.Text()
		
		if rand.Float64() <= float64(size/wordsCount) {
			_ = append(randomWords, line)
		}
		
		if len(randomWords) == 12 {
			break
		} else {
			continue
		}
	}
	return randomWords
}