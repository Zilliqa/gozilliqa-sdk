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
package transaction

import (
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEncodeTransactionProto(t *testing.T) {
	txParams := TxParams{
		Version:      "0",
		Nonce:        "0",
		ToAddr:       "2E3C9B415B19AE4035503A06192A0FAD76E04243",
		SenderPubKey: "0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a",
		Amount:       "10000",
		GasPrice:     "100",
		GasLimit:     "1000",
		Code:         "",
		Data:         "",
	}

	bytes, _ := EncodeTransactionProto(txParams)
	assert.Equal(t, strings.ToUpper(util.EncodeHex(bytes)), "080010001A142E3C9B415B19AE4035503A06192A0FAD76E0424322230A210246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A2A120A100000000000000000000000000000271032120A100000000000000000000000000000006438E807")
}
