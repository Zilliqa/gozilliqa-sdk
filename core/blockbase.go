package core

import "github.com/Zilliqa/gozilliqa-sdk/protobuf"

type BlockBase struct {
	BlockHash [32]byte
	Cosigs    CoSignatures
	Timestamp uint64
}

func (b *BlockBase) ToProtobuf() *protobuf.ProtoBlockBase {
	blockBase := &protobuf.ProtoBlockBase{}
	blockBase.Blockhash = b.BlockHash[:]
	blockBase.Timestamp = b.Timestamp

	cosig := &protobuf.ProtoBlockBase_CoSignatures{
		Cs1: &protobuf.ByteArray{
			Data: b.Cosigs.CS1.Serialize(),
		},
		B1: b.Cosigs.B1,
		Cs2: &protobuf.ByteArray{
			Data: b.Cosigs.CS2.Serialize(),
		},
		B2: b.Cosigs.B2,
	}
	blockBase.Cosigs = cosig

	return blockBase
}
