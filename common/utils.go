package common

import (
	"crypto/sha256"
	"encoding/json"
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
