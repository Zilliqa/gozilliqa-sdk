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
package keytools

import (
	"crypto/rand"
	"math/big"

	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/btcsuite/btcd/btcec"
)

var (
	Secp256k1 = btcec.S256()
)

type PrivateKey [32]byte

func GeneratePrivateKey() (PrivateKey, error) {
	var bytes [32]byte
	for {
		privk, err := btcec.NewPrivateKey(Secp256k1)
		if err == nil {
			pvkInt := privk.D
			if pvkInt.Cmp(big.NewInt(0)) == 1 && pvkInt.Cmp(Secp256k1.N) == -1 {
				privk.D.FillBytes(bytes[:])
				break
			}
		}
	}
	return bytes, nil
}

func GetPublicKeyFromPrivateKey(privateKey []byte, compress bool) []byte {
	x, y := Secp256k1.ScalarBaseMult(privateKey)
	return util.Compress(Secp256k1, x, y, compress)
}

func GetAddressFromPublic(publicKey []byte) string {
	originAddress := util.EncodeHex(util.Sha256(publicKey))
	return originAddress[24:]
}

func GetAddressFromPrivateKey(privateKey []byte) string {
	publicKey := GetPublicKeyFromPrivateKey(privateKey, true)
	return GetAddressFromPublic(publicKey)
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
