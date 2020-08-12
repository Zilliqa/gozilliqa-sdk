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
package go_schnorr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	golangAssert "github.com/stretchr/testify/assert"
)

func TestSignVerify(t *testing.T) {
	run_sign_verify_test(t)
}

func run_sign_verify_test(t *testing.T) {
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

		sig, err := SignMessage(priv, pub, msg)

		if err != nil {
			fmt.Printf("err = %s\n", err.Error())
		} else {
			verify := VerifySignature(pub, msg, sig)

			fmt.Printf("signature = %v\n", sig)
			fmt.Printf("expected verify = %v\n", true)
			fmt.Printf("actually verify = %v\n", verify)
			golangAssert.True(t, verify, "verify should be true")
		}
	}
}
