/*
 * Copyright (C) 2019 Zilliqa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
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
	host := "http://kaya.godscilla.com:5555"
	privateKey := "4d5296a36eb3d95e6ea576e1ef5eeda69be486eaf381fd61d5f5b3d1c5e3e854"
	chainID := 1
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
		Signer:   wallet,
		Provider: provider,
	}

	nonce, _ := provider.GetBalance(address).Result.(map[string]interface{})["nonce"].(json.Number).Int64()

	gasPrice := provider.GetMinimumGasPrice().Result.(string)

	deployParams := DeployParams{
		Version:      strconv.FormatInt(int64(util.Pack(chainID, msgVersion)), 10),
		Nonce:        strconv.FormatInt(nonce+1, 10),
		GasPrice:     gasPrice,
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
		Signer:   wallet,
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
	gasPrice := provider.GetMinimumGasPrice().Result.(string)


	params := CallParams{
		Nonce:        strconv.FormatInt(n, 10),
		Version:      strconv.FormatInt(int64(util.Pack(chainID, msgVersion)), 10),
		GasPrice:     gasPrice,
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

func TestContract_Sign(t *testing.T) {
	host := "https://dev-api.zilliqa.com/"
	privateKey := "e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930"
	chainID := 333
	msgVersion := 1

	publickKey := keytools.GetPublicKeyFromPrivateKey(util.DecodeHex(privateKey), true)
	//address := keytools.GetAddressFromPublic(publickKey)
	pubkey := util.EncodeHex(publickKey)
	provider := provider2.NewProvider(host)

	wallet := account.NewWallet()
	wallet.AddByPrivateKey(privateKey)

	contract := Contract{
		Address:  "84eb5C96Bec8d29eDdFBe36865E9B7F26b816f0F",
		Signer:   wallet,
		Provider: provider,
	}

	args := []Value{
		{
			"proxyTokenContract",
			"ByStr20",
			"0x39550ab45d74cce5fef70e857c1326b2d9bee096",
		},
		{
			"to",
			"ByStr20",
			"0x39550ab45d74cce5fef70e857c1326b2d9bee096",
		},
		{
			"value",
			"Uint128",
			"10000000",
		},
	}

	nonce, _ := provider.GetBalance("9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a").Result.(map[string]interface{})["nonce"].(json.Number).Int64()
	n := nonce + 1
	gasPrice := provider.GetMinimumGasPrice().Result.(string)


	params := CallParams{
		Nonce:        strconv.FormatInt(n, 10),
		Version:      strconv.FormatInt(int64(util.Pack(chainID, msgVersion)), 10),
		GasPrice:     gasPrice,
		GasLimit:     "1000",
		SenderPubKey: pubkey,
		Amount:       "0",
	}

	err, tx := contract.Sign("SubmitCustomMintTransaction", args, params, true)
	if err != nil {
		fmt.Printf(err.Error())
	}

	pl :=  tx.ToTransactionPayload()
	j,_ := pl.ToJson()
	fmt.Println(string(j))

	payload2, err3 := provider2.NewFromJson(j)
	if err3 != nil {
		t.Error(err3.Error())
	}
	rsp := provider.CreateTransaction(*payload2)
	fmt.Println(rsp.Error)
	fmt.Println(rsp.Result)


}
