package transaction

import "github.com/FireStack-Lab/LaksaGo/provider"

type State int

const (
	Initialised State = iota
	Pending
	Confirmed
	Rejected
)

type Transaction struct {
	ID           string
	Version      string
	Nonce        string
	Amount       string
	GasPrice     string
	GasLimit     string
	Signature    string
	Receipt      TransactionReceipt
	SenderPubKey string
	ToAddr       string
	Code         string
	Data         string

	Provider provider.Provider
	Status   State
}
