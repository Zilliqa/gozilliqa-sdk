package crypto

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

func GetDerivedKey(password, salt []byte, iterationCount, keySize int) []byte {
	return pbkdf2.Key(password, salt, iterationCount, keySize, sha256.New)
}
