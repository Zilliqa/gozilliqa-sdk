package crypto

import (
	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	util2 "github.com/Zilliqa/gozilliqa-sdk/util"
	"strings"
	"testing"
)

func TestGetDerivedKey(t *testing.T) {
	p := NewPbkdf2()
	keys := p.GetDerivedKey([]byte("stronk_password"), util2.DecodeHex("0f2274f6c0daf36d5822d97985be5a3d881d11e2e741bad4e038a099eecc3b6d"), 262144, 32)
	iv, _ := keytools.GenerateRandomBytes(16)

	macArray := util2.GenerateMac(keys, util2.DecodeHex("dc55047d51f795509ffb6969db837a4481887ccfb6bfb7c259fb77b19078c2a4"), iv)
	println()
	if strings.Compare(strings.ToLower("DEDC361C53C421974C2811F7F989BC530AEBF9A90C487B4161E0E54AE6FABA31"), strings.ToLower(util2.EncodeHex(macArray))) != 0 {
		t.Error("get derived key error")
	}
}
