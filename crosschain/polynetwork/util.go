/*
 * Copyright (C) 2020 Zilliqa
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
package polynetwork

import (
	"errors"
	"strings"
)

const polyChainPubKeyLen = 67 * 2
const polyChainSignatureLen = 65 * 2

func SplitPubKeys(rawBytes string) ([]string, error) {
	return split(rawBytes, polyChainPubKeyLen)
}

func SplitSignature(rawBytes string) ([]string, error) {
	return split(rawBytes, polyChainSignatureLen)
}

func split(rawBytes string, l int) ([]string, error) {
	if strings.HasPrefix(rawBytes, "0x") {
		rawBytes = rawBytes[2:]
	}
	var keys []string
	keyLen := len(rawBytes)
	if keyLen%l != 0 {
		return keys, errors.New("wrong length of public key list")
	}
	n := keyLen / l
	for i := 0; i < n; i++ {
		publicKey := rawBytes[i*l : (i+1)*l]
		keys = append(keys, "0x"+publicKey)
	}
	return keys, nil
}
