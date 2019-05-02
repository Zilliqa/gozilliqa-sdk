package crypto

import (
	util "github.com/FireStack-Lab/LaksaGo"
	"testing"
)

func TestKeystore_EncryptPrivateKey(t *testing.T) {
	ks := NewDefaultKeystore()
	kv, err := ks.EncryptPrivateKey(util.DecodeHex("24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9"), []byte("xiaohuo"), 0)
	if err != nil {
		t.Error(err.Error())
	} else {
		println(kv)
	}
}
