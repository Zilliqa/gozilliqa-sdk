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
