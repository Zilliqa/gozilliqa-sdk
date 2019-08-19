package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/bech32"
	"github.com/FireStack-Lab/LaksaGo/crypto"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	"github.com/FireStack-Lab/LaksaGo/provider"
	go_schnorr "github.com/FireStack-Lab/LaksaGo/schnorr"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"github.com/FireStack-Lab/LaksaGo/validator"
	"strconv"
	"strings"
)

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
		address := keytools.GetAddressFromPublic(LaksaGo.DecodeHex(tx.SenderPubKey))
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
		response := provider.GetBalance(signer)
		if response.Error == nil {
			result := response.Result.(map[string]interface{})
			n := result["nonce"].(json.Number)
			nonce, _ := n.Int64()
			tx.Nonce = strconv.FormatInt(nonce+1, 10)
		} else {
			tx.Nonce = "1"
		}
	}

	tx.SenderPubKey = LaksaGo.EncodeHex(account.PublicKey)

	message, err := tx.Bytes()

	if err != nil {
		return err
	}

	rb, err2 := keytools.GenerateRandomBytes(keytools.Secp256k1.N.BitLen() / 8)

	if err2 != nil {
		return err2
	}

	r, s, err3 := go_schnorr.TrySign(account.PrivateKey, account.PublicKey, message, rb)

	if err3 != nil {
		return err3
	}

	signature := fmt.Sprintf("%s%s", LaksaGo.EncodeHex(r), LaksaGo.EncodeHex(s))

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
	prik := LaksaGo.DecodeHex(privateKey)
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
