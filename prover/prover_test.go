package prover

import (
	"testing"

	"github.com/Zilliqa/gozilliqa-sdk/v3/core"
	"github.com/Zilliqa/gozilliqa-sdk/v3/provider"
)

func TestStateProver_VerifyStateProof(t *testing.T) {
	if true {
		t.Skip("Skipping verifier test, since this depends on an old environment")
	}
	core.SkipIfCI(t)
	p := provider.NewProvider("https://mpt42-api.dev.z7a.xyz")
	sp := &StateProver{RpcClient: p}

	latestTxBlock, _ := sp.RpcClient.GetLatestTxBlock()

	contractAddr := "5050065f52bd935f9fe58937986f74373657a7fd"
	vname := "zilToPolyTxHashMap"
	indices := []string{"0"}
	blockNum := latestTxBlock.Header.BlockNum
	accountHash := latestTxBlock.Header.StateRootHash

	t.Log("current block number is: ", blockNum)
	t.Log("account root hash from tx block: ", accountHash)

	storageKey := core.GenerateStorageKey(contractAddr, vname, indices)
	stateProof, err := sp.RpcClient.GetStateProof(contractAddr, string(storageKey), &blockNum)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("state proof: ", stateProof)

	value, err2 := sp.VerifyStateProof(contractAddr, vname, indices, stateProof, accountHash)
	if err2 != nil {
		t.Fatal(err2)
	}

	t.Log(string(value))
}
