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
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
)

func TestBuildEventLogSubscriber(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	u := url.URL{Scheme: "wss", Host: "dev-ws.zilliqa.com", Path: ""}
	addresses := []string{"0x2ce491a0fd9e318b39172258101b7c836da7449b", "0x167e3980e04eab1e89ff84523ae8c77e008932dc"}
	subscriber := BuildEventLogSubscriber(u, addresses)

	err, ec, msg := subscriber.Start()
	assert.Nil(t, err, err)

	for {
		select {
		case message := <-msg:
			fmt.Println("Get new message: ", string(message))

		case err := <-ec:
			fmt.Println("Get error: ", err.Error())
		}
	}
}
