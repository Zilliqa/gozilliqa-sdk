package account

import "github.com/FireStack-Lab/LaksaGo/keytools"

type Account struct {
	privateKey []byte
	publicKey  []byte
	address    string
}

func NewAccount(privateKey []byte) Account {
	publicKey := keytools.GetPublicKeyFromPrivateKey(privateKey, true)
	address := keytools.GetAddressFromPublic(publicKey)
	return Account{
		privateKey: privateKey,
		publicKey:  publicKey,
		address:    address,
	}
}
