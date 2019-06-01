package contract

import (
	"encoding/json"
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/account"
	provider2 "github.com/FireStack-Lab/LaksaGo/provider"
	"strconv"
	"testing"
)

func TestContract_Call(t *testing.T) {
	wallet := account.NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")

	contract := Contract{
		Address:  "bd7198209529dC42320db4bC8508880BcD22a9f2",
		singer:   wallet,
		provider: provider,
	}

	args := []Value{
		{
			"to",
			"ByStr20",
			"0x381f4008505e940ad7681ec3468a719060caf796",
		},
		{
			"tokens",
			"Uint128",
			"10",
		},
	}

	nonce, _ := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a").Result.(map[string]interface{})["nonce"].(json.Number).Int64()
	n := nonce + 1
	params := CallParams{
		Nonce:        strconv.FormatInt(n, 10),
		Version:      strconv.FormatInt(int64(LaksaGo.Pack(333, 1)), 10),
		GasPrice:     "1000000000",
		GasLimit:     "1000",
		SenderPubKey: "0246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A",
		Amount:       "0",
	}

	err, tx := contract.Call("Transfer", args, params, 1000, 3)
	if err != nil {
		fmt.Printf(err.Error())
	}

	tx.Confirm(tx.ID, 1000, 3, provider)

}
