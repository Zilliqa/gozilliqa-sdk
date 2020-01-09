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
	"net/url"
	"testing"
)

func TestBuildNewBlockSubscriber(t *testing.T) {
	u := url.URL{Scheme: "wss", Host: "dev-ws.zilliqa.com", Path: ""}
	subscriber := BuildNewBlockSubscriber(u)
	err, ec, msg := subscriber.Start()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	cancel := false

	for {
		if cancel {
			_ = subscriber.Ws.Close()
			break
		}
		select {
		case message := <-msg:
			fmt.Println("Get new message: ", string(message))

		case err := <-ec:
			fmt.Println("Get error: ", err.Error())
			cancel = true
		}

	}

}
