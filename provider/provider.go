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
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/ybbus/jsonrpc"
	"io/ioutil"
	"net/http"
)

type Provider struct {
	host      string
	rpcClient jsonrpc.RPCClient
}

func NewProvider(host string) *Provider {

	rpcClient := jsonrpc.NewClient(host)
	return &Provider{host: host, rpcClient: rpcClient}
}

// Returns the CHAIN_ID of the specified network. This is represented as a String.
func (provider *Provider) GetNetworkId() (string, error) {
	result, err := provider.call("GetNetworkId")
	if err != nil {
		return "", err
	}
	if result.Error != nil {
		return "", result.Error
	}
	return result.Result.(string), nil
}

// Returns the current network statistics for the specified network.
func (provider *Provider) GetBlockchainInfo() (*core.BlockchainInfo, error) {
	result, err := provider.call("GetBlockchainInfo")
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, result.Error
	}

	var blockchainInfo core.BlockchainInfo
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &blockchainInfo)
	if err3 != nil {
		return nil, err3
	}

	return &blockchainInfo, nil

}

func (provider *Provider) GetShardingStructure() (*core.ShardingStructure, error) {
	result, err := provider.call("GetShardingStructure")
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var shardingStructure core.ShardingStructure
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &shardingStructure)
	if err3 != nil {
		return nil, err3
	}

	return &shardingStructure, nil

}

// Returns the details of a specified Directory Service block.
func (provider *Provider) GetDsBlock(block_number string) (*core.DsBlockT, error) {
	result, err := provider.call("GetDsBlock", block_number)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var dsBlock core.DsBlockT

	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &dsBlock)
	if err3 != nil {
		return nil, err3
	}

	return &dsBlock, nil
}

func (provider *Provider) GetDsBlockVerbose(block_number string) (*core.DsBlockT, error) {
	result, err := provider.call("GetDsBlockVerbose", block_number)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var dsBlock core.DsBlockT

	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &dsBlock)
	if err3 != nil {
		return nil, err3
	}

	return &dsBlock, nil
}

// Returns the details of the most recent Directory Service block.
func (provider *Provider) GetLatestDsBlock() (*core.DSBlock, error) {
	result, err := provider.call("GetLatestDsBlock")
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var dsBlock core.DSBlock

	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &dsBlock)
	if err3 != nil {
		return nil, err3
	}

	return &dsBlock, nil
}

// Returns the current number of validated Directory Service blocks in the network.
// This is represented as a String.
func (provider *Provider) GetNumDSBlocks() (string, error) {
	result, err := provider.call("GetNumDSBlocks")
	if err != nil {
		return "", err
	}
	if result.Error != nil {
		return "", result.Error
	}
	return result.Result.(string), nil
}

// Returns the current Directory Service blockrate per second.
func (provider *Provider) GetDSBlockRate() (float64, error) {
	result, err := provider.call("GetDSBlockRate")
	if err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, result.Error
	}

	rate, err2 := result.Result.(json.Number).Float64()
	if err2 != nil {
		return 0, err2
	}

	return rate, nil
}

// Returns a paginated list of up to 10 Directory Service (DS) blocks and their block hashes for a specified page.
// The maxPages variable that specifies the maximum number of pages available is also returned.
func (provider *Provider) DSBlockListing(ds_block_listing int) (*core.BlockList, error) {
	result, err := provider.call("DSBlockListing", ds_block_listing)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var list core.BlockList
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &list)
	if err3 != nil {
		return nil, err3
	}

	return &list, nil
}

// Returns the details of a specified Transaction block.
func (provider *Provider) GetTxBlock(tx_block string) (*core.TxBlockT, error) {
	result, err := provider.call("GetTxBlock", tx_block)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var txBlock core.TxBlockT

	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &txBlock)
	if err3 != nil {
		return nil, err3
	}

	return &txBlock, nil
}

func (provider *Provider) GetTxBlockVerbose(tx_block string) (*core.TxBlockT, error) {
	result, err := provider.call("GetTxBlockVerbose", tx_block)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var txBlock core.TxBlockT

	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &txBlock)
	if err3 != nil {
		return nil, err3
	}

	return &txBlock, nil
}

// Returns the details of the most recent Transaction block.
func (provider *Provider) GetLatestTxBlock() (*core.TxBlockT, error) {
	result, err := provider.call("GetLatestTxBlock")
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var txBlock core.TxBlockT

	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &txBlock)
	if err3 != nil {
		return nil, err3
	}

	return &txBlock, nil
}

