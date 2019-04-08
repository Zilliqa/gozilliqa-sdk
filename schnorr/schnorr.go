package schnorr

import "math/big"

func Sign(privateKey *big.Int, message [32]byte) ([64]byte, error) {
	//todo
}

func Verify(publicKey [33]byte, message [32]byte, signature [64]byte) (bool, error) {
	//todo
}

