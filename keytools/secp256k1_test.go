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
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	for i := 0; i < 100; i++ {
		privateKey, err := GeneratePrivateKey()
		if err != nil {
			panic("cannot generate private key")
		}
		prikeys := util.EncodeHex(privateKey[:])
		assert.Equal(t, len(prikeys), 64)
		t.Log("private key = " + prikeys)
		publickKey := GetPublicKeyFromPrivateKey(util.DecodeHex(prikeys), true)
		pubkeys := util.EncodeHex(publickKey)
		assert.Equal(t, len(pubkeys), 66)
		t.Log("public key = " + pubkeys)
	}
}

func TestGetPublicKeyFromPrivateKey(t *testing.T) {
	privateKey := "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"
	publicKey := GetPublicKeyFromPrivateKey(util.DecodeHex(privateKey), false)
	assert.Equal(t, util.EncodeHex(publicKey), "04163fa604c65aebeb7048c5548875c11418d6d106a20a0289d67b59807abdd299d4cf0efcf07e96e576732dae122b9a8ac142214a6bc133b77aa5b79ba46b3e20")
}

func TestGetAddressFromPublic(t *testing.T) {
	publicKey := "0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a"
	address := GetAddressFromPublic(util.DecodeHex(publicKey))
	assert.Equal(t, address, "9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a")
}

func TestGetAddressFromPrivateKey(t *testing.T) {
	privateKey := "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"
	addr := GetAddressFromPrivateKey(util.DecodeHex(privateKey))
	assert.Equal(t, addr, "b5c2cdd79c37209c3cb59e04b7c4062a8f5d5271")
}
