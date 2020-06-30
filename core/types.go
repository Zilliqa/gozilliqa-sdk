package core

type State int

const (
	Initialised State = iota
	Pending
	Confirmed
	Rejected
)

type BlockchainInfo struct {
	CurrentDSEpoch    string
	CurrentMiniEpoch  string
	DSBlockRate       float64
	NumDSBlocks       string
	NumPeers          int
	NumTransactions   string
	NumTxBlocks       string
	NumTxnsDSEpoch    string
	NumTxnsTxEpoch    string
	ShardingStructure ShardingStructure
	TransactionRate   float64
	TxBlockRate       float64
}

type ShardingStructure struct {
	NumPeers []int
}

type DSBlock struct {
	Header    DsHeader `json:"header"`
	Signature string   `json:"signature"`
}

type DsHeader struct {
	BlockNum     string
	Difficulty   int
	DifficultyDS int
	GasPrice     string
	LeaderPubKey string
	PoWWinners   []interface{}
	PrevHash     string
	Timestamp    string
}

type HashAndNum struct {
	BlockNum int64
	Hash     string
}

type BlockList struct {
	Data     []HashAndNum `json:"data"`
	MaxPages int          `json:"maxPages"`
}

type TxBlock struct {
	Header TxBlockHeader `json:"header"`
	Body   TxBlockBody   `json:"body"`
}

type TxBlockHeader struct {
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

type TxBlockBody struct {
	BlockHash       string
	HeaderSign      string
	MicroBlockInfos []MicroBlockInfo
}

type MicroBlockInfo struct {
	MicroBlockHash        string
	MicroBlockShardId     int
	MicroBlockTxnRootHash string
}

type MinerInfo struct {
	DsCommittee []string    `json:"dscommittee"`
	Shards      []ShardInfo `json:"shards"`
}

type ShardInfo struct {
	Nodes []string `json:"nodes"`
	Size  int      `json:"size"`
}

type Transaction struct {
	ID              string
	Version         string
	Nonce           string
	Amount          string
	GasPrice        string
	GasLimit        string
	Signature       string
	Receipt         TransactionReceipt
	SenderPubKey    string
	ToAddr          string
	Code            string
	Data            interface{}
	Status          State
	ContractAddress string
	Priority        bool
}

type TransactionReceipt struct {
	Accept bool `json:"accept"`
	Errors interface{} `json:"errors"`
	Exceptions []TransactionException `json:"exceptions"`
	Success       bool          `json:"success"`
	CumulativeGas string        `json:"cumulative_gas"`
	EpochNum      string        `json:"epoch_num"`
	EventLogs     []interface{} `json:"event_logs"`
	Transitions []Transition `json:"transitions"`
}

type TransactionException struct {
	Line int `json:"line"`
	Message string `json:"message"`
}

type Transition struct {
	Accept bool `json:"accept"`
	Addr string `json:"addr"`
	Depth int `json:"depth"`
	Msg TransactionMessage `json:"msg"`
}

type TransactionMessage struct {
	Amount string `json:"_amount"`
	Receipt string `json:"_receipt"`
	Tag string `json:"_tag"`
	Params []ContractValue `json:"params"`
}

type Transactions struct {
	TxnHashes []string
}

type ContractValue struct {
	VName string      `json:"vname"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type BalanceAndNonce struct {
	Balance string `json:"balance"`
	Nonce   int64  `json:"nonce"`
}
