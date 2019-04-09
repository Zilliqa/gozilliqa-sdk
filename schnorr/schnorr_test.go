package go_schnorr

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestTrySign(t *testing.T) {
	run_sign_test(t)
}

func run_sign_test(t *testing.T) {
	b, err := ioutil.ReadFile("data")
	if err != nil {
		panic("read file failed")
	}

	var data []map[string]string
	err2 := json.Unmarshal(b, &data)

	if err2 != nil {
		panic("unmarshal failed")
	}

	for _, v := range data {
		msg := hex_bytes(v["msg"])
		pub := hex_bytes(v["pub"])
		priv := hex_bytes(v["priv"])
		k := hex_bytes(v["k"])
		re := v["r"]
		se := v["s"]
		r, s, err := TrySign(priv, pub, msg, k)
		if err != nil {
			fmt.Printf("err = %s\n", err.Error())
		} else {
			fmt.Printf("expected r = %s, s = %s\n", re, se)
			fmt.Printf("actually r = %s, s = %s\n", hex.EncodeToString(r), hex.EncodeToString(s))
			assert(re, hex.EncodeToString(r), t)
			assert(se, hex.EncodeToString(s), t)
		}
	}
}

func TestVerify(t *testing.T) {
	run_verify_test(t)
}

func run_verify_test(t *testing.T) {
	b, err := ioutil.ReadFile("data")
	if err != nil {
		panic("read file failed")
	}

	var data []map[string]string
	err2 := json.Unmarshal(b, &data)

	if err2 != nil {
		panic("unmarshal failed")
	}

	fmt.Printf("test data number = %d", len(data))

	n := 0

	for _, v := range data {
		n++
		msg := hex_bytes(v["msg"])
		pub := hex_bytes(v["pub"])
		r := hex_bytes(v["r"])
		s := hex_bytes(v["s"])
		result := Verify(pub, msg, r, s)
		if !result {
			fmt.Printf("r = %s\n", hex.EncodeToString(r))
			panic("verify failed")
		}
	}

	fmt.Printf("n = %d", n)
}

func hex_bytes(hs string) []byte {
	data, err := hex.DecodeString(hs)
	if err != nil {
		panic("cannot convert hex string to byte array")
	}
	return data
}

func assert(expected string, actually string, t *testing.T) {
	if strings.Compare(expected, strings.ToUpper(actually)) != 0 {
		t.Errorf("expected = %s, actually = %s", expected, actually)
	}
}
