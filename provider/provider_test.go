package provider

import (
	"fmt"
	"testing"
)

func TestGetNetworkId(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNetworkId()
	fmt.Printf("%v\n", response)
}

func TestGetBlockchainInfo(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetBlockchainInfo()
	fmt.Printf("%v\n", response)
}

func TestGetShardingStructure(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetShardingStructure()
	fmt.Printf("%v\n", response)
}

func TestGetDsBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetDsBlock("40")
	fmt.Printf("%v\n", response)
}

func TestGetLatestDsBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetLatestDsBlock()
	fmt.Printf("%v\n", response)
}

func TestGetNumDSBlocks(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumDSBlocks()
	fmt.Printf("%v\n", response)
}

func TestProvider_GetSmartContractSubState(t *testing.T) {
	provider := NewProvider("https://api.zilliqa.com")
	response, err := provider.GetSmartContractSubState("9611c53BE6d1b32058b2747bdeCECed7e1216793", "admins", []interface{}{})
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(response)
}

func TestNewProvider(t *testing.T) {

	//params := []interface{}{
	//	"a", 1, []interface{}{},
	//}
	//
	//b,_  := json.Marshal()
	//fmt.Println(string(b))
}

func TestGetDSBlockRate(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetDSBlockRate()
	fmt.Printf("%v\n", response)
}

func TestDSBlockListing(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.DSBlockListing(1)
	fmt.Printf("%v\n", response)
}

func TestGetTxBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTxBlock("40")
	fmt.Printf("%v\n", response)
}

func TestGetLatestTxBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetLatestTxBlock()
	fmt.Printf("%v\n", response)
}

func TestGetNumTxBlocks(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTxBlocks()
	fmt.Printf("%v\n", response)
}

func TestGetTxBlockRate(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTxBlockRate()
	fmt.Printf("%v\n", response)
}

func TestTxBlockListing(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.TxBlockListing(1)
	fmt.Printf("%v\n", response)
}

func TestGetNumTransactions(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTransactions()
	fmt.Printf("%v\n", response)
}

func TestGetTransactionRate(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTransactionRate()
	fmt.Printf("%v\n", response)
}

func TestGetCurrentMiniEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetCurrentMiniEpoch()
	fmt.Printf("%v\n", response)
}

func TestGetCurrentDSEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetCurrentDSEpoch()
	fmt.Printf("%v\n", response)
}

func TestGetPrevDifficulty(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetPrevDifficulty()
	fmt.Printf("%v\n", response)
}

func TestGetPrevDSDifficulty(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetPrevDSDifficulty()
	fmt.Printf("%v\n", response)
}

func TestGetTransaction(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTransaction("655107c300e86ee6e819af1cbfce097db1510e8cd971d99f32ce2772dcad42f2")
	fmt.Printf("%v\n", response)
}

func TestGetRecentTransactions(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetRecentTransactions()
	fmt.Printf("%v\n", response)
}

func TestGetTransactionsForTxBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTransactionsForTxBlock("1")
	fmt.Printf("%v\n", response)
}

func TestGetNumTxnsTxEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTxnsTxEpoch()
	fmt.Printf("%v\n", response)
}

func TestGetNumTxnsDSEpoch(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetNumTxnsDSEpoch()
	fmt.Printf("%v\n", response)
}

func TestGetMinimumGasPrice(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetMinimumGasPrice()
	fmt.Printf("%v\n", response)
}

func TestGetSmartContractCode(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContractCode("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	fmt.Printf("%v\n", response)
}

func TestGetSmartContractInit(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContractInit("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	fmt.Printf("%v\n", response)
}

func TestGetSmartContractState(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContractState("fe001824823b12b58708bf24edd94d8b5e1cfcf7")
	fmt.Printf("%v\n", response)
}

func TestGetSmartContracts(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetSmartContracts("1eefc4f453539e5ee732b49eb4792b268c2f3908")
	fmt.Printf("%v\n", response)
}

func TestGetContractAddressFromTransactionID(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetContractAddressFromTransactionID("AAF3089596437A7C6984FA2627B6F38B5F5B80FAEAAC6993C2E82C6A8EE2615E")
	fmt.Printf("%v\n", response)
}

func TestGetBalance(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a")
	fmt.Printf("%v\n", response)
}
