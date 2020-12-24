package polynetwork

import (
	"errors"
	"strings"
)

const polyChainPubKeyLen = 67 * 2
const polyChainSignatureLen = 65 * 2

func SplitPubKeys(rawBytes string) ([]string, error) {
	return split(rawBytes, polyChainPubKeyLen)
}

func SplitSignature(rawBytes string) ([]string, error) {
	return split(rawBytes, polyChainSignatureLen)
}

func split(rawBytes string, l int) ([]string, error) {
	if strings.HasPrefix(rawBytes, "0x") {
		rawBytes = rawBytes[2:]
	}
	var keys []string
	keyLen := len(rawBytes)
	if keyLen%l != 0 {
		return keys, errors.New("wrong length of public key list")
	}
	n := keyLen / l
	for i := 0; i < n; i++ {
		publicKey := rawBytes[i*l : (i+1)*l]
		keys = append(keys, "0x"+publicKey)
	}
	return keys, nil
}
