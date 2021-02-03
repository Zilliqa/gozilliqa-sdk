package core

import (
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"math/big"
	"testing"
)

func TestPeer_Serialize(t *testing.T) {
	ip := new(big.Int).SetUint64(16777343)
	peer := Peer{
		IpAddress:      ip,
		ListenPortHost: 0,
	}

	data := peer.Serialize()
	if util.EncodeHex(data) != "0000000000000000000000000100007F00000000" {
		t.Failed()
	}

	ip = new(big.Int).SetUint64(0)
	peer = Peer{
		IpAddress:      ip,
		ListenPortHost: 0,
	}
	data = peer.Serialize()
	if util.EncodeHex(data) != "0000000000000000000000000000000000000000" {
		t.Failed()
	}
}
