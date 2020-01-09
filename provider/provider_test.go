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
package provider

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetNetworkId(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response := provider.GetNetworkId()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetBlockchainInfo(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response := provider.GetBlockchainInfo()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetShardingStructure(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response := provider.GetShardingStructure()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetDsBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetDsBlock("40")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetLatestDsBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetLatestDsBlock()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumDSBlocks(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumDSBlocks()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetSmartContractSubState(t *testing.T) {
	provider := NewProvider("https://api.zilliqa.com")
	response, err := provider.GetSmartContractSubState("9611c53BE6d1b32058b2747bdeCECed7e1216793", "admins", []interface{}{})
	if err != nil {
		t.Error(err.Error())
	}
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetDSBlockRate(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetDSBlockRate()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestDSBlockListing(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.DSBlockListing(1)
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTxBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTxBlock("40")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetLatestTxBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetLatestTxBlock()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTxBlocks(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTxBlocks()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTxBlockRate(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTxBlockRate()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestTxBlockListing(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.TxBlockListing(1)
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTransactions(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTransactions()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTransactionRate(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTransactionRate()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetCurrentMiniEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetCurrentMiniEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetCurrentDSEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetCurrentDSEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetPrevDifficulty(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetPrevDifficulty()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetPendingTxn(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetPendingTxn("2cf109b25f2132c08a4248e2be8add6b95b92aef5b2c77e737faefbc9353ee7c")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetTotalCoinSupply(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTotalCoinSupply()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetPrevDSDifficulty(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetPrevDSDifficulty()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTransaction(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTransaction("655107c300e86ee6e819af1cbfce097db1510e8cd971d99f32ce2772dcad42f2")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetRecentTransactions(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetRecentTransactions()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTransactionsForTxBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTransactionsForTxBlock("1")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTxnsTxEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTxnsTxEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTxnsDSEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTxnsDSEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetMinimumGasPrice(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetMinimumGasPrice()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContractCode(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContractCode("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContractInit(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContractInit("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContractState(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContractState("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContracts(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContracts("1eefc4f453539e5ee732b49eb4792b268c2f3908")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetContractAddressFromTransactionID(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetContractAddressFromTransactionID("AAF3089596437A7C6984FA2627B6F38B5F5B80FAEAAC6993C2E82C6A8EE2615E")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetBalance(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}
