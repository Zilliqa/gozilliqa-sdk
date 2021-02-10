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
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strconv"
)

type TxBlock struct {
	BlockBase
	BlockHeader *TxBlockHeader
}

func NewTxBlockFromTxBlockT(txt *TxBlockT) *TxBlock {
	txBlock := &TxBlock{}
	txBlockHeader := NewTxBlockHeaderFromTxBlockT(txt)
	txBlock.BlockHeader = txBlockHeader

	cs1Ser := util.DecodeHex(txt.Body.CS1)
	cs2Ser := util.DecodeHex(txt.Body.HeaderSign)

	cs1 := NewFromByteArray(cs1Ser)
	cs2 := NewFromByteArray(cs2Ser)

	cosig := CoSignatures{
		CS1: cs1,
		B1:  txt.Body.B1,
		CS2: cs2,
		B2:  txt.Body.B2,
	}
	txBlock.Cosigs = cosig
	timestamp, _ := strconv.ParseUint(txt.Header.Timestamp, 10, 64)
	txBlock.Timestamp = timestamp
	return txBlock
}

type TxBlockT struct {
	Header TxBlockHeaderT `json:"header"`
	Body   TxBlockBodyT   `json:"body"`
}

type TxBlockBodyT struct {
	B1              []bool
	B2              []bool
	BlockHash       string
	CS1             string
	HeaderSign      string
	MicroBlockInfos []MicroBlockInfo
}
