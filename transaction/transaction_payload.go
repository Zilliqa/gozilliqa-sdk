package transaction

type TransactionPayload struct {
	Version   int
	Nonce     int
	ToAddr    string
	Amount    string
	PubKey    string
	GasPrice  string
	GasLimit  string
	Code      string
	Data      string
	Signature string
}
