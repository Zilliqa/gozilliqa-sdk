package core

type TxBlockT struct {
	Header TxBlockHeaderT `json:"header"`
	Body   TxBlockBodyT   `json:"body"`
}

type TxBlockBodyT struct {
	BlockHash       string
	HeaderSign      string
	MicroBlockInfos []MicroBlockInfo
}
