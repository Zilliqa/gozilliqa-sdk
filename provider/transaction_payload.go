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
	"reflect"
	"strconv"
	"strings"
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

type Value struct {
	VName string      `json:"vname"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Data struct {
	Tag    string  `json:"_tag"`
	Params []Value `json:"params"`
}

type payload struct {
	Version   int    `json:"version"`
	Nonce     int    `json:"nonce"`
	ToAddr    string `json:"toAddr"`
	Amount    int64  `json:"amount"`
	PubKey    string `json:"pubKey"`
	GasPrice  int64  `json:"gasPrice"`
	GasLimit  int64  `json:"gasLimit"`
	Code      string `json:"code"`
	Data      Data   `json:"data"`
	Signature string `json:"signature"`
	//Priority  bool
}

type Init struct {
	Version   int           `json:"version"`
	Nonce     int           `json:"nonce"`
	ToAddr    string        `json:"toAddr"`
	Amount    int64         `json:"amount"`
	PubKey    string        `json:"pubKey"`
	GasPrice  int64         `json:"gasPrice"`
	GasLimit  int64         `json:"gasLimit"`
	Code      string        `json:"code"`
	Data      []interface{} `json:"data"`
	Signature string        `json:"signature"`
}

type Payment struct {
	Version   int    `json:"version"`
	Nonce     int    `json:"nonce"`
	ToAddr    string `json:"toAddr"`
	Amount    int64  `json:"amount"`
	PubKey    string `json:"pubKey"`
	GasPrice  int64  `json:"gasPrice"`
	GasLimit  int64  `json:"gasLimit"`
	Code      string `json:"code"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

func (pl *TransactionPayload) ToJson() ([]byte, error) {
	a, err := strconv.ParseInt(pl.Amount, 10, 64)
	if err != nil {
		return nil, err
	}

	price, err2 := strconv.ParseInt(pl.GasPrice, 10, 64)
	if err2 != nil {
		return nil, err2
	}

	limit, err3 := strconv.ParseInt(pl.GasLimit, 10, 64)
	if err3 != nil {
		return nil, err3
	}

	if pl.Data == "" {
		p := Payment{
			Version:   pl.Version,
			Nonce:     pl.Nonce,
			ToAddr:    pl.ToAddr,
			Amount:    a,
			PubKey:    pl.PubKey,
			GasPrice:  price,
			GasLimit:  limit,
			Code:      pl.Code,
			Data:      "",
			Signature: pl.Signature,
			//Priority:  pl.Priority,
		}
		return json.Marshal(&p)
	}

	originData := strings.TrimPrefix(pl.Data, `"`)
	originData = strings.TrimSuffix(originData, `"`)
	originData = strings.ReplaceAll(originData, "\\", "")

	var data Data
	err4 := json.Unmarshal([]byte(originData), &data)
	if err4 != nil {
		var dataArray []interface{}
		err5 := json.Unmarshal([]byte(originData), &dataArray)
		if err5 != nil {
			return nil, errors.New("wrong data")
		} else {
			p := Init{
				Version:   pl.Version,
				Nonce:     pl.Nonce,
				ToAddr:    pl.ToAddr,
				Amount:    a,
				PubKey:    pl.PubKey,
				GasPrice:  price,
				GasLimit:  limit,
				Code:      pl.Code,
				Data:      dataArray,
				Signature: pl.Signature,
				//Priority:  pl.Priority,
			}
			return json.Marshal(&p)
		}
	} else {
		p := payload{
			Version:   pl.Version,
			Nonce:     pl.Nonce,
			ToAddr:    pl.ToAddr,
			Amount:    a,
			PubKey:    pl.PubKey,
			GasPrice:  price,
			GasLimit:  limit,
			Code:      pl.Code,
			Data:      data,
			Signature: pl.Signature,
			//Priority:  pl.Priority,
		}
		return json.Marshal(&p)
	}
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

	type Value struct {
		VName string      `json:"vname"`
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	}

	type Data struct {
		Tag    string  `json:"_tag"`
		Params []Value `json:"params"`
	}

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

	price, ok4 := middle["gasPrice"].(float64)
	if !ok4 {
		return nil, errors.New("parse payload json failed: price")
	}

	limit, ok5 := middle["gasLimit"].(float64)
	if !ok5 {
		return nil, errors.New("parse payload json failed: limit")
	}

	pubkey, ok6 := middle["pubKey"].(string)
	if !ok6 {
		return nil, errors.New("parse payload json failed: public key")
	}

	code, ok7 := middle["code"].(string)
	if !ok7 {
		return nil, errors.New("parse payload json failed: code")
	}

	d, ok8 := middle["data"]
	if !ok8 {
		return nil, errors.New("parse payload json failed: data")
	}

	var sd string

	if reflect.TypeOf(d).Kind() == reflect.Slice {
		dd := d.([]interface{})
		s, err := json.Marshal(dd)
		if err != nil {
			return nil, err
		}
		sd = string(s)
	} else if reflect.TypeOf(d).Kind() == reflect.String {
		sd = d.(string)
	} else {
		j, _ := json.Marshal(d)
		var data Data
		err := json.Unmarshal(j, &data)
		if err != nil {
			return nil, err
		}
		ss, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		sd = string(ss)
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
		Data:      sd,
		Signature: sig,
		Priority:  false,
	}, nil
}
