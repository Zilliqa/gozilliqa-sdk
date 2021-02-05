/*
 * Copyright (C) 2021 Zilliqa
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
package core

type State int

const (
	Initialised State = iota
	Pending
	Confirmed
	Rejected
)

type AccountBase struct {
	Version     uint32
	Balance     uint64
	Nonce       uint64
	StorageRoot []byte
	CodeHash    []byte
}

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

type EventLog struct {
	EventName string          `json:"_eventname"`
	Address   string          `json:"address"`
	Params    []ContractValue `json:"params"`
}

type TransactionReceipt struct {
	Accept        bool                   `json:"accept"`
	Errors        interface{}            `json:"errors"`
	Exceptions    []TransactionException `json:"exceptions"`
	Success       bool                   `json:"success"`
	CumulativeGas string                 `json:"cumulative_gas"`
	EpochNum      string                 `json:"epoch_num"`
	EventLogs     []EventLog             `json:"event_logs"`
	Transitions   []Transition           `json:"transitions"`
}

type TransactionException struct {
	Line    int    `json:"line"`
	Message string `json:"message"`
}

type Transition struct {
	Accept bool               `json:"accept"`
	Addr   string             `json:"addr"`
	Depth  int                `json:"depth"`
	Msg    TransactionMessage `json:"msg"`
}

type TransactionMessage struct {
	Amount    string          `json:"_amount"`
	Recipient string          `json:"_recipient"`
	Tag       string          `json:"_tag"`
	Params    []ContractValue `json:"params"`
}

type Transactions struct {
	TxnHashes []string
}

type ContractValue struct {
	VName string      `json:"vname"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type ParamConstructor struct {
	Constructor string        `json:"constructor"`
	ArgTypes    []interface{} `json:"argtypes"`
	Arguments   []interface{} `json:"arguments"`
}

type BalanceAndNonce struct {
	Balance string `json:"balance"`
	Nonce   int64  `json:"nonce"`
}

var PendingTxnError = map[int]string{
	0:  "Txn was already processed and confirmed",
	1:  "Pending - nonce too high",
	2:  "Pending - blk gas limit exceeded",
	3:  "Pending - consensus failure",
	4:  "Error - txn not found",
	10: "Dropped - math error",
	11: "Dropped - scilla invocation error",
	12: "Dropped - account init error",
	13: "Dropped - invalid source account",
	14: "Dropped - gas limit too high",
	15: "Dropped - txn type unknown",
	16: "Dropped - txn in wrong shard",
	17: "Dropped - account in wrong shard",
	18: "Dropped - code size too large",
	19: "Dropped - txn verification error",
	20: "Dropped - gas limit too low",
	21: "Dropped - insuff balance",
	22: "Dropped - insuff gas for checker",
	23: "Dropped - duplicate txn found",
	24: "Dropped - txn w/ higher gas found",
	25: "Dropped - invalid dest account",
	26: "Dropped - state addition error",
}

type PendingTxnResult struct {
	Code      int  `json:"code"`
	Confirmed bool `json:"confirmed"`
	Info      string
}

type TransactionStatus struct {
	Code    int    `json:"code"`
	TxnHash string `json:"TxnHash"`
	Info    string
}

type PendingTxns struct {
	Txns []*TransactionStatus
}
