package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"go-grow/config"
	"io"

	"github.com/spf13/viper"
)

func AESMethodEncyrpt(stringPassword string) (string, error) {
	config.GetEnvConfig()

	bytePassword := []byte(stringPassword)
	byteKey := []byte(viper.GetString("SECRET_KEY"))
	fmt.Println(byteKey)

	aesCipher, err := aes.NewCipher(byteKey)
	if err != nil {
		return err.Error(), err
	}

	cipherPassword := make([]byte, aes.BlockSize+len(bytePassword))
	iv := cipherPassword[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return err.Error(), err
	}

	encrpytedPassword := cipher.NewCFBEncrypter(aesCipher, iv)
	encrpytedPassword.XORKeyStream(cipherPassword[aes.BlockSize:], bytePassword)

	return base64.RawStdEncoding.EncodeToString(cipherPassword), nil
}