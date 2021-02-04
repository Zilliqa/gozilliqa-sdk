package core

type TxBlockHeaderT struct {
	BlockNum       string
	DSBlockNum     string
	GasLimit       string
	GasUsed        string
	MbInfoHash     string
	MinerPubKey    string
	NumMicroBlocks int
	NumTxns        int
	PrevBlockHash  string
	Rewards        string
	StateDeltaHash string
	StateRootHash  string
	Timestamp      string
	Version        int
}
