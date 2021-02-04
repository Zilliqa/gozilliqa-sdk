package core

import (
	"math/big"
)

// place a number into the destination byte stream at the specified offset
// caller should make sure that the value of numericTypeLen following rules
// Uint128 -> 16 (bytes)
// Uint64 -> 8 (bytes)
// Uint32 -> 4 (bytes)
func UintToByteArray(dst []byte, offset uint, num *big.Int, numericTypeLen uint) []byte {
	lengthAvailable := uint(len(dst)) - offset
	if lengthAvailable < numericTypeLen {
		newDst := make([]byte, uint(len(dst))+numericTypeLen-lengthAvailable)
		copy(newDst, dst)
		dst = newDst
	}

	rightShift := (numericTypeLen - 1) * 8
	ff, _ := new(big.Int).SetString("FF", 16)
	for i := uint(0); i < numericTypeLen; i++ {
		shifted := new(big.Int).Rsh(num, rightShift)
		anded := new(big.Int).And(shifted, ff)
		dst[offset+i] = byte(anded.Int64())
		rightShift -= 8
	}

	return dst
}
