package crypto

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

type pbkdf2Wapper struct {
}

func NewPbkdf2() *pbkdf2Wapper {
	return &pbkdf2Wapper{}
}

func (c *pbkdf2Wapper) GetDerivedKey(password, salt []byte, iterationCount, keySize int) []byte {
	return pbkdf2.Key(password, salt, iterationCount, keySize, sha256.New)
}

type Pbkdf2Params struct {
	Salt []byte
	DkLen int
	Count int
}
