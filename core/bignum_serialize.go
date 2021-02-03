package core

import (
	"math/big"
)

type BIGNumSerialize struct{}

func (bn *BIGNumSerialize) SetNumber(dst []byte, offset uint, size uint, value *big.Int) {
	// check for offset overflow
	if offset+size < size {
		// overflow detected
		return
	}

	// todo any different from openssl function
	bytes := value.Bytes()
	actualByteNumber := len(bytes)
	if actualByteNumber <= int(size) {
		if offset+size > uint(len(dst)) {
			dst = make([]byte, int(offset+size))
		}
		// pad with zeroes as needed
		sizeDiff := size - uint(actualByteNumber)
		for i := uint(0); i < sizeDiff; i++ {
			dst[offset+i] = 0x00
		}

		for i := 0; i < actualByteNumber; i++ {
			dst[offset+sizeDiff+uint(i)] = bytes[i]
		}
	} else {
		// big num size > declared size
	}

}
