package contract

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	provider2 "github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/util"
)

func TestContract_Deploy(t *testing.T) {
	host := "https://dev-api.zilliqa.com/"
	privateKey := "e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930"
	chainID := 333
	msgVersion := 1

	publickKey := keytools.GetPublicKeyFromPrivateKey(util.DecodeHex(privateKey), true)
	address := keytools.GetAddressFromPublic(publickKey)
	pubkey := util.EncodeHex(publickKey)
	provider := provider2.NewProvider(host)

	wallet := account.NewWallet()
	wallet.AddByPrivateKey(privateKey)

	code, _ := ioutil.ReadFile("./fungible.scilla")
	init := []Value{
		{
			"_scilla_version",
			"Uint32",
			"0",
		},
		{
			"owner",
			"ByStr20",
			"0x" + address,
		},
		{
			"total_tokens",
			"Uint128",
			"1000000000",
		},
		{
			"decimals",
			"Uint32",
			"0",
		},
		{
			"name",
			"String",
			"BobCoin",
		},
		{
			"symbol",
			"String",
			"BOB",
		},
	}
	contract := Contract{
		Code:     string(code),
		Init:     init,
		Singer:   wallet,
		Provider: provider,
	}

	nonce, _ := provider.GetBalance(address).Result.(map[string]interface{})["nonce"].(json.Number).Int64()

	deployParams := DeployParams{
		Version:      strconv.FormatInt(int64(util.Pack(chainID, msgVersion)), 10),
		Nonce:        strconv.FormatInt(nonce+1, 10),
		GasPrice:     "10000000000",
		GasLimit:     "10000",
		SenderPubKey: pubkey,
	}

	tx, err := contract.Deploy(deployParams)

	if err != nil {
		panic(err.Error())
	}

	tx.Confirm(tx.ID, 1000, 10, provider)
}

func TestContract_Call(t *testing.T) {
	host := "https://dev-api.zilliqa.com/"
	privateKey := "e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930"
	chainID := 333
	msgVersion := 1

	publickKey := keytools.GetPublicKeyFromPrivateKey(util.DecodeHex(privateKey), true)
	address := keytools.GetAddressFromPublic(publickKey)
	pubkey := util.EncodeHex(publickKey)
	provider := provider2.NewProvider(host)

	wallet := account.NewWallet()
	wallet.AddByPrivateKey(privateKey)

	contract := Contract{
		Address:  "bd7198209529dC42320db4bC8508880BcD22a9f2",
		Singer:   wallet,
		Provider: provider,
	}

	args := []Value{
		{
			"to",
			"ByStr20",
			"0x" + address,
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
		Version:      strconv.FormatInt(int64(util.Pack(chainID, msgVersion)), 10),
		GasPrice:     "1000000000",
		GasLimit:     "1000",
		SenderPubKey: pubkey,
		Amount:       "0",
	}

	err, tx := contract.Call("Transfer", args, params, true, 1000, 3)
	if err != nil {
		fmt.Printf(err.Error())
	}

	tx.Confirm(tx.ID, 1000, 3, provider)

}
