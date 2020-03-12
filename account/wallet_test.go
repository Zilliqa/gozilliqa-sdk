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
package account

import (
	"fmt"
	provider2 "github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strconv"
	"testing"
)

func TestWallet_SignWith(t *testing.T) {
	wallet := NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	tx := &transaction.Transaction{}
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")
	_ = wallet.SignWith(tx, "9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a", *provider)
}

func TestSendTransaction(t *testing.T) {
	wallet := NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")

	gasPrice := provider.GetMinimumGasPrice().Result.(string)

	tx := &transaction.Transaction{
		Version:      strconv.FormatInt(int64(util.Pack(333, 1)), 10),
		SenderPubKey: "0246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A",
		ToAddr:       "4BAF5faDA8e5Db92C3d3242618c5B47133AE003C",
		Amount:       "10000000",
		GasPrice:     gasPrice,
		GasLimit:     "1",
		Code:         "",
		Data:         "",
		Priority:     false,
	}

	err := wallet.Sign(tx, *provider)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	rsp, err := provider.CreateTransaction(tx.ToTransactionPayload())

	if err != nil {
		t.Error(err.Error())
	}

	if rsp.Error != nil {
		fmt.Println(rsp.Error)
		t.Error(err)
	} else {
		result := rsp.Result.(map[string]interface{})
		hash := result["TranID"].(string)
		fmt.Printf("hash is %s\n", hash)
		tx.Confirm(hash, 1000, 3, provider)
	}
}
