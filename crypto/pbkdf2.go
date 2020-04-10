/*
 * Copyright (C) 2019 Zilliqa
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
package crypto

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

type pbkdf2Wapper struct {
}

func NewPbkdf2() *pbkdf2Wapper {
	return &pbkdf2Wapper{}
}

func (c *pbkdf2Wapper) GetDerivedKey(password, salt []byte, iterationCount, keySize int) []byte {
	return pbkdf2.Key(password, salt, iterationCount, keySize, sha256.New)
}

type Pbkdf2Params struct {
	Salt  []byte
	DkLen int
	Count int
}
