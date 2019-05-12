package contract

import (
	"fmt"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"testing"
)

func TestGetAddressFromContract(t *testing.T) {
	tx := &transaction.Transaction{
		ID:           "",
		Version:      "",
		Nonce:        "19",
		Amount:       "",
		GasPrice:     "",
		GasLimit:     "",
		Signature:    "",
		Receipt:      transaction.TransactionReceipt{},
		SenderPubKey: "0246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A",
		ToAddr:       "",
		Code:         "",
		Data:         "",
		Status:       0,
	}

	address := GetAddressFromContract(tx)
	fmt.Printf(address)

	if address != "8f14cb1735b2b5fba397bea1c223d65d12b9a887" {
		t.Error("get address from contract tx failed")
	}
}
