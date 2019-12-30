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

func NewWebsocket(topic Topic, url url.URL,err chan error, msg chan []byte) *Websocket {
	return &Websocket{
		Topic:  topic,
		URL:    url,
		Err:    err,
		Msg:    msg,
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
				break
			}
			w.Msg <- message
		}
	}()
}

func (w *Websocket) Close() error {
	return w.Client.Close()
}
