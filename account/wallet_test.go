package account

import (
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	provider2 "github.com/FireStack-Lab/LaksaGo/provider"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"strconv"
	"strings"
	"testing"
)

func TestWallet_SignWith(t *testing.T) {
	wallet := NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	tx := &transaction.Transaction{}
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")
	_ = wallet.SignWith(tx, "9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a", *provider)
}

func TestSendTransaction(t *testing.T) {
	wallet := NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")

	tx := &transaction.Transaction{
		Version:      strconv.FormatInt(int64(LaksaGo.Pack(333, 2)), 10),
		SenderPubKey: "0246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A",
		ToAddr:       strings.ToLower("0x4baf5fada8e5db92c3d3242618c5b47133ae003c"),
		Amount:       "10000000",
		GasPrice:     "1000000000",
		GasLimit:     "1",
		Code:         "",
		Data:         "",
	}

	err := wallet.Sign(tx, *provider)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	rsp := provider.CreateTransaction(tx.ToTransactionPayload())

	if rsp.Error != nil {
		fmt.Println(rsp.Error)
		t.Error(err)
	} else {
		result := rsp.Result.(map[string]interface{})
		hash := result["TranID"].(string)
		fmt.Printf("hash is %s",hash)
	}

}
