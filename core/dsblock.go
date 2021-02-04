package core

import (
	"github.com/Zilliqa/gozilliqa-sdk/protobuf"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"github.com/golang/protobuf/proto"
	"strconv"
)

type DsBlock struct {
	BlockBase
	BlockHeader *DsBlockHeader
}

func NewDsBlockFromDsBlockT(dst *DsBlockT) *DsBlock {
	dsBlock := &DsBlock{}
	dsBlockHeader := NewDsBlockHeaderFromDsBlockT(dst)
	dsBlock.BlockHeader = dsBlockHeader
	cs1Ser:= util.DecodeHex(dst.CS1)
	cs2Ser := util.DecodeHex(dst.Signatures)

	cs1 := NewFromByteArray(cs1Ser)
	cs2 := NewFromByteArray(cs2Ser)

	cosig := CoSignatures{
		CS1: cs1,
		B1:  dst.B1,
		CS2: cs2,
		B2:  dst.B2,
	}
	dsBlock.Cosigs = cosig
	timestamp,_ := strconv.ParseUint(dst.Header.Timestamp,10,64)
	dsBlock.Timestamp = timestamp

	return dsBlock
}

func (ds *DsBlock) ToProtobuf() []byte {
	protoBlockBase := ds.BlockBase.ToProtobuf()
	protoDSBlockHeader := ds.BlockHeader.ToProtobuf(false)

	protoDSBlock := &protobuf.ProtoDSBlock{
		Blockbase: protoBlockBase,
		Header:    protoDSBlockHeader,
	}

	bytes, _ := proto.Marshal(protoDSBlock)
	return bytes
}
