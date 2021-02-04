package core

import "math/big"

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

type CoSignatures struct {
	CS1 Signature
	B1  []bool
	CS2 Signature
	B2  []bool
}
