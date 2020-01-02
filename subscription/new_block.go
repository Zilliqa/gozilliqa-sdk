package subscription

import "net/url"

type NewBlockSubscriber struct {
	Ws *Websocket
}

func BuildNewBlockSubscriber(url url.URL) *NewBlockSubscriber {
	topic := &NewBlockQuery{Query: "NewBlock"}
	err := make(chan error, 1)
	msg := make(chan []byte, 10)
	ws := NewWebsocket(topic, url, err, msg)
	return &NewBlockSubscriber{Ws: ws}
}

func (subscriber *NewBlockSubscriber) Start() (error, chan error, chan []byte) {
	err := subscriber.Ws.Subscribe()
	if err != nil {
		return err, nil, nil
	}

	subscriber.Ws.Start()
	return nil, subscriber.Ws.Err, subscriber.Ws.Msg
}
