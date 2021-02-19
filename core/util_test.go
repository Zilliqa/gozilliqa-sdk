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
