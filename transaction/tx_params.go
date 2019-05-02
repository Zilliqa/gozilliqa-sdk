package transaction

type TxParams struct {
	ID           string `json:"ID"`
	Version      string `json:"version"`
	Nonce        string `json:"nonce"`
	Amount       string `json:"amount"`
	GasPrice     string `json:"gasPrice"`
	GasLimit     string `json:"gasLimit"`
	Signature    string `json:"signature"`
	Receipt      TransactionReceipt `json:"receipt"`
	SenderPubKey string `json:"senderPubKey"`
	ToAddr       string `json:"toAddr"`
	Code         string `json:"code"`
	Data         string `json:"data"`
}
