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

import "net/url"

type EventLogSubscriber struct {
	Ws *Websocket
}

func BuildEventLogSubscriber(url url.URL, addresses []string) *EventLogSubscriber {
	topic := &NewEventLogQuery{
		Query:     "EventLog",
		Addresses: addresses,
	}

	err := make(chan error, 1)
	msg := make(chan []byte, 10)
	ws := NewWebsocket(topic, url, err, msg)

	return &EventLogSubscriber{Ws: ws}
}

func (subscriber *EventLogSubscriber) Start() (error, chan error, chan []byte) {
	err := subscriber.Ws.Subscribe()
	if err != nil {
		return err, nil, nil
	}

	subscriber.Ws.Start()
	return nil, subscriber.Ws.Err, subscriber.Ws.Msg
}
