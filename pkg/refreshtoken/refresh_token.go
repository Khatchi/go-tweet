package refreshtoken

import (
	"crypto/rand"
	"encoding/hex"
)

func GetRefreshToken() (string, error) {
	b := make([]byte, 18)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
