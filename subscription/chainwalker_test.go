/*
 * Copyright (C) 2019 Zilliqa
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
package subscription

import (
	"fmt"
	provider2 "github.com/Zilliqa/gozilliqa-sdk/provider"
	"testing"
)

func TestWalker_TraversalBlock(t *testing.T) {
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")
	walker := NewWalker(provider, 933750, 933770, "0xab14b0fd133721d7c47ef410908e8ffc2b39167f",50,"Transfer")
	walker.StartTraversalBlock()
	fmt.Println(walker.EventLogs)
}
