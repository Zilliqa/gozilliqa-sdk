package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strconv"
	"strings"
	"time"
)

type Transaction core.Transaction

func NewFromPayload(payload *provider.TransactionPayload) *Transaction {
	v := strconv.FormatInt(int64(payload.Version), 10)
	n := strconv.FormatInt(int64(payload.Nonce), 10)
	var toAddr string
	if payload.ToAddr == "0x0000000000000000000000000000000000000000" {
		toAddr = "0x0000000000000000000000000000000000000000"
	} else {
		toAddr = "0x" + payload.ToAddr
	}
	if payload.Data == "" {
		// payment
		return &Transaction{
			ID:              "",
			Version:         v,
			Nonce:           n,
			Amount:          payload.Amount,
			GasPrice:        payload.GasPrice,
			GasLimit:        payload.GasLimit,
			Signature:       payload.Signature,
			Receipt:         core.TransactionReceipt{},
			SenderPubKey:    payload.PubKey,
			ToAddr:          toAddr,
			Code:            payload.Code,
			Data:            "",
			Status:          0,
			ContractAddress: "",
			Priority:        payload.Priority,
		}
	} else if strings.Contains(payload.Data, "_tag") {
		// contract call
		var data provider.Data
		json.Unmarshal([]byte(payload.Data), &data)
		return &Transaction{
			ID:              "",
			Version:         v,
			Nonce:           n,
			Amount:          payload.Amount,
			GasPrice:        payload.GasPrice,
			GasLimit:        payload.GasLimit,
			Signature:       payload.Signature,
			Receipt:         core.TransactionReceipt{},
			SenderPubKey:    payload.PubKey,
			ToAddr:          toAddr,
			Code:            payload.Code,
			Data:            data,
			Status:          0,
			ContractAddress: "",
			Priority:        payload.Priority,
		}
	} else {
		// contract deployment
		var data []provider.Value
		json.Unmarshal([]byte(payload.Data), &data)
		return &Transaction{
			ID:              "",
			Version:         v,
			Nonce:           n,
			Amount:          payload.Amount,
			GasPrice:        payload.GasPrice,
			GasLimit:        payload.GasLimit,
			Signature:       payload.Signature,
			Receipt:         core.TransactionReceipt{},
			SenderPubKey:    payload.PubKey,
			ToAddr:          toAddr,
			Code:            payload.Code,
			Data:            data,
			Status:          0,
			ContractAddress: "",
			Priority:        payload.Priority,
		}
	}

}

func (t *Transaction) toTransactionParam() TxParams {
	data, _ := json.Marshal(t.Data)
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
		Data:         string(data),
	}

	if t.ToAddr == "" {
		param.ToAddr = "0000000000000000000000000000000000000000"
	} else {
		param.ToAddr = t.ToAddr
	}
	return param
}

func (t *Transaction) ToTransactionPayload() provider.TransactionPayload {
	version, _ := strconv.ParseInt(t.Version, 10, 32)
	nonce, _ := strconv.ParseInt(t.Nonce, 10, 32)
	data, _ := json.Marshal(t.Data)

	p := provider.TransactionPayload{
		Version:   int(version),
		Nonce:     int(nonce),
		ToAddr:    util.ToCheckSumAddress(t.ToAddr)[2:],
		Amount:    t.Amount,
		PubKey:    strings.ToLower(t.SenderPubKey),
		GasPrice:  t.GasPrice,
		GasLimit:  t.GasLimit,
		Code:      t.Code,
		Signature: strings.ToLower(t.Signature),
		Priority:  t.Priority,
	}

	if string(data) != "\"\"" {
		p.Data = string(data)
	}

	if p.ToAddr == "0000000000000000000000000000000000000000" {
		p.ToAddr = "0x0000000000000000000000000000000000000000"
	}
	return p
}

func (t *Transaction) TrackTx(hash string, provider *provider.Provider) bool {
	txn, err := provider.GetTransaction(hash)
	if err != nil {
		fmt.Println("Track error: " + err.Error())
		return false
	}
	t.ID = txn.ID
	t.Receipt = txn.Receipt

	if !t.Receipt.Success {
		t.Status = core.Rejected
	} else {
		t.Status = core.Confirmed
	}
	return true
}

func (t *Transaction) Confirm(hash string, maxAttempts, interval int, provider *provider.Provider) {
	t.Status = core.Pending
	for i := 0; i < maxAttempts; i++ {
		fmt.Println("track " + hash)
		tracked := t.TrackTx(hash, provider)
		time.Sleep(time.Duration(interval) * time.Second)
		if tracked {
			fmt.Println("confirmed! " + hash)
			return
		}
	}
	t.Status = core.Rejected
}

func (t *Transaction) Bytes() ([]byte, error) {
	txParams := t.toTransactionParam()
	bytes, err := EncodeTransactionProto(txParams)

	if err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func (t *Transaction) Hash() ([]byte, error) {
	bytes, err := t.Bytes()
	if err != nil {
		return nil, err
	}
	hash := util.Sha256(bytes)
	return hash, nil
}

func (t *Transaction) isPending() bool {
	return t.Status == core.Pending
}

func (t *Transaction) isInitialised() bool {
	return t.Status == core.Initialised
}

func (t *Transaction) isConfirmed() bool {
	return t.Status == core.Confirmed
}

func (t *Transaction) isRejected() bool {
	return t.Status == core.Rejected
}
