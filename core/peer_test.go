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
	"math/big"
	"testing"
)

func TestPeer_Serialize(t *testing.T) {
	ip := new(big.Int).SetUint64(16777343)
	peer := Peer{
		IpAddress:      ip,
		ListenPortHost: 0,
	}

	data := peer.Serialize()
	if util.EncodeHex(data) != "0000000000000000000000000100007F00000000" {
		t.Failed()
	}

	ip = new(big.Int).SetUint64(0)
	peer = Peer{
		IpAddress:      ip,
		ListenPortHost: 0,
	}
	data = peer.Serialize()
	if util.EncodeHex(data) != "0000000000000000000000000000000000000000" {
		t.Failed()
	}
}
