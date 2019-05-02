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

func TestGetPublicKeyFromPrivateKey(t *testing.T) {
	privateKey := "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"
	publicKey := GetPublicKeyFromPrivateKey(LaksaGo.DecodeHex(privateKey),false)
	println(LaksaGo.EncodeHex(publicKey))
}
