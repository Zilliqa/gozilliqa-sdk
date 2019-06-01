package transaction

import (
	"github.com/FireStack-Lab/LaksaGo/provider"
	"testing"
)

func TestTransaction_TrackTx(t *testing.T) {
	provider := provider.NewProvider("https://dev-api.zilliqa.com/")
	tx := Transaction{
		ID:           "",
		Version:      "",
		Nonce:        "",
		Amount:       "",
		GasPrice:     "",
		GasLimit:     "",
		Signature:    "",
		Receipt:      TransactionReceipt{
			Success:       false,
			CumulativeGas: "",
			EpochNum:      "",
		},
		SenderPubKey: "",
		ToAddr:       "",
		Code:         "",
		Data:         "",
		Status:       0,
	}

	tx.Confirm("846cda64971e259b1739bf15710758803abcf5754507af5af3f779777cd1b0b0",1000,3,provider)

}
