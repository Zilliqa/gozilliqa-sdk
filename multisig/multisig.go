package multisig

import (
	"crypto/elliptic"
	"errors"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"github.com/btcsuite/btcd/btcec"
)

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
		bytes := elliptic.MarshalCompressed(keytools.Secp256k1, x, y)
		pubKey, err2 := btcec.ParsePubKey(bytes, keytools.Secp256k1)
		if err2 != nil {
			return nil, err2
		}

		aggregatedPubKey = pubKey

	}

	return aggregatedPubKey.SerializeCompressed(), nil
}
