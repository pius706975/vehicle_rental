package libs

import (
	"crypto/rand"
	"math/big"
)

func GeneratePaymentCode() string {

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 10)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}

		b[i] = charset[n.Int64()]
	}

	return string(b)
}
