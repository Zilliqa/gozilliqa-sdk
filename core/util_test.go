package core

import (
	"math/big"
	"testing"
)

func TestUint128ToProtobufByteArray(t *testing.T) {
	// prepare a uint128 number
	a, _ := new(big.Int).SetString("7aec9010a5ca23caaeb63e38b4dc92b2", 16)
	dst := make([]byte, 16)
	dst = UintToByteArray(dst, 0, a, 16)
	t.Log(a)
	t.Log(dst)

	b := new(big.Int).SetInt64(1)
	dst2 := make([]byte, 16)
	dst2 = UintToByteArray(dst2, 0, b, 16)
	t.Log(b)
	t.Log(dst2)
}
