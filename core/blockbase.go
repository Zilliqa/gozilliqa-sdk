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

	//cosig := &protobuf.ProtoBlockBase_CoSignatures{}

	return blockBase
}
