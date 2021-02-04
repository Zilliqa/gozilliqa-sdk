package core

import (
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"math/big"
	"strings"
	"testing"
)

func TestSignature_Serialize(t *testing.T) {
	rs := "3AF3D288E830E96FF8ED0769F45ABDA774CD989E2AE32EF9E985C8505F14FF98"
	ss := "E191EB14A70B5B53ADA45AFFF4A04578F5D8BB2B1C8A22985EA159B53826CDE7"
	r, _ := new(big.Int).SetString(rs, 16)
	s, _ := new(big.Int).SetString(ss, 16)
	signature := &Signature{
		r, s,
	}
	sig := signature.Serialize()
	if strings.ToUpper(util.EncodeHex(sig)) != "3AF3D288E830E96FF8ED0769F45ABDA774CD989E2AE32EF9E985C8505F14FF98E191EB14A70B5B53ADA45AFFF4A04578F5D8BB2B1C8A22985EA159B53826CDE7" {
		t.Fail()
	}

	rs = "948AFFFF6E068CA2F2757BFD6085D6E4C3084B038E5533C5927ECB19EA0D329C"
	ss = "DFEE66E2C4799E73F0F778126A23032608408C27C2E7B3FA45A626BB9BDEB53C"
	r, _ = new(big.Int).SetString(rs, 16)
	s, _ = new(big.Int).SetString(ss, 16)
	signature = &Signature{
		r, s,
	}
	sig = signature.Serialize()
	if strings.ToUpper(util.EncodeHex(sig)) != "948AFFFF6E068CA2F2757BFD6085D6E4C3084B038E5533C5927ECB19EA0D329CDFEE66E2C4799E73F0F778126A23032608408C27C2E7B3FA45A626BB9BDEB53C" {
		t.Fail()
	}
}
