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

func TestAccountBaseFromBytes(t *testing.T) {
	bytes := util.DecodeHex("080112120a100000000000000000000000000000000018002220f74e858d851b7035161c66546fc183a5b162a8ee187d10324acb1fa8cf1391ea2a20f95d81f1e266a74b57e3bd6ec484ac9c1b2a006a23a6f3a911ce4cfe73ecd335")
	accountBase, err := AccountBaseFromBytes(bytes)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(util.EncodeHex(accountBase.StorageRoot))
}
