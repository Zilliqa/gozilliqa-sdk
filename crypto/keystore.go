package crypto

import (
	"errors"
	"github.com/FireStack-Lab/LaksaGo/keytools"
)

// 0: p 1:s
type KDFType int

type Keystore struct {
	pbkdf2 *pbkdf2Wapper
	scrypt *scryptWapper
}

func NewDefaultKeystore() *Keystore {
	p := NewPbkdf2()
	s := NewScryptWapper()
	return &Keystore{
		pbkdf2: p,
		scrypt: s,
	}
}

func NewKeystore(p *pbkdf2Wapper, s *scryptWapper) *Keystore {
	return &Keystore{
		pbkdf2: p,
		scrypt: s,
	}
}

func (ks *Keystore) GetDerivedKey(password []byte, params interface{}) ([]byte, error) {
	p, ok := params.(Pbkdf2Params)
	if ok {
		return ks.pbkdf2.GetDerivedKey(password, p.Salt, p.Count, p.DkLen), nil
	}

	s, ok := params.(ScryptParams)
	if ok {
		r, err := ks.scrypt.GetDerivedKey(password, s.Salt, s.N, s.R, s.P, s.DkLen)
		if err != nil {
			return nil, err
		} else {
			return r, nil
		}
	}

	return nil, errors.New("unsupport params")
}

func (ks *Keystore) EncryptPrivateKey(privateKey, passphrase []byte, t KDFType) (string, error) {
	iv, err := keytools.GenerateRandomBytes(16)
	if err != nil {
		return "", err
	}
	salt, err := keytools.GenerateRandomBytes(32)
	if err != nil {
		return "", err
	}

	derivedKey := make([]byte, 32)
	err = nil

	if t == 0 {
		derivedKey = ks.pbkdf2.GetDerivedKey([]byte(passphrase), salt, 262144, 32)
	} else {
		derivedKey, err = ks.scrypt.GetDerivedKey([]byte(passphrase), salt, 8192, 8, 1, 32)
	}
	if err != nil {
		return "", nil
	}

	encryptKey := make([]byte, 16)
	copy(encryptKey, derivedKey[0:16])

	//perform cipher operation

	//build struct
}
