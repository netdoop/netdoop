package utils

import (
	"math/rand"
	"time"
)

func MergeSlices[T comparable](slices ...[]T) []T {
	result := make([]T, 0)
	seen := make(map[T]bool)

	for _, slice := range slices {
		for _, elem := range slice {
			if !seen[elem] {
				seen[elem] = true
				result = append(result, elem)
			}
		}
	}
	return result
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
