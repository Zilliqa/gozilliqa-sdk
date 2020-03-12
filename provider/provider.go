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

func (provider *Provider) GetNetworkId() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetNetworkId")
}

func (provider *Provider) GetBlockchainInfo() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetBlockchainInfo")
}

func (provider *Provider) GetShardingStructure() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetShardingStructure")
}

func (provider *Provider) GetDsBlock(block_number string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetDsBlock", block_number)
}

func (provider *Provider) GetLatestDsBlock() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetLatestDsBlock")
}

func (provider *Provider) GetNumDSBlocks() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetNumDSBlocks")
}

func (provider *Provider) GetDSBlockRate() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetDSBlockRate")
}

func (provider *Provider) DSBlockListing(ds_block_listing int) (*jsonrpc.RPCResponse,error) {
	return provider.call("DSBlockListing", ds_block_listing)
}

func (provider *Provider) GetTxBlock(tx_block string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetTxBlock", tx_block)
}

func (provider *Provider) GetLatestTxBlock() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetLatestTxBlock")
}

func (provider *Provider) GetNumTxBlocks() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetNumTxBlocks")
}

func (provider *Provider) GetTxBlockRate() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetTxBlockRate")
}

func (provider *Provider) TxBlockListing(page int) (*jsonrpc.RPCResponse,error) {
	return provider.call("TxBlockListing", page)
}

func (provider *Provider) GetNumTransactions() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetNumTransactions")
}

func (provider *Provider) GetTransactionRate() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetTransactionRate")
}

func (provider *Provider) GetCurrentMiniEpoch() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetCurrentMiniEpoch")
}

func (provider *Provider) GetCurrentDSEpoch() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetCurrentDSEpoch")
}

func (provider *Provider) GetPrevDifficulty() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetPrevDifficulty")
}

func (provider *Provider) GetPendingTxn(tx string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetPendingTxn", tx)
}

func (provider *Provider) GetPrevDSDifficulty() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetPrevDSDifficulty")
}

func (provider *Provider) GetTotalCoinSupply() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetTotalCoinSupply")
}

func (provider *Provider) CreateTransaction(payload TransactionPayload) (*jsonrpc.RPCResponse,error) {
	r, _ := json.Marshal(payload)
	fmt.Println(string(r))
	//fmt.Println(payload)
	return provider.call("CreateTransaction", &payload)
}

func (provider *Provider) CreateTransactionRaw(payload []byte) (*jsonrpc.RPCResponse,error) {
	var pl TransactionPayload
	err := json.Unmarshal(payload, &pl)
	if err != nil {
		panic(err.Error())
	}
	return provider.call("CreateTransaction", &pl)
}

func (provider *Provider) GetTransaction(transaction_hash string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetTransaction", transaction_hash)
}

func (provider *Provider) GetRecentTransactions() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetRecentTransactions")
}

func (provider *Provider) GetTransactionsForTxBlock(tx_block_number string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetTransactionsForTxBlock", tx_block_number)
}

func (provider *Provider) GetNumTxnsTxEpoch() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetNumTxnsTxEpoch")
}

func (provider *Provider) GetNumTxnsDSEpoch() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetNumTxnsDSEpoch")
}

func (provider *Provider) GetMinimumGasPrice() (*jsonrpc.RPCResponse,error) {
	return provider.call("GetMinimumGasPrice")
}

func (provider *Provider) GetSmartContractCode(contract_address string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetSmartContractCode", contract_address)
}

func (provider *Provider) GetSmartContractInit(contract_address string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetSmartContractInit", contract_address)
}

func (provider *Provider) GetSmartContractState(contract_address string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetSmartContractState", contract_address)
}

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

func (provider *Provider) GetSmartContracts(user_address string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetSmartContracts", user_address)
}

func (provider *Provider) GetContractAddressFromTransactionID(transaction_id string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetContractAddressFromTransactionID", transaction_id)
}

func (provider *Provider) GetBalance(user_address string) (*jsonrpc.RPCResponse,error) {
	return provider.call("GetBalance", user_address)
}

func (provider *Provider) call(method_name string, params ...interface{}) (*jsonrpc.RPCResponse,error) {
	response, err := provider.rpcClient.Call(method_name, params)
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
