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
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	util2 "github.com/Zilliqa/gozilliqa-sdk/util"
	"strings"
	"testing"
)

func TestGetDerivedKey(t *testing.T) {
	p := NewPbkdf2()
	keys := p.GetDerivedKey([]byte("stronk_password"), util2.DecodeHex("0f2274f6c0daf36d5822d97985be5a3d881d11e2e741bad4e038a099eecc3b6d"), 262144, 32)
	iv, _ := keytools.GenerateRandomBytes(16)

	macArray := util2.GenerateMac(keys, util2.DecodeHex("dc55047d51f795509ffb6969db837a4481887ccfb6bfb7c259fb77b19078c2a4"), iv)
	println()
	if strings.Compare(strings.ToLower("DEDC361C53C421974C2811F7F989BC530AEBF9A90C487B4161E0E54AE6FABA31"), strings.ToLower(util2.EncodeHex(macArray))) != 0 {
		t.Error("get derived key error")
	}
}
