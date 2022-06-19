package common

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
)

const (
	HashLength    = 32
	AddressLength = 20
	BlockTime     = 10
	ServerPort    = ":9000"
)

type Hash [HashLength]byte

type Address [AddressLength]byte

func GetHash(v any) Hash {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("byte conversion failed: %v", err)
	}
	hash := sha256.Sum256(bytes)

	return hash
}

func ToHash(b []byte) Hash {
	var hash Hash
	copy(hash[:], b[:32])
	return hash
}

func RandomHash(len int) string {
	data := make([]byte, len)
	if _, err := rand.Read(data); err != nil {
		log.Printf("error %v", err)
	}
	str := fmt.Sprintf("%x", data)
	log.Printf("hash is %v\n", str)

	return fmt.Sprintf("%x", data)
}
