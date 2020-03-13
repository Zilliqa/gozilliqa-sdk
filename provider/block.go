package provider

import (
	"encoding/json"
	"fmt"
	"github.com/ybbus/jsonrpc"
)

type TxBlock struct {
	Header *Header `json:"header"`
	Body   *Body   `json:"body"`
}

type Header struct {
	BlockNum       string
	DSBlockNum     string
	GasLimit       string
	GasUsed        string
	MbInfoHash     string
	MinerPubKey    string
	NumMicroBlocks uint32
	NumTxns        uint32
	PrevBlockHash  string
	Rewards        string
	StateDeltaHash string
	StateRootHash  string
	Timestamp      string
	Version        uint32
}

type Body struct {
	BlockHash       string
	HeaderSign      string
	MicroBlockInfos []*MicroBlockInfo
}

type MicroBlockInfo struct {
	MicroBlockHash        string
	MicroBlockShardId     uint32
	MicroBlockTxnRootHash string
}

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
