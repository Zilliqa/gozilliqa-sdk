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
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransaction_TrackTx(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	provider := provider.NewProvider("https://dev-api.zilliqa.com/")
	tx := Transaction{
		ID:        "",
		Version:   "",
		Nonce:     "",
		Amount:    "",
		GasPrice:  "",
		GasLimit:  "",
		Signature: "",
		Receipt: core.TransactionReceipt{
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

	tx.Confirm("846cda64971e259b1739bf15710758803abcf5754507af5af3f779777cd1b0b0", 1000, 3, provider)
	assert.True(t, tx.Status == core.Confirmed)
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
		"pubKey": ""
}`)

	payload, err2 := provider.NewFromJson(data)
	assert.Nil(t, err2, err2)
	tx := NewFromPayload(payload)
	t.Log(tx)
}

func TestNewFromPayload2(t *testing.T) {
	tx := Transaction{
		ID:           "",
		Version:      "21823489",
		Nonce:        "959",
		Amount:       "0",
		GasPrice:     "1000000000",
		GasLimit:     "10000",
		Signature:    "c0dcffb4f5ef80b9e426c16fc1fb62b31356219deb84c5689ab6a73915ea962c0bc4d4a49985803cd1db8aabb6870e8c749003cab41246e17493767acc6cca90",
		SenderPubKey: "0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a",
		ToAddr:       "84eb5C96Bec8d29eDdFBe36865E9B7F26b816f0F",
		Code:         "",
		Data:         "{\"_tag\":\"SubmitCustomMintTransaction\",\"params\":[{\"vname\":\"proxyTokenContract\",\"type\":\"ByStr20\",\"value\":\"0x39550ab45d74cce5fef70e857c1326b2d9bee096\"},{\"vname\":\"to\",\"type\":\"ByStr20\",\"value\":\"0x39550ab45d74cce5fef70e857c1326b2d9bee096\"},{\"vname\":\"value\",\"type\":\"Uint128\",\"value\":\"10000000\"}]}",
		Status:       0,
	}

	payload := tx.ToTransactionPayload()
	data, err := payload.ToJson()
	assert.Nil(t, err, err)
	t.Log(string(data))
}

func TestNewFromPayload3(t *testing.T) {
	data := []byte(`{
    "version": 65537,
    "nonce": 1,
    "toAddr": "0x0000000000000000000000000000000000000000",
    "amount": 0,
    "gasPrice": 10000000,
    "gasLimit": 9000,
    "code": "",
    "data":"",
    "signature": "",
    "pubKey": ""
}`)

	payload, err2 := provider.NewFromJson(data)
	assert.Nil(t, err2, err2)
	tx := NewFromPayload(payload)
	t.Log(tx)
}
