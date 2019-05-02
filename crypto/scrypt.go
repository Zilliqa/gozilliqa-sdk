package crypto

import (
	"golang.org/x/crypto/scrypt"
)

type scryptWapper struct {
}

func NewScryptWapper() *scryptWapper {
	return &scryptWapper{}
}

func (s *scryptWapper) GetDerivedKey(password, salt []byte, n, r, p, dkLen int) ([]byte, error) {
	dk, err := scrypt.Key(password, salt, n, r, p, dkLen)
	if err != nil {
		return nil, err
	} else {
		return dk, nil
	}
}

type ScryptParams struct {
	Salt  []byte
	DkLen int
	N     int
	R     int
	P     int
}
