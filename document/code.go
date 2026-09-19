package document

import (
	"crypto/rand"
)

const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode() (string, error) {

	code := make([]byte, 6)

	for i := range code {

		randomByte := make([]byte, 1)

		_, err := rand.Read(randomByte)
		if err != nil {
			return "", err
		}

		code[i] = characters[int(randomByte[0])%len(characters)]
	}

	return string(code), nil
}
