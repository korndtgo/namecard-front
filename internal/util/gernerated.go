package util

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	numberBytes  = "1234567890"
	letterBytes  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanumeric = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

//GenerateUUID ...
func GenerateUUID() string {
	return uuid.New().String()
}

// RandStringBytes ...
func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandNumberBytes ...
func RandNumberBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}
	return string(b)
}
