package subscription

import (
	"fmt"
	"net/url"
	"testing"
)

func TestBuildEventLogSubscriber(t *testing.T) {
	u := url.URL{Scheme: "wss", Host: "dev-ws.zilliqa.com", Path: ""}
	addresses := []string{"0x2ce491a0fd9e318b39172258101b7c836da7449b", "0x167e3980e04eab1e89ff84523ae8c77e008932dc"}
	subscriber := BuildEventLogSubscriber(u, addresses)

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
