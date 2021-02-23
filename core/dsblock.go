/*
 * Copyright (C) 2021 Zilliqa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
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

func (dst *DsBlock) Serialize() []byte {
	headerBytes := dst.BlockHeader.Serialize()
	bs := &BitVector{}
	headerBytes = dst.Cosigs.CS1.Serialize(headerBytes, uint(len(headerBytes)))
	headerBytes = bs.SetBitVector(headerBytes, uint(len(headerBytes)), dst.Cosigs.B1)
	return headerBytes
}

func (dst *DsBlock) GetRandS() ([]byte, []byte) {
	data := make([]byte, 0)
	bns := BIGNumSerialize{}
	data = bns.SetNumber(data, 0, signatureChallengeSize, dst.Cosigs.CS2.R)
	data = bns.SetNumber(data, 0+signatureChallengeSize, signatureChallengeSize, dst.Cosigs.CS2.S)

	signature := util.EncodeHex(data)
	r := util.DecodeHex(signature[0:64])
	s := util.DecodeHex(signature[64:128])
	return r, s
}

func NewDsBlockFromDsBlockT(dst *DsBlockT) *DsBlock {
	dsBlock := &DsBlock{}
	dsBlockHeader := NewDsBlockHeaderFromDsBlockT(dst)
	dsBlock.BlockHeader = dsBlockHeader
	cs1Ser := util.DecodeHex(dst.CS1)
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
	timestamp, _ := strconv.ParseUint(dst.Header.Timestamp, 10, 64)
	dsBlock.Timestamp = timestamp

	copy(dsBlock.BlockHash[:],util.Sha256(dsBlock.BlockHeader.Serialize()))
	return dsBlock
}

func (dst *DsBlock) ToProtobuf() []byte {
	protoBlockBase := dst.BlockBase.ToProtobuf()
	protoDSBlockHeader := dst.BlockHeader.ToProtobuf(false)

	protoDSBlock := &protobuf.ProtoDSBlock{
		Blockbase: protoBlockBase,
		Header:    protoDSBlockHeader,
	}

	bytes, _ := proto.Marshal(protoDSBlock)
	return bytes
}
