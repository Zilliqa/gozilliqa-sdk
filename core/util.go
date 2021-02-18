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
	"bytes"
	"encoding/binary"
	"math/big"
	"net"
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

// extract a number from the source byte stream at the specific offset.
// Uint128 -> 16 (bytes)
// Uint64 -> 8 (bytes)
// Uint32 -> 4 (bytes)
func ByteArrayToUint(src []byte, offset uint, numericTypeLen uint) *big.Int {
	resultNum := new(big.Int)
	if offset+numericTypeLen <= uint(len(src)) {
		leftShift := (numericTypeLen - 1) * 8
		for i := uint(0); i < numericTypeLen; i++ {
			tmp := new(big.Int).SetBytes(src[offset+i : offset+i+1])
			tmp = new(big.Int).Lsh(tmp, leftShift)
			resultNum = new(big.Int).And(resultNum, tmp)
			leftShift -= 8
		}
	}
	return resultNum
}

func IP2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.LittleEndian, &long)
	return long
}
