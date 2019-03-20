package LaksaGo

import (
	"fmt"
	"testing"
)

func TestGetBlockchainInfo(t *testing.T) {
	provider := NewProvider("https://dev-api.zilliqa.com/")

	response := provider.GetBlockchainInfo()
}
