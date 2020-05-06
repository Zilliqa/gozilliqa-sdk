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
	"fmt"
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
func (provider *Provider) GetNetworkId() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetNetworkId")
}

// Returns the current network statistics for the specified network.
func (provider *Provider) GetBlockchainInfo() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetBlockchainInfo")
}

func (provider *Provider) GetShardingStructure() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetShardingStructure")
}

// Returns the details of a specified Directory Service block.
func (provider *Provider) GetDsBlock(block_number string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetDsBlock", block_number)
}

// Returns the details of the most recent Directory Service block.
func (provider *Provider) GetLatestDsBlock() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetLatestDsBlock")
}

// Returns the current number of validated Directory Service blocks in the network.
// This is represented as a String.
func (provider *Provider) GetNumDSBlocks() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetNumDSBlocks")
}

// Returns the current Directory Service blockrate per second.
func (provider *Provider) GetDSBlockRate() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetDSBlockRate")
}

// Returns a paginated list of up to 10 Directory Service (DS) blocks and their block hashes for a specified page.
// The maxPages variable that specifies the maximum number of pages available is also returned.
func (provider *Provider) DSBlockListing(ds_block_listing int) (*jsonrpc.RPCResponse, error) {
	return provider.call("DSBlockListing", ds_block_listing)
}

// Returns the details of a specified Transaction block.
func (provider *Provider) GetTxBlock(tx_block string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetTxBlock", tx_block)
}

// Returns the details of the most recent Transaction block.
func (provider *Provider) GetLatestTxBlock() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetLatestTxBlock")
}

// Returns the current number of Transaction blocks in the network.
// This is represented as a String.
func (provider *Provider) GetNumTxBlocks() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetNumTxBlocks")
}

// Returns the current Transaction blockrate per second for the network.
func (provider *Provider) GetTxBlockRate() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetTxBlockRate")
}

// Returns a paginated list of up to 10 Transaction blocks and their block hashes for a specified page.
// The maxPages variable that specifies the maximum number of pages available is also returned.
func (provider *Provider) TxBlockListing(page int) (*jsonrpc.RPCResponse, error) {
	return provider.call("TxBlockListing", page)
}

// Returns the current number of validated Transactions in the network.
// This is represented as a String.
func (provider *Provider) GetNumTransactions() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetNumTransactions")
}

// Returns the current Transaction rate per second (TPS) of the network.
// This is represented as an Number.
func (provider *Provider) GetTransactionRate() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetTransactionRate")
}

// Returns the current TX block number of the network.
// This is represented as a String.
func (provider *Provider) GetCurrentMiniEpoch() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetCurrentMiniEpoch")
}

// Returns the current number of DS blocks in the network.
// This is represented as a String.
func (provider *Provider) GetCurrentDSEpoch() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetCurrentDSEpoch")
}

// Returns the minimum shard difficulty of the previous block.
// This is represented as an Number.
func (provider *Provider) GetPrevDifficulty() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetPrevDifficulty")
}

func (provider *Provider) GetPendingTxn(tx string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetPendingTxn", tx)
}

// Returns the minimum DS difficulty of the previous block.
// This is represented as an Number.
func (provider *Provider) GetPrevDSDifficulty() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetPrevDSDifficulty")
}

// Returns the total supply (ZIL) of coins in the network. This is represented as a String.
func (provider *Provider) GetTotalCoinSupply() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetTotalCoinSupply")
}

// Create a new Transaction object and send it to the network to be process.
func (provider *Provider) CreateTransaction(payload TransactionPayload) (*jsonrpc.RPCResponse, error) {
	r, _ := json.Marshal(payload)
	fmt.Println(string(r))
	//fmt.Println(payload)
	return provider.call("CreateTransaction", &payload)
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
func (provider *Provider) GetTransaction(transaction_hash string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetTransaction", transaction_hash)
}

// Returns the most recent 100 transactions that are validated by the Zilliqa network.
func (provider *Provider) GetRecentTransactions() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetRecentTransactions")
}

// Returns the validated transactions included within a specfied final transaction block as an array of length i,
// where i is the number of shards plus the DS committee. The transactions are grouped based on the group that processed
// the transaction. The first element of the array refers to the first shard. The last element of the array at index, i,
// refers to the transactions processed by the DS Committee.
func (provider *Provider) GetTransactionsForTxBlock(tx_block_number string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetTransactionsForTxBlock", tx_block_number)
}

func (provider *Provider) GetTxnBodiesForTxBlock(tx_block_number string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetTxnBodiesForTxBlock", tx_block_number)
}

// Returns the number of validated transactions included in this Transaction epoch.
// This is represented as String.
func (provider *Provider) GetNumTxnsTxEpoch() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetNumTxnsTxEpoch")
}

// Returns the number of validated transactions included in this DS epoch.
// This is represented as String.
func (provider *Provider) GetNumTxnsDSEpoch() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetNumTxnsDSEpoch")
}

// Returns the minimum gas price for this DS epoch, measured in the smallest price unit Qa (or 10^-12 Zil) in Zilliqa.
// This is represented as a String.
func (provider *Provider) GetMinimumGasPrice() (*jsonrpc.RPCResponse, error) {
	return provider.call("GetMinimumGasPrice")
}

// Returns the Scilla code associated with a smart contract address.
// This is represented as a String.
func (provider *Provider) GetSmartContractCode(contract_address string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetSmartContractCode", contract_address)
}

// Returns the initialization (immutable) parameters of a given smart contract, represented in a JSON format.
func (provider *Provider) GetSmartContractInit(contract_address string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetSmartContractInit", contract_address)
}

// Returns the state (mutable) variables of a smart contract address, represented in a JSON format.
func (provider *Provider) GetSmartContractState(contract_address string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetSmartContractState", contract_address)
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
		return "", nil
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(result), nil

}

// Returns the list of smart contract addresses created by an User's account and the contracts' latest states.
func (provider *Provider) GetSmartContracts(user_address string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetSmartContracts", user_address)
}

// Returns a smart contract address of 20 bytes. This is represented as a String.
// NOTE: This only works for contract deployment transactions.
func (provider *Provider) GetContractAddressFromTransactionID(transaction_id string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetContractAddressFromTransactionID", transaction_id)
}

// Returns the current balance of an account, measured in the smallest accounting unit Qa (or 10^-12 Zil).
// This is represented as a String
// Returns the current nonce of an account. This is represented as an Number.
func (provider *Provider) GetBalance(user_address string) (*jsonrpc.RPCResponse, error) {
	return provider.call("GetBalance", user_address)
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
