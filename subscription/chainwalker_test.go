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
	provider2 "github.com/Zilliqa/gozilliqa-sdk/provider"
	"os"
	"testing"
)

func TestWalker_TraversalBlock(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	provider := provider2.NewProvider("https://api.zilliqa.com/")
	walker := NewWalker(provider, 613524, 613525, "0x367093f076490df47b9d7ec3400d8104649175d6", 50, "Burnt")
	walker.StartTraversalBlock()
	t.Log(walker.EventLogs)
}
