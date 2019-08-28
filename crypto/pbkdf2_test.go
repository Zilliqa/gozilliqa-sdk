package crypto

import (
	util "github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	"strings"
	"testing"
)

func TestGetDerivedKey(t *testing.T) {
	p := NewPbkdf2()
	keys := p.GetDerivedKey([]byte("stronk_password"), util.DecodeHex("0f2274f6c0daf36d5822d97985be5a3d881d11e2e741bad4e038a099eecc3b6d"), 262144, 32)
	iv, _ := keytools.GenerateRandomBytes(16)

	macArray := util.GenerateMac(keys, util.DecodeHex("dc55047d51f795509ffb6969db837a4481887ccfb6bfb7c259fb77b19078c2a4"), iv)
	println()
	if strings.Compare(strings.ToLower("DEDC361C53C421974C2811F7F989BC530AEBF9A90C487B4161E0E54AE6FABA31"), strings.ToLower(util.EncodeHex(macArray))) != 0 {
		t.Error("get derived key error")
	}
}
