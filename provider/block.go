package provider

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
