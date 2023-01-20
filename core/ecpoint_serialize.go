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
	"crypto/elliptic"
	"github.com/Zilliqa/gozilliqa-sdk/v3/keytools"
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
