package account

import (
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/crypto"
	"github.com/FireStack-Lab/LaksaGo/keytools"
)

type Wallet struct {
	Accounts       map[string]*Account
	DefaultAccount *Account
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
