package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	gey = []byte("super-secret-key")
)

func hashCookie(cookievalue string, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	io.WriteString(hash, cookievalue)
	return hash.Sum(nil)
}

func verifyCookie(cookieToVerify []byte, podpis []byte, key []byte) bool {
	h := hmac.New(sha256.New, key)
	h.Write(cookieToVerify)
	calculated := h.Sum(nil)
	return hmac.Equal(calculated, podpis)
}

func main() {
	text := "testtext"
	var hashed []byte
	hashed = hashCookie(text, gey)

	fmt.Println(fmt.Sprintf("%x", hashed))

	b := verifyCookie([]byte(text), hashed, gey)

	fmt.Println(b)
}
