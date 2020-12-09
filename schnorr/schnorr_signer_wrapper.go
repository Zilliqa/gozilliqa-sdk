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

/*
 * This module is a mini wrapper around the main schnorr module used for
 * zilliqa rosetta project. It includes the generation of the random bytes
 * to make the signing process more convenient.
 */
package go_schnorr

import (
	"encoding/hex"
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"github.com/Zilliqa/gozilliqa-sdk/util"
)

func SignMessage(privateKey []byte, publicKey []byte, message []byte) ([]byte, error) {
	// generate a random byte k
	rb, err := keytools.GenerateRandomBytes(keytools.Secp256k1.N.BitLen() / 8)

	if err != nil {
		return nil, err
	}

	r, s, err2 := TrySign(privateKey, publicKey, message, rb)

	if err2 != nil {
		return nil, err2
	}

	sig := fmt.Sprintf("%064s%064s", util.EncodeHex(r), util.EncodeHex(s))
	sigBytes, err := hex.DecodeString(sig)

	if err != nil {
		panic("cannot convert hex string to byte array")
	}

	return sigBytes, nil
}

func VerifySignature(publicKey []byte, message []byte, signature []byte) bool {
	sig := util.EncodeHex(signature)

	if len(sig) != 128 {
		return false
	}

	r := util.DecodeHex(sig[0:64])
	s := util.DecodeHex(sig[64:128])

	return Verify(publicKey, message, r, s)
}
