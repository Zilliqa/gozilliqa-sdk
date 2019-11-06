package account

import (
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

func NewHDAccount(mnemonic string, index uint32) (*Account, error) {
	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	acc44H, err := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return nil, err
	}
	acc44H313H, err := acc44H.Child(hdkeychain.HardenedKeyStart + 313)
	if err != nil {
		return nil, err
	}
	acc44H313H0H, err := acc44H313H.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		return nil, err
	}
	acc44H313H0H0, err := acc44H313H0H.Child(0)
	if err != nil {
		return nil, err
	}
	acc44H60H0H00, err := acc44H313H0H0.Child(index)
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
