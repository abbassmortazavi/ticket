package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateRandomKey() string {
	key := make([]byte, 32)
	read, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	if read != 32 {
		log.Fatal("rand.Read() returned wrong length")
	}
	return base64.URLEncoding.EncodeToString(key)
}

func GenerateRandomKeySession(num int) []byte {
	key := make([]byte, num)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("Failed to generate random key: %v", err)
	}
	return key
}
