package provider

type TransactionPayload struct {
	Version   int    `json:"version"`
	Nonce     int    `json:"nonce"`
	ToAddr    string `json:"toAddr"`
	Amount    string `json:"amount"`
	PubKey    string `json:"pubKey"`
	GasPrice  string `json:"gasPrice"`
	GasLimit  string `json:"gasLimit"`
	Code      string `json:"code"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
	Priority  bool   `json:"priority"`
}
