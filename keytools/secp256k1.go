package keytools

import (
	"crypto/rand"
	"io"
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

