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
	"math/big"
	"strconv"
)

type TxBlockHeader struct {
	BlockHeaderBase BlockHeaderBase
	GasLimit        uint64
	GasUsed         uint64
	// uint128
	Rewards *big.Int
	// block index, starting from 0 in the genesis block
	BlockNum uint64
	HashSet  TxBlockHashSet
	NumTxs   uint32
	// Leader of the committee who proposed this block
	// base16 string
	MinerPubKey string
	// DS Block index at the time this Tx block was proposed
	DSBlockNum uint64
}

func NewTxBlockHeaderFromTxBlockT(txt *TxBlockT) *TxBlockHeader {
	header := &TxBlockHeader{}
	gasLimit, _ := strconv.ParseUint(txt.Header.GasLimit, 10, 64)
	gasUsed, _ := strconv.ParseUint(txt.Header.GasUsed, 10, 64)
	header.GasLimit = gasLimit
	header.GasUsed = gasUsed

	rewards, _ := new(big.Int).SetString(txt.Header.Rewards, 10)
	header.Rewards = rewards

	blockNum, _ := strconv.ParseUint(txt.Header.BlockNum, 10, 64)
	header.BlockNum = blockNum

	var txHashSet TxBlockHashSet
	copy(txHashSet.stateRootHash[:], util.DecodeHex(txt.Header.StateRootHash))
	copy(txHashSet.deltaHash[:], util.DecodeHex(txt.Header.StateDeltaHash))
	copy(txHashSet.mbInfoHash[:], util.DecodeHex(txt.Header.MbInfoHash))
	header.HashSet = txHashSet

	header.NumTxs = txt.Header.NumTxns
	header.MinerPubKey = txt.Header.MinerPubKey
	dsBlockNum, _ := strconv.ParseUint(txt.Header.DSBlockNum, 10, 64)
	header.DSBlockNum = dsBlockNum

	header.BlockHeaderBase.Version = uint32(txt.Header.Version)

	ch := util.DecodeHex(txt.Header.CommitteeHash)
	var commitHash [32]byte
	copy(commitHash[:], ch)
	header.BlockHeaderBase.CommitteeHash = commitHash

	ph := util.DecodeHex(txt.Header.PrevBlockHash)
	var prevHash [32]byte
	copy(prevHash[:], ph)
	header.BlockHeaderBase.PrevHash = prevHash
	return header
}

func (t *TxBlockHeader) Serialize() []byte {
	h := t.ToProtoBuf()
	bytes, _ := proto.Marshal(h)
	return bytes
}

func (t *TxBlockHeader) ToProtoBuf() *protobuf.ProtoTxBlock_TxBlockHeader {
	protoTxBlockHeader := &protobuf.ProtoTxBlock_TxBlockHeader{}
	protoBlockHeaderBase := t.BlockHeaderBase.ToProtobuf()
	protoTxBlockHeader.Blockheaderbase = protoBlockHeaderBase

	protoTxBlockHeader.Gaslimit = t.GasLimit
	protoTxBlockHeader.Gasused = &t.GasUsed

	data := make([]byte, 0)
	data = UintToByteArray(data, 0, t.Rewards, 16)
	protoTxBlockHeader.Rewards = &protobuf.ByteArray{Data: data}

	protoTxBlockHeader.Blocknum = t.BlockNum

	hashset := &protobuf.ProtoTxBlock_TxBlockHashSet{
		Stateroothash:  t.HashSet.stateRootHash[:],
		Statedeltahash: t.HashSet.deltaHash[:],
		Mbinfohash:     t.HashSet.mbInfoHash[:],
	}
	protoTxBlockHeader.Hash = hashset

	protoTxBlockHeader.Numtxs = &t.NumTxs
	protoTxBlockHeader.Minerpubkey = &protobuf.ByteArray{Data: util.DecodeHex(t.MinerPubKey)}
	protoTxBlockHeader.Dsblocknum = t.DSBlockNum

	return protoTxBlockHeader
}

type TxBlockHeaderT struct {
	BlockNum       string
	CommitteeHash  string
	DSBlockNum     string
	GasLimit       string
	GasUsed        string
	MbInfoHash     string
	MinerPubKey    string
	NumMicroBlocks int
	NumTxns        uint32
	PrevBlockHash  string
	Rewards        string
	StateDeltaHash string
	StateRootHash  string
	Timestamp      string
	TxnFees        string
	Version        int
}
