package core

import (
	protobuf "github.com/Zilliqa/gozilliqa-sdk/protobuf"
	"github.com/golang/protobuf/proto"
)

type AccountBase struct {
	Version     uint32
	Balance     uint64
	Nonce       uint64
	StorageRoot []byte
	CodeHash    []byte
}

func AccountBaseFromBytes(bytes []byte) (*AccountBase, error) {
	var protoAccountBase protobuf.ProtoAccountBase
	err := proto.Unmarshal(bytes, &protoAccountBase)
	if err != nil {
		return nil, err
	}

	var accountBase AccountBase
	accountBase.Version = *protoAccountBase.Version

	balanceNum := ByteArrayToUint(protoAccountBase.Balance.Data, 0, 16)
	accountBase.Balance = balanceNum.Uint64()
	accountBase.Nonce = *protoAccountBase.Nonce
	accountBase.CodeHash = protoAccountBase.Codehash
	accountBase.StorageRoot = protoAccountBase.Storageroot

	return &accountBase, nil
}
