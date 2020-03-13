package provider

import (
	"encoding/json"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"github.com/ybbus/jsonrpc"
)

func ParseTxHashArray(rpcResult *jsonrpc.RPCResponse) ([][]string, error) {
	jsonResult, err := json.Marshal(rpcResult.Result)
	if err != nil {
		return nil, fmt.Errorf("ParseTxHashArray: marshal rpc result, %s", err)
	}
	result := [][]string{}
	if err := json.Unmarshal(jsonResult, &result); err != nil {
		return nil, fmt.Errorf("ParseTxHashArray: unmarshal result, %s", err)
	}
	return result, nil
}

func ParseTx(rpcResult *jsonrpc.RPCResponse) (*transaction.Transaction, error) {
	jsonResult, err := json.Marshal(rpcResult.Result)
	if err != nil {
		return nil, fmt.Errorf("ParseTx: marshal rpc result, %s", err)
	}
	result := &transaction.Transaction{}
	if err := json.Unmarshal(jsonResult, &result); err != nil {
		return nil, fmt.Errorf("ParseTx: unmarshal result, %s", err)
	}
	return result, nil
}
