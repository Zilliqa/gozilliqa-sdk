package go_schnorr

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	"github.com/btcsuite/btcd/btcec"
)

var bintZero = big.NewInt(0)

func TrySign(privateKey []byte, publicKey []byte, message []byte, k []byte) ([]byte, []byte, error) {
	priKey := new(big.Int).SetBytes(privateKey)
	bintK := new(big.Int).SetBytes(k)

	// 1a. check if private key is 0
	if priKey.Cmp(new(big.Int).SetInt64(0)) <= 0 {
		return nil, nil, errors.New("private key must be > 0")
	}

	// 1b. check if private key is less than curve order, i.e., within [1...n-1]
	if priKey.Cmp(keytools.Secp256k1.N) >= 0 {
		return nil, nil, errors.New("private key cannot be greater than curve order")
	}

	if bintK.Cmp(bintZero) == 0 {
		return nil, nil, errors.New("k cannot be zero")
	}

	if bintK.Cmp(keytools.Secp256k1.N) > 0 {
		return nil, nil, errors.New("k cannot be greater than order of secp256k1")
	}

	// 2. Compute commitment Q = kG, where G is the base point
	Qx, Qy := keytools.Secp256k1.ScalarBaseMult(k)

	Q := LaksaGo.Compress(keytools.Secp256k1, Qx, Qy, true)

	// 3. Compute the challenge r = H(Q || pubKey || msg)
	// mod reduce r by the order of secp256k1, n
	r := new(big.Int).SetBytes(LaksaGo.Hash(Q, publicKey, message[:]))
	r = r.Mod(r, keytools.Secp256k1.N)

	if r.Cmp(bintZero) == 0 {
		return nil, nil, errors.New("invalid r")
	}

	//4. Compute s = k - r * prv
	// 4a. Compute r * prv
	_r := *r
	s := new(big.Int).Mod(_r.Sub(bintK, _r.Mul(&_r, priKey)), keytools.Secp256k1.N)

	if s.Cmp(big.NewInt(0)) == 0 {
		return nil, nil, errors.New("invalid s")
	}

	return r.Bytes(), s.Bytes(), nil
}

func Verify(publicKey []byte, msg []byte, r []byte, s []byte) bool {
	bintR := new(big.Int).SetBytes(r)
	bintS := new(big.Int).SetBytes(s)

	//cannot be zero
	if bintR.Cmp(bintZero) == 0 || bintS.Cmp(bintZero) == 0 {
		fmt.Printf("Invalid R or S value: cannot be zero")
		return false
	}

	//cannot be negative
	if bintR.Sign() == -1 || bintS.Sign() == -1 {
		fmt.Printf("Invalid R or S value: cannot be negative")
		return false
	}

	// cannot be greater than curve.N
	if bintR.Cmp(keytools.Secp256k1.N) == 1 || bintS.Cmp(keytools.Secp256k1.N) == 1 {
		fmt.Printf("Invalid R or S value: cannot be greater than order of secp256k1")
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
	Q := LaksaGo.Compress(keytools.Secp256k1, Qx, Qy, true)

	_r := LaksaGo.Hash(Q, publicKey, msg)

	rn := new(big.Int).SetBytes(r)
	_rn := new(big.Int).SetBytes(_r)
	fmt.Printf("r = %s, _r = %s\n", hex.EncodeToString(r), hex.EncodeToString(_r))
	return rn.Cmp(_rn) == 0
}
