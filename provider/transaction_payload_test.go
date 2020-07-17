package provider

import (
	"fmt"
	"os"
	"testing"
)

func TestNewFromJson(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	data := []byte(`{"version":21823489,"nonce":166471,"toAddr":"bd7198209529dC42320db4bC8508880BcD22a9f2","amount":0,"pubKey":"0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a","gasPrice":1000000000,"gasLimit":1000,"code":"","data":{"_tag":"Transfer","params":[{"vname":"to","type":"ByStr20","value":"0x9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a"},{"vname":"tokens","type":"Uint128","value":"10"}]},"signature":"e4ae77ba4534598a723a3792705dd7477ffef8a475da94388d1af0cd29b38a3737b64fbedc4eb72c66b28303ac1b0bb8c45a68da2d40d6def367cbf20e751747"}`)

	payload, err2 := NewFromJson(data)
	if err2 != nil {
		t.Error(err2.Error())
	}

	fmt.Println(payload)

	payload2, err3 := NewFromJson(data)
	if err3 != nil {
		t.Error(err3.Error())
	}
	provider := NewProvider("https://dev-api.zilliqa.com")
	rsp, err := provider.CreateTransaction(*payload2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(rsp.Error)
		fmt.Println(rsp.Result)
	}

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
		Data:      "",
		Signature: "",
		//Priority:  false,
	}

	data, err := pl.ToJson()
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(string(data))
	expect := "{\"version\":65537,\"nonce\":1,\"toAddr\":\"0x0000000000000000000000000000000000000000\",\"amount\":0,\"pubKey\":\"\",\"gasPrice\":10000000,\"gasLimit\":9000,\"code\":\"\",\"data\":\"\",\"signature\":\"\"}"
	if len(string(data)) != len(expect) {
		t.Fail()
	}
}
