package core

import "github.com/Zilliqa/gozilliqa-sdk/protobuf"

type BlockHeaderBase struct {
	Version uint32
	// Hash for the committee that generated the block
	CommitteeHash [32]byte
	PrevHash [32]byte
}

func (b *BlockHeaderBase) ToProtobuf() *protobuf.ProtoBlockHeaderBase {
	protoBlockHeaderBase := &protobuf.ProtoBlockHeaderBase{}
	protoBlockHeaderBase.Version = b.Version
	protoBlockHeaderBase.Committeehash = b.CommitteeHash[:]
	protoBlockHeaderBase.Prevhash = b.PrevHash[:]
	return protoBlockHeaderBase
}
