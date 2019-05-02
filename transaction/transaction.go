package transaction

import (
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/provider"
	"strconv"
	"strings"
)

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

func (t *Transaction) toTransactionParam() TxParams {
	param := TxParams{
		ID:           t.ID,
		Version:      t.Version,
		Nonce:        t.Nonce,
		Amount:       t.Amount,
		GasPrice:     t.GasPrice,
		GasLimit:     t.GasLimit,
		Signature:    t.Signature,
		Receipt:      t.Receipt,
		SenderPubKey: t.SenderPubKey,
		Code:         t.Code,
		Data:         t.Data,
	}

	if t.ToAddr == "" {
		param.ToAddr = "0000000000000000000000000000000000000000"
	} else {
		param.ToAddr = t.ToAddr
	}
	return param
}

func (t *Transaction) toTransactionPayload() TransactionPayload {
	version, _ := strconv.ParseInt(t.Version, 10, 32)
	nonce, _ := strconv.ParseInt(t.Nonce, 10, 32)

	return TransactionPayload{
		Version:   int(version),
		Nonce:     int(nonce),
		ToAddr:    LaksaGo.ToCheckSumAddress(t.ToAddr)[2:],
		Amount:    t.Amount,
		PubKey:    strings.ToLower(t.SenderPubKey),
		GasPrice:  t.GasPrice,
		GasLimit:  t.GasLimit,
		Code:      t.Code,
		Data:      t.Data,
		Signature: strings.ToLower(t.Signature),
	}
}

func (t *Transaction) bytes() ([]byte, error) {
	txParams := t.toTransactionParam()
	bytes, err := EncodeTransactionProto(txParams)
	if err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func (t *Transaction) isPending() bool {
	return t.Status == Pending
}

func (t *Transaction) isInitialised() bool {
	return t.Status == Initialised
}

func (t *Transaction) isConfirmed() bool {
	return t.Status == Confirmed
}

func (t *Transaction) isRejected() bool {
	return t.Status == Rejected
}
