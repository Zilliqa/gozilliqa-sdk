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
	tx.TrackTx("13d24a3d9137a7a047ca468511a3856d4a173777a0c35d78b44624e5fc0b91bc",provider)
	if tx.Status != Confirmed {
		t.Error("track tx failed")
	}

}
