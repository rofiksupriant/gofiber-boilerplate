package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"io"
)

func Encrypt(plainText string) (*string, error) {
	block, err := aes.NewCipher([]byte("thisis32bitlongpassphraseimusing"))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	encryptedTextEncoded := base64.URLEncoding.EncodeToString(encryptedText)
	return &encryptedTextEncoded, nil
}

func Decrypt(encryptedTextEncoded string) (*string, error) {
	// decode from plain base64 to []byte
	encryptedText, err := base64.URLEncoding.DecodeString(encryptedTextEncoded)
	if err != nil {
		log.Errorf("error decode base64 %v", err)
		return nil, fiber.ErrInternalServerError
	}

	// init cipher
	block, err := aes.NewCipher([]byte("thisis32bitlongpassphraseimusing"))
	if err != nil {
		log.Errorf("error initialize cipher block %v", err)
		return nil, err
	}

	// init aes Galois Counter Mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Errorf("error initialize aes gcm")
		return nil, err
	}

	// is nonce-size valid
	nonceSize := aesGCM.NonceSize()
	if len(encryptedText) < nonceSize {
		log.Errorf("invalid noncesize token")
		return nil, fiber.ErrInternalServerError
	}

	// separate nonce and plain
	nonce, encryptedPlainText := encryptedText[:nonceSize], encryptedText[nonceSize:]

	// decrypt
	plainTextBytes, err := aesGCM.Open(nil, nonce, encryptedPlainText, nil)
	if err != nil {
		log.Errorf("failed decrypt token %v", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	plainText := string(plainTextBytes)
	return &plainText, nil
}
