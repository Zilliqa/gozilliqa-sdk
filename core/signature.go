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

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) Serialize(data []byte, offset uint) []byte {
	bns := BIGNumSerialize{}
	data = bns.SetNumber(data, offset, signatureChallengeSize, s.R)
	data = bns.SetNumber(data, offset+signatureChallengeSize, signatureChallengeSize, s.S)
	return data
}

func NewFromByteArray(bytes []byte) *Signature {
	rb := make([]byte, 32)
	sb := make([]byte, 32)
	copy(rb, bytes[0:32])
	copy(sb, bytes[32:])

	r := new(big.Int).SetBytes(rb)
	s := new(big.Int).SetBytes(sb)

	return &Signature{
		R: r,
		S: s,
	}
}

type CoSignatures struct {
	CS1 *Signature
	B1  []bool
	CS2 *Signature
	B2  []bool
}
