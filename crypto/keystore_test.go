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
	json := "{\"address\":\"0f168fda7796d07735d2cf5763c7511e88a37fc0\",\"id\":\"383ad669-999b-421f-8aa2-e6f7a5f11e99\",\"version\":3,\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"eebde49dcb493c11e9c2055371e47aa3fb76138a8b36cbbd5bb7ce3269dd6d2b\",\"kdf\":\"pbkdf2\",\"mac\":\"beb17af089f467ce5e1213ab980cb08471572281a905a20aeb94e5df6058d6f5\",\"cipherparams\":{\"iv\":\"76389a457903294c36cd31509dfd89b0\"},\"kdfparams\":{\"n\":8192,\"c\":262144,\"r\":8,\"p\":1,\"dklen\":32,\"salt\":\"wIWo67McUXbpakz0t0nraBNUklJXOzatmlDPa6+Y5Ng=\"}}}\n"

	ks := NewDefaultKeystore()
	privateKey, err := ks.DecryptPrivateKey(json, "xiaohuo")
	if err != nil {
		t.Error(err.Error())
	} else {
		if strings.Compare(strings.ToLower(privateKey), "24180e6b0c3021aedb8f5a86f75276ee6fc7ff46e67e98e716728326102e91c9") != 0 {
			t.Error("decrypt private key failed")
		}
	}

}
