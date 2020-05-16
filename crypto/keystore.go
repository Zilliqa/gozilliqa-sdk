/*
 * Copyright (C) 2019 Zilliqa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"

	"github.com/Zilliqa/gozilliqa-sdk/keytools"
	util2 "github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/google/uuid"
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

	ciphertext := util2.DecodeHex(kv.Crypto.Ciphertext)
	iv := util2.DecodeHex(kv.Crypto.CipherParams.IV)
	kdfparams := kv.Crypto.KDFParams
	kdf := kv.Crypto.KDF
	if kdf == "pbkdf2" {
		derivedKey = ks.pbkdf2.GetDerivedKey([]byte(passphrase), util2.DecodeHex(kdfparams.Salt), 262144, 32)

	} else {
		derivedKey, err = ks.scrypt.GetDerivedKey([]byte(passphrase), util2.DecodeHex(kdfparams.Salt), 8192, 8, 1, 32)
	}

	if err != nil {
		return "", nil
	}

	mac := hex.EncodeToString(util2.GenerateMac(derivedKey, ciphertext, iv))

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
	return util2.EncodeHex(privateKey), nil

}

func (ks *Keystore) EncryptPrivateKey(privateKey, passphrase []byte, t KDFType) (string, error) {
	address := keytools.GetAddressFromPrivateKey(privateKey)
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

	mac := util2.GenerateMac(derivedKey, ciphertext, iv)

	//build struct
	cp := CipherParams{
		IV: util2.EncodeHex(iv),
	}

	kp := NewKDFParams(util2.EncodeHex(salt))

	var kdf string

	if t == 0 {
		kdf = "pbkdf2"
	} else {
		kdf = "scrypt"
	}

	crypto := Crypto{
		Cipher:       "aes-128-ctr",
		CipherParams: cp,
		Ciphertext:   util2.EncodeHex(ciphertext),
		KDF:          kdf,
		KDFParams:    kp,
		MAC:          util2.EncodeHex(mac),
	}

	uid := uuid.New()
	kv := KeystoreV3{
		Address: address,
		Crypto:  crypto,
		ID:      uid.String(),
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
