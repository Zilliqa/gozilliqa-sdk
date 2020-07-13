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
	id, _ := provider.GetNetworkId()
	fmt.Println(id)
}

func TestGetBlockchainInfo(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	info, _ := provider.GetBlockchainInfo()
	fmt.Println(info)
}

func TestGetShardingStructure(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetShardingStructure()
	fmt.Println(result)
}

func TestGetDsBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetDsBlock("40")
	fmt.Println(result)
}

func TestGetLatestDsBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetLatestDsBlock()
	fmt.Println(result)
}

func TestGetNumDSBlocks(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetNumDSBlocks()
	fmt.Println(result)
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
	result, _ := provider.GetDSBlockRate()
	fmt.Println(result)
}

func TestDSBlockListing(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.DSBlockListing(1)
	fmt.Println(result)
}

func TestGetTxBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetTxBlock("40")
	fmt.Println(result)
}

func TestGetLatestTxBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetLatestTxBlock()
	fmt.Println(result)
}

func TestGetNumTxBlocks(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetNumTxBlocks()
	fmt.Println(result)
}

func TestGetTxBlockRate(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetTxBlockRate()
	fmt.Println(result)
}

func TestTxBlockListing(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.TxBlockListing(1)
	fmt.Println(result)
}

func TestGetNumTransactions(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetNumTransactions()
	fmt.Println(result)
}

func TestGetTransactionRate(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetTransactionRate()
	fmt.Println(result)
}

func TestGetCurrentMiniEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetCurrentMiniEpoch()
	fmt.Println(result)
}

func TestGetCurrentDSEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetCurrentDSEpoch()
	fmt.Println(result)
}

func TestGetPrevDifficulty(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetPrevDifficulty()
	fmt.Println(result)
}

func TestProvider_GetPendingTxn(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetPendingTxn("2cf109b25f2132c08a4248e2be8add6b95b92aef5b2c77e737faefbc9353ee7c")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetPendingTxns(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	response, _ := provider.GetPendingTxns()
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestProvider_GetTotalCoinSupply(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetTotalCoinSupply()
	fmt.Println(result)
}

func TestProvider_GetMinerInfo(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://api.zilliqa.com/")
	result, _ := provider.GetMinerInfo("6000")
	fmt.Println(result)
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
	result, _ := provider.GetTransaction("c7d6550a6558edcddbf4b3c7cf14db9f1025200b89bcbcd6a570c84db58d554f")
	fmt.Println(result)
}

func TestProvider_GetTransactionBatch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	transactions, _ := provider.GetTransactionBatch([]string{"c7d6550a6558edcddbf4b3c7cf14db9f1025200b89bcbcd6a570c84db58d554f", "c7d6550a6558edcddbf4b3c7cf14db9f1025200b89bcbcd6a570c84db58d554f"})
	st, _ := json.Marshal(transactions)
	fmt.Println(string(st))
}

func TestGetRecentTransactions(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetRecentTransactions()
	fmt.Println(result)
}

func TestGetTransactionsForTxBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetTransactionsForTxBlock("1442201")
	fmt.Println(result)
}

func TestProvider_GetTxnBodiesForTxBlock(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetTxnBodiesForTxBlock("1364221")
	fmt.Println(result)
}

func TestGetNumTxnsTxEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetNumTxnsTxEpoch()
	fmt.Println(result)
}

func TestGetNumTxnsDSEpoch(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetNumTxnsDSEpoch()
	fmt.Println(result)
}

func TestGetMinimumGasPrice(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetMinimumGasPrice()
	fmt.Println(result)
}

func TestGetSmartContractCode(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetSmartContractCode("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	fmt.Println(result)
}

func TestGetSmartContractInit(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetSmartContractInit("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	fmt.Println(result)
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
	response, _ := provider.GetContractAddressFromTransactionID("5283d3a37d90b960ff2e7c6b2a6e8b0f5e62ed74f63b268b1b9485aa08026551")
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}

func TestGetBalance(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://dev-api.zilliqa.com/")
	result, _ := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a")
	fmt.Println(result)
}

func TestProvider_CreateTransactionRaw(t *testing.T) {
	SkipIfCI(t)
	provider := NewProvider("https://api.zilliqa.com/")
	payload := []byte(`{"version":65537,"nonce":1,"toAddr":"39550aB45D74cCe5feF70e857c1326b2d9bEE096","pubKey":"03bb0637134af801bcc912f7cf61448aed05fea21f4a6460a7f15a48c8704f2aea","amount":"0","gasPrice":"1000000000","gasLimit":"40000","code":"","data":"{\"_tag\":\"ProxyTransfer\",\"params\":[{\"vname\":\"to\",\"type\":\"ByStr20\",\"value\":\"0x0200a288be83e2a2061d7519d3397b3c6da05f29\"},{\"vname\":\"value\",\"type\":\"Uint128\",\"value\":\"10000000\"}]}","signature":"edf8d36c24e1d9e8e6832aa8e513690afe6df2f076756b45286200dfa02e202525ad0a7349b9c00f10ceb8ce2b71e9cfc745a679e68a6b034d581383155213b0","priority":false}`)
	response, _ := provider.CreateTransactionRaw(payload)
	result, _ := json.Marshal(response)
	fmt.Println(string(result))
}
