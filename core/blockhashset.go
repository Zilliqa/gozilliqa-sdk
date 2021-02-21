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

type DSBlockHashSet struct {
	// should be 32 bytes
	ShadingHash   []byte
	ReservedField [128]byte
}

type TxBlockHashSet struct {
	// State merkle tree root hash only valid in vacuous epoch
	// should be 32 bytes as well
	StateRootHash [32]byte
	// State Delta Hash on DS
	// 32 bytes
	DeltaHash [32]byte
	// Hash concatenated from all microblock infos
	// 32 bytes
	MbInfoHash [32]byte
}
