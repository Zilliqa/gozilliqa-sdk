package core

import (
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) Serialize() []byte {
	bns := BIGNumSerialize{}
	data := make([]byte, 0)
	data = bns.SetNumber(data, 0, signatureChallengeSize, s.R)
	data = bns.SetNumber(data, signatureChallengeSize, signatureChallengeSize, s.S)
	return data
}

func NewFromByteArray(bytes []byte) *Signature {
	rb := make([]byte,32)
	sb := make([]byte,32)
	copy(rb,bytes[0:32])
	copy(sb,bytes[32:])

	r := new(big.Int).SetBytes(rb)
	s := new(big.Int).SetBytes(sb)

	return &Signature{
		R: r,
		S: s,
	}
}

type CoSignatures struct {
	CS1 *Signature
	B1  []bool
	CS2 *Signature
	B2  []bool
}
