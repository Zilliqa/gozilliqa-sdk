package account

import (
	provider2 "github.com/FireStack-Lab/LaksaGo/provider"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"testing"
)

func TestWallet_SignWith(t *testing.T) {
	wallet := NewWallet()
	wallet.AddByPrivateKey("e19d05c5452598e24caad4a0d85a49146f7be089515c905ae6a19e8a578a6930")
	tx := &transaction.Transaction{}
	provider := provider2.NewProvider("https://dev-api.zilliqa.com/")
	wallet.SignWith(tx, "9bfec715a6bd658fcb62b0f8cc9bfa2ade71434a", *provider)
}
