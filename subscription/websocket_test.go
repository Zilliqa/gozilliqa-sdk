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
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
)

func TestWebsocket_Start(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
	u := url.URL{Scheme: "wss", Host: "dev-ws.zilliqa.com", Path: ""}
	t.Logf("connecting to %s\n", u.String())
	topic := &NewBlockQuery{Query: "NewBlock"}
	ws := &Websocket{
		Topic: topic,
		URL:   u,
		Err:   make(chan error, 1),
		Msg:   make(chan []byte, 10),
	}

	err := ws.Subscribe()
	assert.Nil(t, err, err)

	ws.Start()
	for {
		select {
		case message := <-ws.Msg:
			t.Log("Get new message: ", string(message))

		case err := <-ws.Err:
			t.Log("Get error: ", err.Error())
		}

	}
}
