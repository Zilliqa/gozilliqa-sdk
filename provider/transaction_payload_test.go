package provider

import (
	"fmt"
	"testing"
)

func TestNewFromJson(t *testing.T) {
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

	payload, err2 := NewFromJson(data)
	if err2 != nil {
		t.Error(err2.Error())
	}

	fmt.Println(payload)

	data = []byte(`{
    "version": 21823489,
    "nonce": 962,
    "toAddr": "84eb5C96Bec8d29eDdFBe36865E9B7F26b816f0F",
    "amount": 0,
    "gasPrice": 1000000000,
    "gasLimit": 10000,
    "code": "",
    "data": {"_tag":"SubmitCustomMintTransaction","params":[{"vname":"proxyTokenContract","type":"ByStr20","value":"0x39550ab45d74cce5fef70e857c1326b2d9bee096"},{"vname":"to","type":"ByStr20","value":"0x39550ab45d74cce5fef70e857c1326b2d9bee096"},{"vname":"value","type":"Uint128","value":"10000000"}]},
    "signature": "bc44a4cb98184578ab57a8648624a65dc6a74ca8e0cb39df041f30ec373faadf3c2d7ec7e664e8a319b6215b3ec1f5dd948686e46dde31dd580bdd0e79b2890e",
    "pubKey": "0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a"
	}`)

	payload2, err3 := NewFromJson(data)
	if err3 != nil {
		t.Error(err3.Error())
	}
	fmt.Println(payload2)
}

func TestTransactionPayload_ToJson(t *testing.T) {
	pl := TransactionPayload{
		Version:   65537,
		Nonce:     1,
		ToAddr:    "0x0000000000000000000000000000000000000000",
		Amount:    "0",
		PubKey:    "",
		GasPrice:  "10000000",
		GasLimit:  "9000",
		Code:      "",
		Data:      "[{\"vname\":\"_scilla_version\",\"type\":\"Uint32\",\"value\":\"0\"},{\"vname\":\"initial_owners\",\"type\":\"List ByStr20\",\"value\":{\"constructor\":\"Cons\",\"argtypes\":[\"ByStr20\"],\"arguments\":[\"0x1234567890123456789012345678906784567890\",{\"constructor\":\"Cons\",\"argtypes\":[\"ByStr20\"],\"arguments\":[\"0xabcdeabcde123456786782345678901234567890\",{\"constructor\":\"Cons\",\"argtypes\":[\"ByStr20\"],\"arguments\":[\"0xffcdeabcde126786789012345678901234567890\",{\"constructor\":\"Nil\",\"argtypes\":[\"ByStr20\"],\"arguments\":[]}]}]}]}},{\"vname\":\"required_signatures\",\"type\":\"Uint32\",\"value\":\"2\"}]",
		Signature: "",
		//Priority:  false,
	}

	data,err := pl.ToJson()
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(string(data))
	expect := "{\"version\":65537,\"nonce\":1,\"toAddr\":\"0x0000000000000000000000000000000000000000\",\"amount\":0,\"pubKey\":\"\",\"gasPrice\":10000000,\"gasLimit\":9000,\"code\":\"\",\"data\":[{\"vname\":\"_scilla_version\",\"type\":\"Uint32\",\"value\":\"0\"},{\"vname\":\"initial_owners\",\"type\":\"List ByStr20\",\"value\":{\"constructor\":\"Cons\",\"argtypes\":[\"ByStr20\"],\"arguments\":[\"0x1234567890123456789012345678906784567890\",{\"constructor\":\"Cons\",\"argtypes\":[\"ByStr20\"],\"arguments\":[\"0xabcdeabcde123456786782345678901234567890\",{\"constructor\":\"Cons\",\"argtypes\":[\"ByStr20\"],\"arguments\":[\"0xffcdeabcde126786789012345678901234567890\",{\"constructor\":\"Nil\",\"argtypes\":[\"ByStr20\"],\"arguments\":[]}]}]}]}},{\"vname\":\"required_signatures\",\"type\":\"Uint32\",\"value\":\"2\"}],\"signature\":\"\"}"
	fmt.Println(expect)
	if len(string(data)) != len(expect) {
		t.Fail()
	}
}
