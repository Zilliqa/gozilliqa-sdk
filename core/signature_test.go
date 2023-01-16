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
	"github.com/Zilliqa/gozilliqa-sdk/v3/util"
	"math/big"
	"strings"
	"testing"
)

func TestSignature_Serialize(t *testing.T) {
	rs := "3AF3D288E830E96FF8ED0769F45ABDA774CD989E2AE32EF9E985C8505F14FF98"
	ss := "E191EB14A70B5B53ADA45AFFF4A04578F5D8BB2B1C8A22985EA159B53826CDE7"
	r, _ := new(big.Int).SetString(rs, 16)
	s, _ := new(big.Int).SetString(ss, 16)
	signature := &Signature{
		r, s,
	}
	data := make([]byte, 0)
	sig := signature.Serialize(data, 0)
	if strings.ToUpper(util.EncodeHex(sig)) != "3AF3D288E830E96FF8ED0769F45ABDA774CD989E2AE32EF9E985C8505F14FF98E191EB14A70B5B53ADA45AFFF4A04578F5D8BB2B1C8A22985EA159B53826CDE7" {
		t.Fail()
	}

	rs = "948AFFFF6E068CA2F2757BFD6085D6E4C3084B038E5533C5927ECB19EA0D329C"
	ss = "DFEE66E2C4799E73F0F778126A23032608408C27C2E7B3FA45A626BB9BDEB53C"
	r, _ = new(big.Int).SetString(rs, 16)
	s, _ = new(big.Int).SetString(ss, 16)
	signature = &Signature{
		r, s,
	}
	data = make([]byte, 0)
	sig = signature.Serialize(data, 0)
	if strings.ToUpper(util.EncodeHex(sig)) != "948AFFFF6E068CA2F2757BFD6085D6E4C3084B038E5533C5927ECB19EA0D329CDFEE66E2C4799E73F0F778126A23032608408C27C2E7B3FA45A626BB9BDEB53C" {
		t.Fail()
	}
}