// Returns the current number of Transaction blocks in the network.
// This is represented as a String.
func (provider *Provider) GetNumTxBlocks() (string, error) {
	result, err := provider.call("GetNumTxBlocks")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the current Transaction blockrate per second for the network.
func (provider *Provider) GetTxBlockRate() (float64, error) {
	result, err := provider.call("GetTxBlockRate")
	if err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, result.Error
	}

	rate, err2 := result.Result.(json.Number).Float64()
	if err2 != nil {
		return 0, err2
	}

	return rate, nil
}

// Returns a paginated list of up to 10 Transaction blocks and their block hashes for a specified page.
// The maxPages variable that specifies the maximum number of pages available is also returned.
func (provider *Provider) TxBlockListing(page int) (*core.BlockList, error) {
	result, err := provider.call("TxBlockListing", page)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var list core.BlockList
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &list)
	if err3 != nil {
		return nil, err3
	}

	return &list, nil
}

func (provider *Provider) GetCurrentDSComm() (*core.DSComm, error) {
	result, err := provider.call("GetCurrentDSComm")
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var dscomm core.DSComm

	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &dscomm)
	if err3 != nil {
		return nil, err3
	}

	return &dscomm, nil
}

// Returns the current number of validated Transactions in the network.
// This is represented as a String.
func (provider *Provider) GetNumTransactions() (string, error) {
	result, err := provider.call("GetNumTransactions")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the current Transaction rate per second (TPS) of the network.
// This is represented as an Number.
func (provider *Provider) GetTransactionRate() (float64, error) {
	result, err := provider.call("GetTransactionRate")

	if err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, result.Error
	}

	rate, err2 := result.Result.(json.Number).Float64()
	if err2 != nil {
		return 0, err2
	}

	return rate, nil
}

// Returns the current TX block number of the network.
// This is represented as a String.
func (provider *Provider) GetCurrentMiniEpoch() (string, error) {
	result, err := provider.call("GetCurrentMiniEpoch")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the current number of DS blocks in the network.
// This is represented as a String.
func (provider *Provider) GetCurrentDSEpoch() (string, error) {
	result, err := provider.call("GetCurrentDSEpoch")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the minimum shard difficulty of the previous block.
// This is represented as an Number.
func (provider *Provider) GetPrevDifficulty() (int64, error) {
	result, err := provider.call("GetPrevDifficulty")
	if err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, result.Error
	}

	difficulty, err2 := result.Result.(json.Number).Int64()
	if err2 != nil {
		return 0, err2
	}

	return difficulty, nil
}

// Returns the minimum DS difficulty of the previous block.
// This is represented as an Number.
func (provider *Provider) GetPrevDSDifficulty() (int64, error) {
	result, err := provider.call("GetPrevDSDifficulty")
	if err != nil {
		return 0, err
	}

	if result.Error != nil {
		return 0, result.Error
	}

	difficulty, err2 := result.Result.(json.Number).Int64()
	if err2 != nil {
		return 0, err2
	}

	return difficulty, nil
}

// Returns the total supply (ZIL) of coins in the network. This is represented as a String.
func (provider *Provider) GetTotalCoinSupply() (string, error) {
	result, err := provider.call("GetTotalCoinSupply")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the mining nodes (i.e., the members of the DS committee and shards) at the specified DS block.
// Notes: 1. Nodes owned by Zilliqa Research are omitted. 2. dscommittee has no size field since the DS committee size
// is fixed for a given chain. 3. For the Zilliqa Mainnet, this API is only available from DS block 5500 onwards.
func (provider *Provider) GetMinerInfo(dsNumber string) (*core.MinerInfo, error) {
	result, err := provider.call("GetMinerInfo", dsNumber)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var minerInfo core.MinerInfo
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &minerInfo)
	if err3 != nil {
		return nil, err3
	}

	return &minerInfo, nil
}

// Returns the pending status of a specified Transaction.
func (provider *Provider) GetPendingTxn(tx string) (*core.PendingTxnResult, error) {
	result, err := provider.call("GetPendingTxn", tx)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var pendingResult core.PendingTxnResult
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &pendingResult)
	if err3 != nil {
		return nil, err3
	}

	pendingResult.Info = core.PendingTxnError[pendingResult.Code]

	return &pendingResult, nil

}

// Returns the pending status of all unvalidated Transactions.
func (provider *Provider) GetPendingTxns() (*core.PendingTxns, error) {
	result, err := provider.call("GetPendingTxns")
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var pendingTxns core.PendingTxns
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &pendingTxns)
	if err3 != nil {
		return nil, err3
	}

	for _, tnx := range pendingTxns.Txns {
		tnx.Info = core.PendingTxnError[tnx.Code]
	}

	return &pendingTxns, err
}

