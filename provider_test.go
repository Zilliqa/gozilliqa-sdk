package LaksaGo

import (
	"fmt"
	"testing"
)

// func TestGetNetworkId(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetNetworkId()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetBlockchainInfo(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetBlockchainInfo()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetShardingStructure(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetShardingStructure()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetDsBlock(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetDsBlock("40")
// 	fmt.Printf("%v\n", response)
// }

// func TestGetLatestDsBlock(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetLatestDsBlock()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetNumDSBlocks(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetNumDSBlocks()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetDSBlockRate(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetDSBlockRate()
// 	fmt.Printf("%v\n", response)
// }

// func TestDSBlockListing(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.DSBlockListing(1)
// 	fmt.Printf("%v\n", response)
// }

func TestGetTxBlock(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetTxBlock("40")
	fmt.Printf("%v\n", response)
}

// func TestGetLatestTxBlock(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetLatestTxBlock()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetNumTxBlocks(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetNumTxBlocks()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetTxBlockRate(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetTxBlockRate()
// 	fmt.Printf("%v\n", response)
// }

// func TestTxBlockListing(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.TxBlockListing()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetNumTransactions(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetNumTransactions()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetTransactionRate(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetTransactionRate()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetCurrentMiniEpoch(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetCurrentMiniEpoch()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetCurrentDSEpoch(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetCurrentDSEpoch()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetPrevDifficulty(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetPrevDifficulty()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetPrevDSDifficulty(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetPrevDSDifficulty()
// 	fmt.Printf("%v\n", response)
// }

// func TestCreateTransaction(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.CreateTransaction()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetTransaction(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetTransaction()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetRecentTransactions(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetRecentTransactions()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetTransactionsForTxBlock(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetTransactionsForTxBlock()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetNumTxnsTxEpoch(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetNumTxnsTxEpoch()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetNumTxnsDSEpoch(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetNumTxnsDSEpoch()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetMinimumGasPrice(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetMinimumGasPrice()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetSmartContractCode(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetSmartContractCode()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetSmartContractInit(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetSmartContractInit()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetSmartContractState(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetSmartContractState()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetSmartContracts(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetSmartContracts()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetContractAddressFromTransactionID(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetContractAddressFromTransactionID()
// 	fmt.Printf("%v\n", response)
// }

// func TestGetBalance(t *testing.T) {
// 	provider := NewProvider("https://dev-api.zilliqa.com/")

// 	response := provider.GetBalance()
// 	fmt.Printf("%v\n", response)
// }
