package provider

import (
	"encoding/json"
	"fmt"
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
