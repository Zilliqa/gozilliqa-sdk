package account

import (
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/crypto"
	"github.com/FireStack-Lab/LaksaGo/keytools"
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

func FromFile(file, passphrase string) (*Account, error) {
	ks := crypto.NewDefaultKeystore()
	privateKey, err := ks.DecryptPrivateKey(file, passphrase)
	if err != nil {
		return nil, err
	}
	return NewAccount(LaksaGo.DecodeHex(privateKey)), nil
}

func ToFile(privateKey, passphrase string, t crypto.KDFType) (string, error) {
	ks := crypto.NewDefaultKeystore()
	file, err := ks.EncryptPrivateKey(LaksaGo.DecodeHex(privateKey), LaksaGo.DecodeHex(passphrase), t)
	if err != nil {
		return "", nil
	}

	return file, nil
}
