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

func TestScryptWapper_GetDerivedKey(t *testing.T) {
	s := NewScryptWapper()
	keys, err := s.GetDerivedKey([]byte("stronk_password"), util2.DecodeHex("2c37db13a633c5a5e5b8c699109690e33860b7eb43bbc81bbab47d4e9c29f1b9"),
		8192, 8, 1, 32)
	if err != nil {
		t.Error("scrypt: get derived key")
	}
	iv, err := keytools.GenerateRandomBytes(16)
	if err != nil {
		t.Error(err.Error())
	}

	macArray := util2.GenerateMac(keys, util2.DecodeHex("ecdf81453d031ac2fa068b7185ddac044fa4632d3b061400d3c07a86510b4823"),iv)
	if strings.Compare(strings.ToLower(util2.EncodeHex(macArray)), "ed7fa37a4adbc8b7bbe0d43a329a047f89e2dcf7f2dfc96babfe79edd955f7a3") != 0 {
		t.Error("crypt: get derived key")
	}
}
