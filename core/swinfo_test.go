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
	"testing"
)

func TestSWInfo_Serialize(t *testing.T) {
	swinfo := &SWInfo{
		ZilliqaMajorVersion: 7,
		ZilliqaMinorVersion: 1,
		ZilliqaFixVersion:   1,
		ZilliqaUpgradeDS:    0,
		ZilliqaCommit:       0,
		ScillaMajorVersion:  3,
		ScillaMinorVersion:  2,
		ScillaFixVersion:    0,
		ScillaUpgradeDS:     0,
		ScillaCommit:        1,
	}

	data := swinfo.Serialize()
	t.Log(util.EncodeHex(data))
	// 000000070000000100000001000000000000000000000000000000030000000200000000000000000000000000000001
	// todo assertion
}
