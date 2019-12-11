package subscription

import (
	"fmt"
	provider2 "github.com/Zilliqa/gozilliqa-sdk/provider"
	"testing"
)

func TestWalker_TraversalBlock(t *testing.T) {
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")
	walker := NewWalker(provider, 960215, 960420, "0x557b6d169b6dafedd2cbe3f7db486c1ef40d1ff3",50)
	walker.StartTraversalBlock()
	fmt.Println(walker.EventLogs)
}
