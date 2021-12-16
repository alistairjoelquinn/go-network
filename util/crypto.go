package util

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomString(n int) (string, error) {
	const values = "0123456789"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(values))))
		if err != nil {
			return "", err
		}
		ret[i] = values[num.Int64()]
	}

	return string(ret), nil
}