// Create a new Transaction object and send it to the network to be process.
func (provider *Provider) CreateTransaction(payload TransactionPayload) (*jsonrpc.RPCResponse, error) {
	return provider.call("CreateTransaction", &payload)
}

func (provider *Provider) CreateTransactionBatch(payloads [][]TransactionPayload) (jsonrpc.RPCResponses, error) {
	var requests jsonrpc.RPCRequests
	for _, payload := range payloads {
		r := jsonrpc.NewRequest("CreateTransaction", payload)
		requests = append(requests, r)
	}
	return provider.rpcClient.CallBatch(requests)
}

func (provider *Provider) CreateTransactionRaw(payload []byte) (*jsonrpc.RPCResponse, error) {
	var pl TransactionPayload
	err := json.Unmarshal(payload, &pl)
	if err != nil {
		panic(err.Error())
	}
	return provider.call("CreateTransaction", &pl)
}

// Returns the details of a specified Transaction.
// Note: If the transaction had an data field or code field, it will be displayed
func (provider *Provider) GetTransaction(transaction_hash string) (*core.Transaction, error) {
	result, err := provider.call("GetTransaction", transaction_hash)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var transaction core.Transaction
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &transaction)
	if err3 != nil {
		return nil, err3
	}

	return &transaction, nil
}

func (provider *Provider) GetTransactionStatus(transactionHash string) (*core.TransactionStatus, error) {
	result, err := provider.call("GetTransactionStatus", transactionHash)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var status core.TransactionStatus
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &status)
	if err3 != nil {
		return nil, err3
	}

	return &status, nil
}

func (provider *Provider) GetTransactionBatch(transactionHashes []string) ([]*core.Transaction, error) {
	var requests jsonrpc.RPCRequests
	for _, hash := range transactionHashes {
		r := jsonrpc.NewRequest("GetTransaction", []string{hash})
		requests = append(requests, r)
	}

	results, err := provider.rpcClient.CallBatch(requests)
	if err != nil {
		return nil, err
	}

	var transactions []*core.Transaction

	for _, result := range results {
		var transaction core.Transaction
		jsonString, err2 := json.Marshal(result.Result)
		if err2 != nil {
			return transactions, err2
		}
		err3 := json.Unmarshal(jsonString, &transaction)
		if err3 != nil {
			return transactions, err3
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil

}

// Returns the most recent 100 transactions that are validated by the Zilliqa network.
func (provider *Provider) GetRecentTransactions() (*core.Transactions, error) {
	result, err := provider.call("GetRecentTransactions")
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var transactions core.Transactions
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &transactions)
	if err3 != nil {
		return nil, err3
	}

	return &transactions, nil
}

// Returns the validated transactions included within a specfied final transaction block as an array of length i,
// where i is the number of shards plus the DS committee. The transactions are grouped based on the group that processed
// the transaction. The first element of the array refers to the first shard. The last element of the array at index, i,
// refers to the transactions processed by the DS Committee.
func (provider *Provider) GetTransactionsForTxBlock(tx_block_number string) ([][]string, error) {
	result, err := provider.call("GetTransactionsForTxBlock", tx_block_number)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var transactions [][]string
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &transactions)
	if err3 != nil {
		return nil, err3
	}

	return transactions, nil
}

func (provider *Provider) GetTxnBodiesForTxBlock(tx_block_number string) ([]core.Transaction, error) {
	result, err := provider.call("GetTxnBodiesForTxBlock", tx_block_number)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var transactions []core.Transaction
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &transactions)
	if err3 != nil {
		return nil, err3
	}

	return transactions, nil
}

