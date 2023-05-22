package libs

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
)

func CodeCrypt(len int) (string, error) {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("could not has a password %w", err)
	}

	return base32.StdEncoding.EncodeToString(bytes)[:len], nil
}