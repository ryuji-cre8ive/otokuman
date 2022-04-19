package plugins

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encryption (word string) string {
	b := []byte(word)
	sha256 := sha256.Sum256(b)
	hashed := hex.EncodeToString(sha256[:])
	return hashed
}