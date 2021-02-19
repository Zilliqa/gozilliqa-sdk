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

	cs1 := make([]byte, 0)
	cs2 := make([]byte, 0)

	cosig := &protobuf.ProtoBlockBase_CoSignatures{
		Cs1: &protobuf.ByteArray{
			Data: b.Cosigs.CS1.Serialize(cs1, 0),
		},
		B1: b.Cosigs.B1,
		Cs2: &protobuf.ByteArray{
			Data: b.Cosigs.CS2.Serialize(cs2, 0),
		},
		B2: b.Cosigs.B2,
	}
	blockBase.Cosigs = cosig

	return blockBase
}
