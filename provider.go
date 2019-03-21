package LaksaGo

import (
	"github.com/ybbus/jsonrpc"
)

type Provider struct {
	rpcClient jsonrpc.RPCClient
}

func NewProvider(host string) *Provider {
	rpcClient := jsonrpc.NewClient(host)
	return &Provider{rpcClient}
}

func (provider *Provider) GetNetworkId() *jsonrpc.RPCResponse {
	return provider.call("GetNetworkId")
}

func (provider *Provider) GetBlockchainInfo() *jsonrpc.RPCResponse {
	return provider.call("GetBlockchainInfo")
}

func (provider *Provider) GetShardingStructure() *jsonrpc.RPCResponse {
	return provider.call("GetShardingStructure")
}

func (provider *Provider) GetDsBlock() *jsonrpc.RPCResponse {
	return provider.call("GetDsBlock")
}

func (provider *Provider) GetLatestDsBlock() *jsonrpc.RPCResponse {
	return provider.call("GetLatestDsBlock")
}

func (provider *Provider) GetNumDSBlocks() *jsonrpc.RPCResponse {
	return provider.call("GetNumDSBlocks")
}

func (provider *Provider) GetDSBlockRate() *jsonrpc.RPCResponse {
	return provider.call("GetDSBlockRate")
}

func (provider *Provider) DSBlockListing() *jsonrpc.RPCResponse {
	return provider.call("DSBlockListing")
}

func (provider *Provider) GetTxBlock() *jsonrpc.RPCResponse {
	return provider.call("GetTxBlock")
}

func (provider *Provider) GetLatestTxBlock() *jsonrpc.RPCResponse {
	return provider.call("GetLatestTxBlock")
}

func (provider *Provider) GetNumTxBlocks() *jsonrpc.RPCResponse {
	return provider.call("GetNumTxBlocks")
}

func (provider *Provider) GetTxBlockRate() *jsonrpc.RPCResponse {
	return provider.call("GetTxBlockRate")
}

func (provider *Provider) TxBlockListing() *jsonrpc.RPCResponse {
	return provider.call("TxBlockListing")
}

func (provider *Provider) GetNumTransactions() *jsonrpc.RPCResponse {
	return provider.call("GetNumTransactions")
}

func (provider *Provider) GetTransactionRate() *jsonrpc.RPCResponse {
	return provider.call("GetTransactionRate")
}

func (provider *Provider) GetCurrentMiniEpoch() *jsonrpc.RPCResponse {
	return provider.call("GetCurrentMiniEpoch")
}

func (provider *Provider) GetCurrentDSEpoch() *jsonrpc.RPCResponse {
	return provider.call("GetCurrentDSEpoch")
}

func (provider *Provider) GetPrevDifficulty() *jsonrpc.RPCResponse {
	return provider.call("GetPrevDifficulty")
}

func (provider *Provider) GetPrevDSDifficulty() *jsonrpc.RPCResponse {
	return provider.call("GetPrevDSDifficulty")
}

func (provider *Provider) CreateTransaction() *jsonrpc.RPCResponse {
	return provider.call("CreateTransaction")
}

func (provider *Provider) GetTransaction() *jsonrpc.RPCResponse {
	return provider.call("GetTransaction")
}

func (provider *Provider) GetRecentTransactions() *jsonrpc.RPCResponse {
	return provider.call("GetRecentTransactions")
}

func (provider *Provider) GetTransactionsForTxBlock() *jsonrpc.RPCResponse {
	return provider.call("GetTransactionsForTxBlock")
}

func (provider *Provider) GetNumTxnsTxEpoch() *jsonrpc.RPCResponse {
	return provider.call("GetNumTxnsTxEpoch")
}

func (provider *Provider) GetNumTxnsDSEpoch() *jsonrpc.RPCResponse {
	return provider.call("GetNumTxnsDSEpoch")
}

func (provider *Provider) GetMinimumGasPrice() *jsonrpc.RPCResponse {
	return provider.call("GetMinimumGasPrice")
}

func (provider *Provider) GetSmartContractCode() *jsonrpc.RPCResponse {
	return provider.call("GetSmartContractCode")
}

func (provider *Provider) GetSmartContractInit() *jsonrpc.RPCResponse {
	return provider.call("GetSmartContractInit")
}

func (provider *Provider) GetSmartContractState() *jsonrpc.RPCResponse {
	return provider.call("GetSmartContractState")
}

func (provider *Provider) GetSmartContracts() *jsonrpc.RPCResponse {
	return provider.call("GetSmartContracts")
}

func (provider *Provider) GetContractAddressFromTransactionID() *jsonrpc.RPCResponse {
	return provider.call("GetContractAddressFromTransactionID")
}

func (provider *Provider) GetBalance() *jsonrpc.RPCResponse {
	return provider.call("GetBalance")
}

func (provider *Provider) call(method_name string) *jsonrpc.RPCResponse {
	response, err := provider.rpcClient.Call(method_name)
	if err != nil {
		return nil
	} else {
		return response
	}
}
