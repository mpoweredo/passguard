package utils

import "crypto/rand"

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateId(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes)
}
