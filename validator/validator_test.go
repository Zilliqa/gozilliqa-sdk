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
package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPublicKey(t *testing.T) {
	publicKey := "039E43C9810E6CC09F46AAD38E716DAE3191629534967DC457D3A687D2E2CDDC6A"
	assert.Equal(t, true, IsPublicKey(publicKey))
	assert.Equal(t, false, IsPublicKey(publicKey[0:64]))
}

func TestIsPrivateKey(t *testing.T) {
	privateKey := "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"
	assert.Equal(t, true, IsPrivateKey(privateKey))
	assert.Equal(t, false, IsPrivateKey(privateKey[0:63]))
}

func TestIsAddress(t *testing.T) {
	address := "B5C2CDD79C37209C3CB59E04B7C4062A8F5D5271"
	assert.Equal(t, true, IsAddress(address))
	assert.Equal(t, false, IsAddress(address[0:38]))
}

func TestIsSignature(t *testing.T) {
	signature := "3AF3D288E830E96FF8ED0769F45ABDA774CD989E2AE32EF9E985C8505F14FF98E191EB14A70B5B53ADA45AFFF4A04578F5D8BB2B1C8A22985EA159B53826CDE7"
	assert.Equal(t, true, IsSignature(signature))
	assert.Equal(t, false, IsSignature(signature[0:38]))
}

func TestIsChecksumAddress(t *testing.T) {
	checksumAddress := "0x4BAF5faDA8e5Db92C3d3242618c5B47133AE003C"
	assert.Equal(t, true, IsChecksumAddress(checksumAddress))
	assert.Equal(t, false, IsChecksumAddress(checksumAddress[0:38]))
}

func TestIsBech32(t *testing.T) {
	assert.True(t, IsBech32("zil16jrfrs8vfdtc74yzhyy83je4s4c5sqrcasjlc4"))
}
