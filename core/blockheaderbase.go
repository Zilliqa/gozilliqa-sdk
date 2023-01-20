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

import "github.com/Zilliqa/gozilliqa-sdk/v3/protobuf"

type BlockHeaderBase struct {
	Version uint32
	// Hash for the committee that generated the block
	CommitteeHash [32]byte
	PrevHash      [32]byte
}

func (b *BlockHeaderBase) ToProtobuf() *protobuf.ProtoBlockHeaderBase {
	protoBlockHeaderBase := &protobuf.ProtoBlockHeaderBase{}
	protoBlockHeaderBase.Version = b.Version
	protoBlockHeaderBase.Committeehash = b.CommitteeHash[:]
	protoBlockHeaderBase.Prevhash = b.PrevHash[:]
	return protoBlockHeaderBase
}
