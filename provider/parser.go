package provider

import (
	"encoding/json"
	"fmt"
	"github.com/ybbus/jsonrpc"
	"strconv"
)

func ParseTxBlock(rpcResult *jsonrpc.RPCResponse) (*TxBlock, error) {
	jsonResult, err := json.Marshal(rpcResult.Result)
	if err != nil {
		return nil, fmt.Errorf("ParseTxBlock: marshal rpc result, %s", err)
	}
	block := &TxBlock{}
	if err := json.Unmarshal(jsonResult, block); err != nil {
		return nil, fmt.Errorf("ParseTxBlock: unmarshal txBlock, %s", err)
	}
	return block, nil
}

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

func ParseBlockHeight(rpcResult *jsonrpc.RPCResponse) (uint64, error) {
	if heightString, ok := rpcResult.Result.(string); ok {
		height, err := strconv.ParseUint(heightString, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("ParseBlockHeight: result %s invalid, %s", heightString, err)
		}
		return height, nil
	} else {
		return 0, fmt.Errorf("ParseBlockHeight: type unmatch")
	}
}
