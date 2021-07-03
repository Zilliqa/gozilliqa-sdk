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

import "math/big"

type SWInfo struct {
	ZilliqaMajorVersion uint32
	ZilliqaMinorVersion uint32
	ZilliqaFixVersion   uint32
	ZilliqaUpgradeDS    uint64
	ZilliqaCommit       uint32
	ScillaMajorVersion  uint32
	ScillaMinorVersion  uint32
	ScillaFixVersion    uint32
	ScillaUpgradeDS     uint64
	ScillaCommit        uint32
}

func (sw *SWInfo) Serialize() []byte {
	// length should be 48
	data := make([]byte, 48)
	UintToByteArray(data, 0, new(big.Int).SetUint64(uint64(sw.ZilliqaMajorVersion)), 4)
	UintToByteArray(data, 4, new(big.Int).SetUint64(uint64(sw.ZilliqaMinorVersion)), 4)
	UintToByteArray(data, 8, new(big.Int).SetUint64(uint64(sw.ZilliqaFixVersion)), 4)
	UintToByteArray(data, 12, new(big.Int).SetUint64(uint64(sw.ZilliqaUpgradeDS)), 8)
	UintToByteArray(data, 20, new(big.Int).SetUint64(uint64(sw.ZilliqaCommit)), 4)
	UintToByteArray(data, 24, new(big.Int).SetUint64(uint64(sw.ScillaMajorVersion)), 4)
	UintToByteArray(data, 28, new(big.Int).SetUint64(uint64(sw.ScillaMinorVersion)), 4)
	UintToByteArray(data, 32, new(big.Int).SetUint64(uint64(sw.ScillaFixVersion)), 4)
	UintToByteArray(data, 36, new(big.Int).SetUint64(uint64(sw.ScillaUpgradeDS)), 8)
	UintToByteArray(data, 44, new(big.Int).SetUint64(uint64(sw.ScillaCommit)), 4)
	return data
}

type SWInfoT struct {
	Scilla  []interface{}
	Zilliqa []interface{}
}
