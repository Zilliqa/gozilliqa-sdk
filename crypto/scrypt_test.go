package crypto

import (
	util "github.com/Zilliqa/gozilliqa-sdk"
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	"strings"
	"testing"
)

func TestScryptWapper_GetDerivedKey(t *testing.T) {
	s := NewScryptWapper()
	keys, err := s.GetDerivedKey([]byte("stronk_password"), util.DecodeHex("2c37db13a633c5a5e5b8c699109690e33860b7eb43bbc81bbab47d4e9c29f1b9"),
		8192, 8, 1, 32)
	if err != nil {
		t.Error("scrypt: get derived key")
	}
	iv, err := keytools.GenerateRandomBytes(16)
	if err != nil {
		t.Error(err.Error())
	}

	macArray := util.GenerateMac(keys, util.DecodeHex("ecdf81453d031ac2fa068b7185ddac044fa4632d3b061400d3c07a86510b4823"),iv)
	if strings.Compare(strings.ToLower(util.EncodeHex(macArray)), "ed7fa37a4adbc8b7bbe0d43a329a047f89e2dcf7f2dfc96babfe79edd955f7a3") != 0 {
		t.Error("crypt: get derived key")
	}
}
