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
	"golang.org/x/crypto/scrypt"
)

type scryptWapper struct {
}

func NewScryptWapper() *scryptWapper {
	return &scryptWapper{}
}

func (s *scryptWapper) GetDerivedKey(password, salt []byte, n, r, p, dkLen int) ([]byte, error) {
	dk, err := scrypt.Key(password, salt, n, r, p, dkLen)
	if err != nil {
		return nil, err
	} else {
		return dk, nil
	}
}

type ScryptParams struct {
	Salt  []byte
	DkLen int
	N     int
	R     int
	P     int
}
