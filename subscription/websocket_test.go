package subscription

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestWebsocket_Start(t *testing.T) {
	u := url.URL{Scheme: "wss", Host: "dev-ws.zilliqa.com", Path: ""}
	log.Printf("connecting to %s", u.String())
	topic := &NewBlockQuery{Query: "NewBlock"}
	ws := &Websocket{
		Topic: topic,
		URL:   u,
		Err:   make(chan error, 1),
		Msg:   make(chan []byte, 10),
	}

	err := ws.Subscribe()
	if err != nil {
		fmt.Println("Init websocket failed: ", err)
	}

	ws.Start()
	for {
		select {
		case message := <-ws.Msg:
			fmt.Println("Get new message: ", string(message))

		case err := <-ws.Err:
			fmt.Println("Get error: ", err.Error())
		}

	}
}
