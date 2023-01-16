/*
* Copyright (C) 2020 Zilliqa
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

package polynetwork

import (
	"github.com/Zilliqa/gozilliqa-sdk/v3/util"
	"strings"
)

func DeserializeProof(proof string, position uint64) *ProofEntity {
	if strings.HasPrefix(proof, "0x") {
		proof = proof[2:]
	}

	proofBytes := util.DecodeHex(proof)
	zsc := ZeroCopySource{
		s:   proofBytes,
		off: position,
	}

	// todo handle eof
	value, _ := zsc.NextVarBytes()

	fullLen := uint64(len(proofBytes))
	llen := fullLen - zsc.off

	// ByStr1 + Bystr32
	compSize := uint64(33)
	size := llen / compSize

	var pairs []ProofPair
	for i := uint64(0); i < size; i++ {
		key, _ := zsc.NextByte()
		hash, _ := zsc.NextHash()
		pairs = append(pairs, ProofPair{
			Key:  util.EncodeHex([]byte{key}),
			Hash: util.EncodeHex(hash[:]),
		})
	}

	return &ProofEntity{
		Proof: util.EncodeHex(value),
		Pair:  pairs,
	}

}
