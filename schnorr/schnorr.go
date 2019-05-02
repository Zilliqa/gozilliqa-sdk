package go_schnorr

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	"github.com/btcsuite/btcd/btcec"
	"math/big"
)



func TrySign(privateKey []byte, publicKey []byte, message []byte, k []byte) ([]byte, []byte, error) {

	priKey := new(big.Int).SetBytes(privateKey)

	// 1a. check if private key is 0
	if priKey.Cmp(new(big.Int).SetInt64(0)) <= 0 {
		return nil, nil, errors.New("private key must be > 0")
	}

	// 1b. check if private key is less than curve order, i.e., within [1...n-1]
	if priKey.Cmp(keytools.Secp256k1.N) >= 0 {
		return nil, nil, errors.New("private key cannot be greater than curve order")
	}

	// 2. Compute commitment Q = kG, where G is the base point
	Qx, Qy := keytools.Secp256k1.ScalarBaseMult(k)

	Q := LaksaGo.Compress(keytools.Secp256k1, Qx, Qy,true)

	// 3. Compute the challenge r = H(Q || pubKey || msg)
	// mod reduce r by the order of secp256k1, n
	r := new(big.Int).SetBytes(LaksaGo.Hash(Q, publicKey, message[:]))
	r = r.Mod(r, keytools.Secp256k1.N)

	if r.Cmp(new(big.Int).SetInt64(0)) == 0 {
		return nil, nil, errors.New("invalid r")
	}

	//4. Compute s = k - r * prv
	// 4a. Compute r * prv
	_r := *r
	s := new(big.Int).Mod(_r.Sub(new(big.Int).SetBytes(k), _r.Mul(&_r, priKey)), keytools.Secp256k1.N)

	if s.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("invalid s")
	}

	return r.Bytes(), s.Bytes(), nil
}

func Verify(publicKey []byte, msg []byte, r []byte, s []byte) bool {

	//cannot be zero
	if new(big.Int).SetBytes(r).Cmp(new(big.Int).SetInt64(0)) == 0 || new(big.Int).SetBytes(s).Cmp(new(big.Int).SetInt64(0)) == 0 {
		fmt.Printf("Invalid R or S value: cannot be zero")
		return false
	}

	//cannot be negative
	if new(big.Int).SetBytes(r).Sign() == -1 || new(big.Int).SetBytes(s).Sign() == -1 {
		fmt.Printf("Invalid R or S value: cannot be negative")
		return false
	}

	puk, err := btcec.ParsePubKey(publicKey, keytools.Secp256k1)

	if err != nil {
		panic("parse public key error")
	}

	pkx, pky := puk.X, puk.Y

	lx, ly := keytools.Secp256k1.ScalarMult(pkx, pky, r)
	rx, ry := keytools.Secp256k1.ScalarBaseMult(s)
	Qx, Qy := keytools.Secp256k1.Add(rx, ry, lx, ly)
	Q := LaksaGo.Compress(keytools.Secp256k1, Qx, Qy,true)

	_r := LaksaGo.Hash(Q, publicKey, msg)

	rn := new(big.Int).SetBytes(r)
	_rn := new(big.Int).SetBytes(_r)
	fmt.Printf("r = %s, _r = %s\n", hex.EncodeToString(r), hex.EncodeToString(_r))
	return rn.Cmp(_rn) == 0
}
