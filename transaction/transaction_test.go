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
package transaction

import (
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"testing"
)

func TestTransaction_TrackTx(t *testing.T) {
	provider := provider.NewProvider("https://dev-api.zilliqa.com/")
	tx := Transaction{
		ID:           "",
		Version:      "",
		Nonce:        "",
		Amount:       "",
		GasPrice:     "",
		GasLimit:     "",
		Signature:    "",
		Receipt:      TransactionReceipt{
			Success:       false,
			CumulativeGas: "",
			EpochNum:      "",
		},
		SenderPubKey: "",
		ToAddr:       "",
		Code:         "",
		Data:         "",
		Status:       0,
	}

	tx.Confirm("846cda64971e259b1739bf15710758803abcf5754507af5af3f779777cd1b0b0",1000,3,provider)
}

func TestNewFromPayload(t *testing.T) {
	data := []byte(`{
    "version": 65537,
    "nonce": 1,
    "toAddr": "0x0000000000000000000000000000000000000000",
    "amount": 0,
    "gasPrice": 10000000,
    "gasLimit": 9000,
    "code": "",
    "data": [
        {
            "vname": "_scilla_version",
            "type": "Uint32",
            "value": "0"
        },
        {
            "vname": "initial_owners",
            "type": "List ByStr20",
            "value": {
                "constructor": "Cons",
                "argtypes": [
                    "ByStr20"
                ],
                "arguments": [
                    "0x1234567890123456789012345678906784567890",
                    {
                        "constructor": "Cons",
                        "argtypes": [
                            "ByStr20"
                        ],
                        "arguments": [
                            "0xabcdeabcde123456786782345678901234567890",
                            {
                                "constructor": "Cons",
                                "argtypes": [
                                    "ByStr20"
                                ],
                                "arguments": [
                                    "0xffcdeabcde126786789012345678901234567890",
                                    {
                                        "constructor": "Nil",
                                        "argtypes": [
                                            "ByStr20"
                                        ],
                                        "arguments": []
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        },
        {
            "vname": "required_signatures",
            "type": "Uint32",
            "value": "2"
        }
    ],
        "signature": "",
		"PubKey": ""
}`)

	payload, err2 := provider.NewFromJson(data)
	if err2 != nil {
		t.Error(err2.Error())
	}
	tx := NewFromPayload(payload)
	fmt.Println(tx)
}
