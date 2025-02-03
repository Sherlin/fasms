package utilities

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	log "github.com/sirupsen/logrus"
)

// Decrypt decrypts a base64-encoded ciphertext using AES
func Decrypt(encrypted, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %v", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	nonce := data[:aes.BlockSize]
	cipherText := data[aes.BlockSize:]

	plainText := make([]byte, len(cipherText))
	stream := cipher.NewCFBDecrypter(block, nonce)
	stream.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func GetDatabaseUrl()( string, error) {

	key:= os.Getenv("PASSKEY")
	// Read encrypted text from file
	file, err := os.Open("/home/ec2-user/backend/fasms")
	if err != nil {
		log.Error("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	var encrypted string
	if _, err := fmt.Fscan(file, &encrypted); err != nil && err != io.EOF {
		log.Error("Error reading file:", err)
		return "",err
	}

	// Decrypt the content
	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		log.Error("Error decrypting:", err)
		return "",err
	}

	return decrypted, nil
}