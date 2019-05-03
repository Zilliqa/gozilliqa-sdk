package account

import (
	"encoding/json"
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/crypto"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	"github.com/FireStack-Lab/LaksaGo/provider"
	go_schnorr "github.com/FireStack-Lab/LaksaGo/schnorr"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"strconv"
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

//todo unit test
func (w *Wallet) SignWith(tx *transaction.Transaction, signer string, provider provider.Provider) error {
	account, ok := w.Accounts[signer]
	if !ok {
		panic("account not exists")
	}

	if tx.Nonce == "" {
		response := provider.GetBalance(signer)
		if response.Error == nil {
			result := response.Result.(map[string]interface{})
			n := result["nonce"].(json.Number)
			nonce, _ := n.Int64()
			tx.Nonce = strconv.FormatInt(nonce+1, 10)
			println(tx.Nonce)
		}
	}

	tx.SenderPubKey = LaksaGo.EncodeHex(account.publicKey)

	message, err := tx.Bytes()

	if err != nil {
		return err
	}

	rb, err2 := keytools.GenerateRandomBytes(keytools.Secp256k1.N.BitLen() / 8)

	if err2 != nil {
		return err2
	}

	r, s, err3 := go_schnorr.TrySign(account.privateKey, account.publicKey, message, rb)

	if err3 != nil {
		return err3
	}

	signature := fmt.Sprintf("%s %s", LaksaGo.EncodeHex(r), LaksaGo.EncodeHex(s))

	tx.Signature = signature

	return nil
}

func (w *Wallet) CreateAccount() {
	privateKey, _ := keytools.GeneratePrivateKey()
	account := NewAccount(privateKey[:])

	address := keytools.GetAddressFromPrivateKey(privateKey[:])
	w.Accounts[address] = account

	if w.DefaultAccount == nil {
		w.DefaultAccount = account
	}
}

func (w *Wallet) AddByPrivateKey(privateKey string) {
	prik := LaksaGo.DecodeHex(privateKey)
	account := NewAccount(prik[:])
	address := keytools.GetAddressFromPrivateKey(prik[:])
	println(address)
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
	account, ok := w.Accounts[address]
	if ok {
		w.DefaultAccount = account
	}
}
