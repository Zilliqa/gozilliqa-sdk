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
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"testing"
)

func TestTransaction_TrackTx(t *testing.T) {
	provider := provider.NewProvider("https://dev-api.zilliqa.com/")
	tx := Transaction{
		ID:           "",
		Version:      "",
		Nonce:        "",
		Amount:       "",
		GasPrice:     "",
		GasLimit:     "",
		Signature:    "",
		Receipt:      TransactionReceipt{
			Success:       false,
			CumulativeGas: "",
			EpochNum:      "",
		},
		SenderPubKey: "",
		ToAddr:       "",
		Code:         "",
		Data:         "",
		Status:       0,
	}

	tx.Confirm("846cda64971e259b1739bf15710758803abcf5754507af5af3f779777cd1b0b0",1000,3,provider)

}
