package crypto

import (
	util "github.com/FireStack-Lab/LaksaGo"
	"strings"
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

func TestKeystore_DecryptPrivateKey(t *testing.T) {
	json := "{\"address\": \"c57e04e20f452ff617de1f97c75ab14bff64bbe6\", \"crypto\": {\"cipher\": \"aes-128-ctr\", \"cipherparams\": {\"iv\": \"43ecfcbfbc28973d778cd1069ab7148e\"}, \"ciphertext\": \"d5fa57ed9571fe03c2d54e00b6f9e42bb1260750ac8817424089af47f796e19c\", \"kdf\": \"pbkdf2\", \"kdfparams\": {\"salt\": \"8b04769629839019a9b65e7392506730da70331312dc049869d182224f579a56\", \"n\": 8192, \"c\": 262144, \"r\": 8, \"p\": 1, \"dklen\": 32}, \"mac\": \"f1634b61f8e0e713b6d3aa052c46da0f2cc0b922f9c91d3405655f3e92672f61\"}, \"id\": \"810c2f4d-8b69-4012-bd9f-41e86695c72d\", \"version\": 3}"

	ks := NewDefaultKeystore()
	privateKey, err := ks.DecryptPrivateKey(json, "MyKey")
	if err != nil {
		t.Error(err.Error())
	} else {
		if strings.Compare(strings.ToUpper(privateKey), "3B6674116AF2B954675E6373AC27E6A5CE03BCC8675ECDB7915AC8EE68B7ADCF") != 0 {
			t.Error("decrypt private key failed")
		}
	}

}
