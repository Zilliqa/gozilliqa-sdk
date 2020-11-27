package multisig

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/btcsuite/btcd/btcec"
	"math/big"
)

var bintZero = big.NewInt(0)

const ThirdDomainSeparatedHashFunctionByte = 0x11

func AggregatedPubKey(pubKeys [][]byte) ([]byte, error) {
	if len(pubKeys) == 0 {
		return nil, errors.New("empty public key list")
	}
	var aggregatedPubKey *btcec.PublicKey
	key, err := btcec.ParsePubKey(pubKeys[0], keytools.Secp256k1)
	if err != nil {
		return nil, err
	}
	aggregatedPubKey = key
	for i := 1; i < len(pubKeys); i++ {
		puk, err1 := btcec.ParsePubKey(pubKeys[i], keytools.Secp256k1)
		if err1 != nil {
			return nil, err1
		}
		x, y := keytools.Secp256k1.Add(aggregatedPubKey.X, aggregatedPubKey.Y, puk.X, puk.Y)
		bytes := util.Marshal(keytools.Secp256k1, x, y, true)
		pubKey, err2 := btcec.ParsePubKey(bytes, keytools.Secp256k1)
		if err2 != nil {
			return nil, err2
		}

		aggregatedPubKey = pubKey

	}

	return aggregatedPubKey.SerializeCompressed(), nil
}

func MultiVerify(publicKey []byte, msg []byte, r []byte, s []byte) bool {
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
	Q := util.Compress(keytools.Secp256k1, Qx, Qy, true)

	_r := hash(ThirdDomainSeparatedHashFunctionByte, Q, publicKey, msg)
	_rn := new(big.Int).Mod(new(big.Int).SetBytes(_r), keytools.Secp256k1.N)

	rn := new(big.Int).SetBytes(r)
	return rn.Cmp(_rn) == 0
}

func hash(first byte, Q []byte, pubKey []byte, msg []byte) []byte {
	var buffer bytes.Buffer
	buffer.WriteByte(first)
	buffer.Write(Q)
	buffer.Write(pubKey[:])
	buffer.Write(msg)
	return util.Sha256(buffer.Bytes())
}
