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

type BitVector struct{}

func (b *BitVector) GetBitVectorLengthInBytes(lengthInBits uint) uint {
	a := lengthInBits & 0x07
	if a > 0 {
		return (lengthInBits >> 3) + 1
	} else {
		return lengthInBits >> 3
	}
}

func (b *BitVector) GetBitVectorSerializedSize(lengthInBits uint) uint {
	return 2 + b.GetBitVectorLengthInBytes(lengthInBits)
}

func (b *BitVector) SetBitVector(dst []byte, offset uint, value []bool) uint {
	lengthNeeded := b.GetBitVectorSerializedSize(uint(len(value)))
	if (offset + lengthNeeded) > uint(len(dst)) {
		newDst := make([]byte, offset+lengthNeeded)
		copy(newDst, dst)
		dst = newDst
	}

	for i := offset; i < offset+lengthNeeded; i++ {
		dst[i] = 0x00
	}

	dst[offset] = byte(len(value) >> 8)
	dst[offset+1] = byte(len(value))

	index := uint(0)
	for _, b := range value {
		if b {
			dst[offset+2+(index>>3)] |= 1 << (7 - (index & 0x07))
		}
		index++
	}

	return lengthNeeded
}
