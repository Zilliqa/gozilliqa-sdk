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

func (provider *Provider) GetBlockchainInfo() *jsonrpc.RPCResponse {
	return provider.call("GetBlockchainInfo")
}

func (provider *Provider) call(method_name string) *jsonrpc.RPCResponse {
	response, err := provider.rpcClient.Call(method_name)
	if err != nil {
		return nil
	} else {
		return response
	}
}
