package transaction

import (
	"github.com/Zilliqa/gozilliqa-sdk"
	"strings"
	"testing"
)

func TestEncodeTransactionProto(t *testing.T) {
	txParams := TxParams{
		Version:      "0",
		Nonce:        "0",
		ToAddr:       "2E3C9B415B19AE4035503A06192A0FAD76E04243",
		SenderPubKey: "0246e7178dc8253201101e18fd6f6eb9972451d121fc57aa2a06dd5c111e58dc6a",
		Amount:       "10000",
		GasPrice:     "100",
		GasLimit:     "1000",
		Code:         "",
		Data:         "",
	}

	bytes, _ := EncodeTransactionProto(txParams)
	if strings.Compare(strings.ToUpper(LaksaGo.EncodeHex(bytes)),"080010001A142E3C9B415B19AE4035503A06192A0FAD76E0424322230A210246E7178DC8253201101E18FD6F6EB9972451D121FC57AA2A06DD5C111E58DC6A2A120A100000000000000000000000000000271032120A100000000000000000000000000000006438E807") != 0 {
		t.Error("encode error")
	}
}
