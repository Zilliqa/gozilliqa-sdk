package account

import (
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"testing"
)

var f = "{\"address\":\"9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a\",\"id\":\"1497eb45-3a52-4c5a-97eb-88d5e790fcd0\",\"version\":3,\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"3ddd39cb13c95ccdc150c962fadaebfa7a2fca3221c81e276491d70a5d621dd5\",\"kdf\":\"pbkdf2\",\"mac\":\"980f95923582693dad2038ea4e1119a934332c53d620ebe38b7e3b7928e57d05\",\"cipherparams\":{\"iv\":\"39a7beef25795f912572718363dba9f4\"},\"kdfparams\":{\"n\":8192,\"c\":262144,\"r\":8,\"p\":1,\"dklen\":32,\"salt\":\"4f3ddae640ebe3cb45a133c583d03e5da25c36baf4472343fb5f6a0c899b78f1\"}}}"

func TestToFile(t *testing.T) {
	file, err := ToFile("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930", "xiaohuo", 0)
	if err != nil {
		t.Error(err.Error())
	}
	if f != file {
		t.Failed()
	}
}

func TestFromFile(t *testing.T) {
	a, err := FromFile(f, "xiaohuo")
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(util.EncodeHex(a.PrivateKey))
	if util.EncodeHex(a.PrivateKey) != "e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930" {
		t.Failed()
	}
}
