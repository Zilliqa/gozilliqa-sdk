package keytools

import (
	"github.com/FireStack-Lab/LaksaGo"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		panic("cannot generate private key")
	}
	println(LaksaGo.EncodeHex(privateKey[:]))
}
