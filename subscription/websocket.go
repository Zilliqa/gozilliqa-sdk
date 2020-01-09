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
	"github.com/gorilla/websocket"
	"net/url"
)

type Websocket struct {
	Topic  Topic
	URL    url.URL
	Err    chan error
	Msg    chan []byte
	Client *websocket.Conn
}

func NewWebsocket(topic Topic, url url.URL, err chan error, msg chan []byte) *Websocket {
	return &Websocket{
		Topic: topic,
		URL:   url,
		Err:   err,
		Msg:   msg,
	}
}

func (w *Websocket) Subscribe() error {
	c, _, err := websocket.DefaultDialer.Dial(w.URL.String(), nil)
	if err != nil {
		return err
	}
	w.Client = c

	sub, err := w.Topic.Stringify()
	if err != nil {
		return err
	}

	err2 := c.WriteMessage(websocket.TextMessage, sub)
	if err2 != nil {
		return err2
	}

	return nil
}

func (w *Websocket) Start() {
	go func() {
		for {
			_, message, err := w.Client.ReadMessage()
			if err != nil {
				w.Err <- err
				c, _, err := websocket.DefaultDialer.Dial(w.URL.String(), nil)
				if err != nil {
					w.Err <- err
				} else {
					w.Client = c
					sub, _ := w.Topic.Stringify()
					err2 := c.WriteMessage(websocket.TextMessage, sub)
					if err2 != nil {
						w.Err <- err2
					}
				}
			} else {
				w.Msg <- message
			}

		}
	}()
}

func (w *Websocket) Close() error {
	return w.Client.Close()
}
