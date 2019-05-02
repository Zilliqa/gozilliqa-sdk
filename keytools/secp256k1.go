package keytools

import (
	"crypto/rand"
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/btcsuite/btcd/btcec"
	"io"
)

var (
	Secp256k1 = btcec.S256()
)

type PrivateKey [32]byte

func GeneratePrivateKey() (PrivateKey, error) {
	pvk := [32]byte{}
	_, err := io.ReadFull(rand.Reader, pvk[:])

	if err != nil {
		return pvk, err
	}
	return PrivateKey(pvk), nil
}

func GetPublicKeyFromPrivateKey(privateKey []byte,compress bool) []byte {
	x, y := Secp256k1.ScalarBaseMult(privateKey)
	return LaksaGo.Compress(Secp256k1, x, y,compress)
}
