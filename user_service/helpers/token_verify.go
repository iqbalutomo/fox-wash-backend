package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateTokenVerify() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", fmt.Errorf("failed to generate rand token: %v", err)
	}

	encodedToken := base64.StdEncoding.EncodeToString(token)

	return encodedToken, nil
}
