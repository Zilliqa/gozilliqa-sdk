/*
 * Copyright (C) 2021 Zilliqa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package core

import (
	"math/big"
)

type BIGNumSerialize struct{}

func (bn *BIGNumSerialize) SetNumber(dst []byte, offset uint, size uint, value *big.Int) []byte {
	// check for offset overflow
	if offset+size < size {
		// overflow detected
		return nil
	}

	bytes := value.Bytes()
	actualByteNumber := len(bytes)
	if actualByteNumber <= int(size) {
		if offset+size > uint(len(dst)) {
			newDst := make([]byte, int(offset+size))
			copy(newDst, dst)
			dst = newDst
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
	return dst
}
