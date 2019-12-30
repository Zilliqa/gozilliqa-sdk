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

	return &EventLogSubscriber{Ws:ws}
}

func (subscriber *EventLogSubscriber) Start() (error, chan error, chan []byte) {
	err := subscriber.Ws.Subscribe()
	if err != nil {
		return err, nil, nil
	}

	subscriber.Ws.Start()
	return nil, subscriber.Ws.Err, subscriber.Ws.Msg
}

