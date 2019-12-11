package subscription

import "github.com/Zilliqa/gozilliqa-sdk/provider"

type Walker struct {
	Provider  *provider.Provider
	FromBlock uint64
	ToBlock   uint64 // if 0, then forever
}
