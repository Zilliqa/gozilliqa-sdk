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
package provider

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TransactionPayload struct {
	Version   int    `json:"version"`
	Nonce     int    `json:"nonce"`
	ToAddr    string `json:"toAddr"`
	Amount    string `json:"amount"`
	PubKey    string `json:"pubKey"`
	GasPrice  string `json:"gasPrice"`
	GasLimit  string `json:"gasLimit"`
	Code      string `json:"code"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
	Priority  bool   `json:"priority"`
}

// some data fields don't match, so we need middle map
// see the unit test
func NewFromJson(data []byte) (*TransactionPayload, error) {
	var middle map[string]interface{}
	err := json.Unmarshal(data, &middle)
	if err != nil {
		return nil, err
	}

	return NewFromMap(middle)

}

func NewFromMap(middle map[string]interface{}) (*TransactionPayload, error) {
	v, ok := middle["version"].(float64)
	if !ok {
		return nil, errors.New("parse payload json failed: version")
	}

	n, ok1 := middle["nonce"].(float64)
	if !ok1 {
		return nil, errors.New("parse payload json failed: nonce")
	}

	toAddr, ok2 := middle["toAddr"]
	if !ok2 {
		return nil, errors.New("parse payload json failed: toAddr")
	}

	amount, ok3 := middle["amount"].(float64)
	if !ok3 {
		return nil, errors.New("parse payload json failed: amount")
	}
	fmt.Println(amount)

	price, ok4 := middle["gasPrice"].(float64)
	if !ok4 {
		return nil, errors.New("parse payload json failed: price")
	}

	limit, ok5 := middle["gasLimit"].(float64)
	if !ok5 {
		return nil, errors.New("parse payload json failed: limit")
	}

	pubkey, ok6 := middle["pubKey"].(string)
	// todo resolve me
	if !ok6 {
		//return nil, errors.New("parse payload json failed: public key")
	}

	code, ok7 := middle["code"].(string)
	if !ok7 {
		return nil, errors.New("parse payload json failed: code")
	}

	d, ok8 := middle["data"].([]interface{})
	if !ok8 {
		return nil, errors.New("parse payload json failed: data")
	}

	data, err := json.Marshal(d)
	if err != nil {
		return nil, errors.New("parse data failed")
	}

	sig, ok9 := middle["signature"].(string)
	if !ok9 {
		return nil, errors.New("parse payload json failed: signature")
	}

	return &TransactionPayload{
		Version:   int(v),
		Nonce:     int(n),
		ToAddr:    toAddr.(string),
		Amount:    fmt.Sprintf("%.0f", amount),
		PubKey:    pubkey,
		GasPrice:  fmt.Sprintf("%.0f", price),
		GasLimit:  fmt.Sprintf("%.0f", limit),
		Code:      code,
		Data:      string(data),
		Signature: sig,
		Priority:  false,
	}, nil
}
