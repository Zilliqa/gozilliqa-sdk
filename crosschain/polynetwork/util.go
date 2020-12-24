package polynetwork

import (
	"errors"
	"strings"
)

const polyChainPubKeyLen = 67

func SplitPubKeys(rawBytes string) ([]string, error) {
	if strings.HasPrefix(rawBytes,"0x") {
		rawBytes = rawBytes[2:]
	}
	var keys []string
	keyLen := len(rawBytes)
	if keyLen % polyChainPubKeyLen != 0 {
		return keys, errors.New("wrong length of public key list")
	}
	n := keyLen / polyChainPubKeyLen
	for i := 0; i < n; i++ {
		publicKey := rawBytes[i*polyChainPubKeyLen : (i+1)*polyChainPubKeyLen]
		keys = append(keys, "0x" + publicKey)
	}
	return keys, nil
}
