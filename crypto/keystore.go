package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	util "github.com/FireStack-Lab/LaksaGo"
	"github.com/FireStack-Lab/LaksaGo/keytools"
	uuid "github.com/satori/go.uuid"
	"strings"
)

// 0: p 1:s
type KDFType int

type Keystore struct {
	pbkdf2 *pbkdf2Wapper
	scrypt *scryptWapper
}

func NewDefaultKeystore() *Keystore {
	p := NewPbkdf2()
	s := NewScryptWapper()
	return &Keystore{
		pbkdf2: p,
		scrypt: s,
	}
}

func NewKeystore(p *pbkdf2Wapper, s *scryptWapper) *Keystore {
	return &Keystore{
		pbkdf2: p,
		scrypt: s,
	}
}

func (ks *Keystore) GetDerivedKey(password []byte, params interface{}) ([]byte, error) {
	p, ok := params.(Pbkdf2Params)
	if ok {
		return ks.pbkdf2.GetDerivedKey(password, p.Salt, p.Count, p.DkLen), nil
	}

	s, ok := params.(ScryptParams)
	if ok {
		r, err := ks.scrypt.GetDerivedKey(password, s.Salt, s.N, s.R, s.P, s.DkLen)
		if err != nil {
			return nil, err
		} else {
			return r, nil
		}
	}

	return nil, errors.New("unsupport params")
}

func (ks *Keystore) DecryptPrivateKey(encryptJson, passphrase string) (string, error) {
	var kv KeystoreV3
	err := json.Unmarshal([]byte(encryptJson), &kv)
	if err != nil {
		return "", err
	}

	derivedKey := make([]byte, 32)
	err = nil

	ciphertext := util.DecodeHex(kv.Crypto.Ciphertext)
	iv := util.DecodeHex(kv.Crypto.CipherParams.IV)
	kdfparams := kv.Crypto.KDFParams
	kdf := kv.Crypto.KDF
	fmt.Println(kdfparams.Salt)
	if kdf == "pbkdf2" {
		derivedKey = ks.pbkdf2.GetDerivedKey([]byte(passphrase), util.DecodeHex(kdfparams.Salt), 262144, 32)

	} else {
		derivedKey, err = ks.scrypt.GetDerivedKey([]byte(passphrase), util.DecodeHex(kdfparams.Salt), 8192, 8, 1, 32)
	}

	if err != nil {
		return "", nil
	}

	mac := hex.EncodeToString(util.GenerateMac(derivedKey, ciphertext, iv))

	if strings.Compare(strings.ToLower(mac), strings.ToLower(kv.Crypto.MAC)) != 0 {
		return "", errors.New("Failed to decrypt.")
	}

	encryptKey := make([]byte, 16)
	copy(encryptKey[:], derivedKey[0:16])

	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return "", err
	}

	privateKey := make([]byte, len(ciphertext))
	mode := cipher.NewCTR(block, iv)
	mode.XORKeyStream(privateKey, ciphertext)
	return util.EncodeHex(privateKey), nil

}

func (ks *Keystore) EncryptPrivateKey(privateKey, passphrase []byte, t KDFType) (string, error) {
	address := keytools.GetAddressFromPublic(privateKey)
	iv, err := keytools.GenerateRandomBytes(16)
	if err != nil {
		return "", err
	}
	salt, err := keytools.GenerateRandomBytes(32)
	if err != nil {
		return "", err
	}

	derivedKey := make([]byte, 32)
	err = nil

	if t == 0 {
		derivedKey = ks.pbkdf2.GetDerivedKey([]byte(passphrase), salt, 262144, 32)
	} else {
		derivedKey, err = ks.scrypt.GetDerivedKey([]byte(passphrase), salt, 8192, 8, 1, 32)
	}
	if err != nil {
		return "", nil
	}

	encryptKey := make([]byte, 16)
	copy(encryptKey, derivedKey[0:16])

	//perform cipher operation
	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(privateKey))
	mode := cipher.NewCTR(block, iv)
	mode.XORKeyStream(ciphertext, privateKey)

	mac := util.GenerateMac(derivedKey, ciphertext, iv)

	//build struct
	cp := CipherParams{
		IV: util.EncodeHex(iv),
	}

	kp := NewKDFParams(util.EncodeHex(salt))

	var kdf string

	if t == 0 {
		kdf = "pbkdf2"
	} else {
		kdf = "scrypt"
	}

	crypto := Crypto{
		Cipher:       "aes-128-ctr",
		CipherParams: cp,
		Ciphertext:   util.EncodeHex(ciphertext),
		KDF:          kdf,
		KDFParams:    kp,
		MAC:          util.EncodeHex(mac),
	}

	kv := KeystoreV3{
		Address: address,
		Crypto:  crypto,
		ID:      uuid.NewV4().String(),
		Version: 3,
	}

	kvstring, err := json.Marshal(kv)
	if err != nil {
		return "", err
	}
	return string(kvstring), nil
}

type KeystoreV3 struct {
	Address string `json:"address"`
	ID      string `json:"id"`
	Version int    `json:"version"`
	Crypto  `json:"crypto"`
}

type CipherParams struct {
	IV string `json:"iv"`
}

type Crypto struct {
	Cipher       string `json:"cipher"`
	Ciphertext   string `json:"ciphertext"`
	KDF          string `json:"kdf"`
	MAC          string `json:"mac"`
	CipherParams `json:"cipherparams"`
	KDFParams    `json:"kdfparams"`
}
type KDFParams struct {
	N     int    `json:"n"`
	C     int    `json:"c"`
	R     int    `json:"r"`
	P     int    `json:"p"`
	DKlen int    `json:"dklen"`
	Salt  string `json:"salt"`
}

func NewKDFParams(salt string) KDFParams {
	return KDFParams{
		N:     8192,
		C:     262144,
		R:     8,
		P:     1,
		DKlen: 32,
		Salt:  salt,
	}
}
