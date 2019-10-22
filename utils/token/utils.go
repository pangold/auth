package token

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	return generateRandomString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
}

func generateRandomString(bytes string, length int) string {
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, []byte(bytes)[r.Intn(len(bytes))])
	}
	return string(result)
}

