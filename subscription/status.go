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

type SocketConnectT string

const (
	Ready     SocketConnectT = "ready"
	Connect   SocketConnectT = "connect"
	Error     SocketConnectT = "error"
	Close     SocketConnectT = "close"
	Reconnect SocketConnectT = "reconnect"
)

type SocketStateT string

const (
	SocketConnect SocketStateT = "socket_connect"
	SocketMessage SocketStateT = "socket_message"
	SocketReady   SocketStateT = "socket_ready"
	SocketClose   SocketStateT = "socket_close"
	SocketError   SocketStateT = "socket_error"
)

// message type pushed by server side and we can query with to server
type MessageTypeT string

const (
	NewBlock     MessageTypeT = "NewBlock"
	EventLog     MessageTypeT = "EventLog"
	Notification MessageTypeT = "Notification"
	UnSubscribe  MessageTypeT = "Unsubscribe"
)

type StatusType string

const (
	SubscribeNewBlock StatusType = "SubscribeNewBlock"
	SubscribeEventLog StatusType = "SubscribeEventLog"
)
