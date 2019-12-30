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

	for {
		select {
		case message := <-msg:
			fmt.Println("Get new message: ", string(message))

		case err := <-ec:
			fmt.Println("Get error: ", err.Error())
		}

	}

}
