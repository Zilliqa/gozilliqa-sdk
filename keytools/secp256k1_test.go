package keytools

import (
	"github.com/FireStack-Lab/LaksaGo"
	"strings"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	for i := 0; i < 100000; i++ {
		privateKey, err := GeneratePrivateKey()
		if err != nil {
			panic("cannot generate private key")
		}

		prikeys := LaksaGo.EncodeHex(privateKey[:])
		if len(prikeys) != 64 {
			panic("generate private key error")
		}
		println("private key = " + prikeys)
		publickKey := GetPublicKeyFromPrivateKey(LaksaGo.DecodeHex(prikeys), true)
		pubkeys := LaksaGo.EncodeHex(publickKey)
		if len(pubkeys) != 66 {
			panic("generate public key error")
		}
		println("public key = " + pubkeys)

	}
}

func TestGetPublicKeyFromPrivateKey(t *testing.T) {
	privateKey := "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"
	publicKey := GetPublicKeyFromPrivateKey(LaksaGo.DecodeHex(privateKey), false)
	if strings.Compare(LaksaGo.EncodeHex(publicKey), "04163fa604c65aebeb7048c5548875c11418d6d106a20a0289d67b59807abdd299d4cf0efcf07e96e576732dae122b9a8ac142214a6bc133b77aa5b79ba46b3e20") != 0 {
		t.Error("wrong public key")
	}
}

func TestGetAddressFromPublic(t *testing.T) {
	publicKey := "0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a"
	address := GetAddressFromPublic(LaksaGo.DecodeHex(publicKey))
	println(address)
	if strings.Compare("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a", address) != 0 {
		t.Error("wrong address")
	}
}
