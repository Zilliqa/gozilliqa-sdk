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
package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
	"github.com/Zilliqa/gozilliqa-sdk/crypto"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	go_schnorr "github.com/Zilliqa/gozilliqa-sdk/schnorr"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/Zilliqa/gozilliqa-sdk/validator"
	"strconv"
	"strings"
)

const signatureSize = 128

type Wallet struct {
	Accounts       map[string]*Account
	DefaultAccount *Account
}

func NewWallet() *Wallet {
	accounts := make(map[string]*Account)
	return &Wallet{
		Accounts: accounts,
	}
}

func (w *Wallet) Sign(tx *transaction.Transaction, provider provider.Provider) error {
	if strings.HasPrefix(tx.ToAddr, "0x") {
		tx.ToAddr = strings.TrimPrefix(tx.ToAddr, "0x")
	}

	if !validator.IsBech32(tx.ToAddr) && !validator.IsChecksumAddress("0x" + tx.ToAddr) {
		return errors.New("not checksum Address or bech32")
	}

	if validator.IsBech32(tx.ToAddr) {
		adddress, err := bech32.FromBech32Addr(tx.ToAddr)
		if err != nil {
			return err
		}
		tx.ToAddr = adddress
	}

	if validator.IsChecksumAddress("0x" + tx.ToAddr) {
		tx.ToAddr = "0x" + tx.ToAddr
	}

	if tx.SenderPubKey != "" {
		address := keytools.GetAddressFromPublic(util.DecodeHex(tx.SenderPubKey))
		err := w.SignWith(tx, address, provider)
		if err != nil {
			return err
		}
		return nil
	}

	if w.DefaultAccount == nil {
		return errors.New("this wallet has no default account")
	}

	err2 := w.SignWith(tx, w.DefaultAccount.Address, provider)
	if err2 != nil {
		return err2
	}

	return nil

}

func (w *Wallet) SignWith(tx *transaction.Transaction, signer string, provider provider.Provider) error {
	account, ok := w.Accounts[strings.ToUpper(signer)]
	if !ok {
		return errors.New("account does not exist")
	}

	if tx.Nonce == "" {
		response,err := provider.GetBalance(signer)
		if err != nil {
			return err
		}
		if response == nil {
			return errors.New("get balance response err")
		}
		if response.Error == nil {
			result := response.Result.(map[string]interface{})
			n := result["nonce"].(json.Number)
			nonce, _ := n.Int64()
			tx.Nonce = strconv.FormatInt(nonce+1, 10)
		} else {
			tx.Nonce = "1"
		}
	}

	tx.SenderPubKey = util.EncodeHex(account.PublicKey)

	message, err := tx.Bytes()

	if err != nil {
		return err
	}

	rb, err2 := keytools.GenerateRandomBytes(keytools.Secp256k1.N.BitLen() / 8)

	if err2 != nil {
		return err2
	}

	var signature string

	for {
		r, s, err3 := go_schnorr.TrySign(account.PrivateKey, account.PublicKey, message, rb)
		if err3 != nil {
			return err3
		}
		sig := fmt.Sprintf("%s%s", util.EncodeHex(r), util.EncodeHex(s))
		if len(sig) == signatureSize {
			signature = sig
			break
		}
	}

	tx.Signature = signature

	return nil
}

func (w *Wallet) CreateAccount() {
	privateKey, _ := keytools.GeneratePrivateKey()
	account := NewAccount(privateKey[:])

	address := strings.ToUpper(keytools.GetAddressFromPrivateKey(privateKey[:]))
	w.Accounts[address] = account

	if w.DefaultAccount == nil {
		w.DefaultAccount = account
	}
}

func (w *Wallet) AddByPrivateKey(privateKey string) {
	prik := util.DecodeHex(privateKey)
	account := NewAccount(prik[:])
	address := strings.ToUpper(keytools.GetAddressFromPrivateKey(prik[:]))
	w.Accounts[address] = account

	if w.DefaultAccount == nil {
		w.DefaultAccount = account
	}
}

func (w *Wallet) AddByKeyStore(keystore, passphrase string) {
	ks := crypto.NewDefaultKeystore()
	privateKey, _ := ks.DecryptPrivateKey(keystore, passphrase)
	w.AddByPrivateKey(privateKey)
}

func (w *Wallet) SetDefault(address string) {
	account, ok := w.Accounts[strings.ToUpper(address)]
	if ok {
		w.DefaultAccount = account
	}
}