// Returns the number of validated transactions included in this Transaction epoch.
// This is represented as String.
func (provider *Provider) GetNumTxnsTxEpoch() (string, error) {
	result, err := provider.call("GetNumTxnsTxEpoch")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the number of validated transactions included in this DS epoch.
// This is represented as String.
func (provider *Provider) GetNumTxnsDSEpoch() (string, error) {
	result, err := provider.call("GetNumTxnsDSEpoch")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the minimum gas price for this DS epoch, measured in the smallest price unit Qa (or 10^-12 Zil) in Zilliqa.
// This is represented as a String.
func (provider *Provider) GetMinimumGasPrice() (string, error) {
	result, err := provider.call("GetMinimumGasPrice")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil

}

// Returns the Scilla code associated with a smart contract address.
// This is represented as a String.
func (provider *Provider) GetSmartContractCode(contract_address string) (string, error) {
	result, err := provider.call("GetSmartContractCode")
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the initialization (immutable) parameters of a given smart contract, represented in a JSON format.
func (provider *Provider) GetSmartContractInit(contract_address string) ([]core.ContractValue, error) {
	result, err := provider.call("GetSmartContractInit", contract_address)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	var init []core.ContractValue
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &init)
	if err3 != nil {
		return nil, err3
	}

	return init, nil
}

// Returns the state (mutable) variables of a smart contract address, represented in a JSON format.
func (provider *Provider) GetSmartContractState(contract_address string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetSmartContractState", contract_address)
}

func (provider *Provider) GetSmartContractStateBatch(contract_addresses []string) (jsonrpc.RPCResponses, error) {
	var requests jsonrpc.RPCRequests
	for _, payload := range contract_addresses {
		r := jsonrpc.NewRequest("GetSmartContractState", payload)
		requests = append(requests, r)
	}
	return provider.rpcClient.CallBatch(requests)
}

// Returns the state (or a part specified) of a smart contract address, represented in a JSON format.
func (provider *Provider) GetSmartContractSubState(contractAddress string, params ...interface{}) (string, error) {
	//we should hack here for now
	type req struct {
		Id      string      `json:"id"`
		Jsonrpc string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
	}

	p := []interface{}{
		contractAddress,
	}

	for _, v := range params {
		p = append(p, v)
	}

	r := &req{
		Id:      "1",
		Jsonrpc: "2.0",
		Method:  "GetSmartContractSubState",
		Params:  p,
	}

	b, _ := json.Marshal(r)
	reader := bytes.NewReader(b)
	request, err := http.NewRequest("POST", provider.host, reader)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil

}

// Returns the list of smart contract addresses created by an User's account and the contracts' latest states.
func (provider *Provider) GetSmartContracts(user_address string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetSmartContracts", user_address)
}

// Returns a smart contract address of 20 bytes. This is represented as a String.
// NOTE: This only works for contract deployment transactions.
func (provider *Provider) GetContractAddressFromTransactionID(transaction_id string) (string, error) {
	result, err := provider.call("GetContractAddressFromTransactionID", transaction_id)
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", result.Error
	}

	return result.Result.(string), nil
}

// Returns the current balance of an account, measured in the smallest accounting unit Qa (or 10^-12 Zil).
// This is represented as a String
// Returns the current nonce of an account. This is represented as an Number.
func (provider *Provider) GetBalance(user_address string) (*core.BalanceAndNonce, error) {
	result, err := provider.call("GetBalance", user_address)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, result.Error
	}

	balanceAndNonce := core.BalanceAndNonce{
		Balance: "0",
		Nonce:   0,
	}
	jsonString, err2 := json.Marshal(result.Result)
	if err2 != nil {
		return nil, err2
	}

	err3 := json.Unmarshal(jsonString, &balanceAndNonce)
	if err3 != nil {
		return nil, err3
	}

	return &balanceAndNonce, nil
}

func (provider *Provider) GetStateProof(contractAddress string, hashedStorageKey string, blockNum *string) (*core.StateProof, error) {
	type req struct {
		Id      string      `json:"id"`
		Jsonrpc string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
	}

	var blocknum string

	if blockNum == nil {
		blocknum = "latest"
	} else {
		blocknum = *blockNum
	}

	p := []interface{}{
		contractAddress,
		hashedStorageKey,
		blocknum,
	}

	r := &req{
		Id:      "1",
		Jsonrpc: "2.0",
		Method:  "GetStateProof",
		Params:  p,
	}

	b, _ := json.Marshal(r)
	reader := bytes.NewReader(b)
	request, err := http.NewRequest("POST", provider.host, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err2 := client.Do(request)
	if err2 != nil {
		return nil, err2
	}
	defer resp.Body.Close()

	result, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return nil, err3
	}

	type rsp struct {
		Id      string           `json:"id"`
		Jsonrpc string           `json:"jsonrpc"`
		Result  *core.StateProof `json:"result"`
	}
	var stateProof rsp
	err4 := json.Unmarshal(result, &stateProof)
	if err4 != nil {
		return nil, err4
	}

	return stateProof.Result, nil
}

func (provider *Provider) call(method_name string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	response, err := provider.rpcClient.Call(method_name, params)

	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, errors.New("rpc response is nil, please check your network status")
	}

	return response, nil
}
