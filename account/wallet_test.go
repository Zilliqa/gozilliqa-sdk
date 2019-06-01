package account

import (
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	provider2 "github.com/FireStack-Lab/LaksaGo/provider"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"strconv"
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
		ToAddr:       "bd7198209529dC42320db4bC8508880BcD22a9f2",
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

//{"version":21823489,"nonce":226,"toAddr":"bd7198209529dC42320db4bC8508880BcD22a9f2","amount":"0","pubKey":"0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a","gasPrice":"1000000000","gasLimit":"1000","code":"","data":"{\"_tag\":\"Transfer\",\"params\":[{\"vname\":\"to\",\"type\":\"ByStr20\",\"value\":\"0x381f4008505e940ad7681ec3468a719060caf796\"},{\"vname\":\"tokens\",\"type\":\"Uint128\",\"value\":\"10\"}]}","signature":"746fab24f8940fff741fe46dfb228c01e965c9c54b106959af60d2a7ddfbef49af0c2a8ee996155fc772fa4f90de2c10bd5e49aff29df534f7d162b2aaa80f61"}
//{"version":21823489,"nonce":225,"toAddr":"bd7198209529dC42320db4bC8508880BcD22a9f2","amount":"0","pubKey":"0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a","gasPrice":"1000000000","gasLimit":"1000","code":"","data":"{\"_tag\":\"Transfer\",\"params\":[{\"vname\":\"to\",\"type\":\"ByStr20\",\"value\":\"0x381f4008505e940ad7681ec3468a719060caf796\"},{\"vname\":\"tokens\",\"type\":\"Uint128\",\"value\":\"10\"}]}","signature":"e5cb1859bb712e961df8a8c2be84e8185ed9468562baf546d706f87e663093906fae8fd086df6b142c8348b2fe11ec0e55355fcc53f2de07a485f6cf720d1628"}