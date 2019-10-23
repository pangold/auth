package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string{
	return generateRandomString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
}

func GenerateRandomNumber(length int) string {
	return generateRandomString("0123456789", length)
}

func generateRandomString(bytes string, length int) string {
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, []byte(bytes)[r.Intn(len(bytes))])
	}
	return string(result)
}

func GenerateMD5String(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

