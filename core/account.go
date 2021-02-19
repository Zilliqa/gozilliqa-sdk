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
