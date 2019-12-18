package subscription

import (
	"fmt"
	provider2 "github.com/Zilliqa/gozilliqa-sdk/provider"
	"testing"
)

func TestWalker_TraversalBlock(t *testing.T) {
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")
	walker := NewWalker(provider, 933750, 933770, "0xab14b0fd133721d7c47ef410908e8ffc2b39167f",50,"Transfer")
	walker.StartTraversalBlock()
	fmt.Println(walker.EventLogs)
}
