package core

import (
	"crypto/elliptic"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"math/big"
)

type ECPointSerialize struct {
	BIGNumSerialize
}

// x and y represent the point on the curve
func (ec *ECPointSerialize) SetNumber(dst []byte, offset uint, size uint, x, y *big.Int) {
	bytes := elliptic.MarshalCompressed(keytools.Secp256k1, x, y)
	bnValue := new(big.Int).SetBytes(bytes)
	ec.BIGNumSerialize.SetNumber(dst, offset, size, bnValue)
}
