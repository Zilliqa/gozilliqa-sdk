package account

import (
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/crypto"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type Account struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    string
}

func NewAccount(privateKey []byte) *Account {
	publicKey := keytools.GetPublicKeyFromPrivateKey(privateKey, true)
	address := keytools.GetAddressFromPublic(publicKey)
	return &Account{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
	}
}

func NewHDAccountWithDerivationPath(mnemonic, path string) (*Account, error) {
	derivationPath, err := ParseDerivationPath(path)
	if err != nil {
		return nil, err
	}
	return newHDAccount(mnemonic, derivationPath)
}

func newHDAccount(mnemonic string, path DerivationPath) (*Account, error) {
	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	acc44H, err := masterKey.Child(hdkeychain.HardenedKeyStart + canonical(path[0]))
	if err != nil {
		return nil, err
	}
	acc44H313H, err := acc44H.Child(hdkeychain.HardenedKeyStart + canonical(path[1]))
	if err != nil {
		return nil, err
	}
	acc44H313H0H, err := acc44H313H.Child(hdkeychain.HardenedKeyStart + canonical(path[2]))
	if err != nil {
		return nil, err
	}
	acc44H313H0H0, err := acc44H313H0H.Child(canonical(path[3]))
	if err != nil {
		return nil, err
	}
	acc44H60H0H00, err := acc44H313H0H0.Child(canonical(path[4]))
	if err != nil {
		return nil, err
	}
	btcecPrivKey, err := acc44H60H0H00.ECPrivKey()
	if err != nil {
		return nil, err
	}
	privateKey := btcecPrivKey.ToECDSA()
	account := NewAccount(privateKey.D.Bytes())
	return account, nil
}

func NewDefaultHDAccount(mnemonic string, index uint32) (*Account, error) {
	path := fmt.Sprintf("m/44'/313'/0'/0/%d", index)
	derivationPath, err := ParseDerivationPath(path)
	if err != nil {
		return nil, err
	}
	return newHDAccount(mnemonic, derivationPath)
}

func FromFile(file, passphrase string) (*Account, error) {
	ks := crypto.NewDefaultKeystore()
	privateKey, err := ks.DecryptPrivateKey(file, passphrase)
	if err != nil {
		return nil, err
	}
	return NewAccount(util.DecodeHex(privateKey)), nil
}

func ToFile(privateKey, passphrase string, t crypto.KDFType) (string, error) {
	ks := crypto.NewDefaultKeystore()
	file, err := ks.EncryptPrivateKey(util.DecodeHex(privateKey), []byte(passphrase), t)
	if err != nil {
		return "", nil
	}

	return file, nil
}

func canonical(component uint32) uint32 {
	if component >= 0x80000000 {
		component -= 0x80000000
	}
	return component
}
