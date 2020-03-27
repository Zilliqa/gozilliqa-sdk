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
	"os"
	"testing"
)

func SkipIfCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestGetNetworkId(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetNetworkId()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetBlockchainInfo(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetBlockchainInfo()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetShardingStructure(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetShardingStructure()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetDsBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetDsBlock("40")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetLatestDsBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetLatestDsBlock()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumDSBlocks(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetNumDSBlocks()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetSmartContractSubState(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://api.zilliqa.com")
	response, err := provider.GetSmartContractSubState("9611c53BE6d1b32058b2747bdeCECed7e1216793", "admins", []interface{}{})
	if err != nil {
		t.Error(err.Error())
	}
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetDSBlockRate(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetDSBlockRate()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestDSBlockListing(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.DSBlockListing(1)
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTxBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetTxBlock("40")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetLatestTxBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetLatestTxBlock()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTxBlocks(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetNumTxBlocks()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTxBlockRate(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetTxBlockRate()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestTxBlockListing(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.TxBlockListing(1)
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTransactions(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetNumTransactions()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTransactionRate(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetTransactionRate()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetCurrentMiniEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetCurrentMiniEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetCurrentDSEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetCurrentDSEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetPrevDifficulty(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetPrevDifficulty()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetPendingTxn(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetPendingTxn("2cf109b25f2132c08a4248e2be8add6b95b92aef5b2c77e737faefbc9353ee7c")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetTotalCoinSupply(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetTotalCoinSupply()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetPrevDSDifficulty(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetPrevDSDifficulty()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTransaction(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetTransaction("655107c300e86ee6e819af1cbfce097db1510e8cd971d99f32ce2772dcad42f2")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetRecentTransactions(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetRecentTransactions()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetTransactionsForTxBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetTransactionsForTxBlock("1")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTxnsTxEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetNumTxnsTxEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetNumTxnsDSEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetNumTxnsDSEpoch()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetMinimumGasPrice(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetMinimumGasPrice()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContractCode(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetSmartContractCode("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContractInit(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetSmartContractInit("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContractState(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetSmartContractState("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetSmartContracts(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetSmartContracts("1eefc4f453539e5ee732b49eb4792b268c2f3908")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetContractAddressFromTransactionID(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetContractAddressFromTransactionID("AAF3089596437A7C6984FA2627B6F38B5F5B80FAEAAC6993C2E82C6A8EE2615E")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetBalance(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_CreateTransactionRaw(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://api.zilliqa.com/")
	payload := []byte(`{"version":65537,"nonce":1,"toAddr":"39550aB45D74cCe5feF70e857c1326b2d9bEE096","pubKey":"03bb0637134af801bcc912f7cf61448aed05fea21f4a6460a7f15a48c8704f2aea","amount":"0","gasPrice":"1000000000","gasLimit":"40000","code":"","data":"{\"_tag\":\"ProxyTransfer\",\"params\":[{\"vname\":\"to\",\"type\":\"ByStr20\",\"value\":\"0x0200a288be83e2a2061d7519d3397b3c6da05f29\"},{\"vname\":\"value\",\"type\":\"Uint128\",\"value\":\"10000000\"}]}","signature":"edf8d36c24e1d9e8e6832aa8e513690afe6df2f076756b45286200dfa02e202525ad0a7349b9c00f10ceb8ce2b71e9cfc745a679e68a6b034d581383155213b0","priority":false}`)
	response, _ := provider.CreateTransactionRaw(payload)
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}
