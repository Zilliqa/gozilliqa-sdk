package prover

import (
	"errors"
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/v3/core"
	"github.com/Zilliqa/gozilliqa-sdk/v3/mpt"
	"github.com/Zilliqa/gozilliqa-sdk/v3/provider"
	"github.com/Zilliqa/gozilliqa-sdk/v3/util"
)

type StateProver struct {
	RpcClient *provider.Provider
}

func (p *StateProver) VerifyStateProof(contractAddr string, vname string, indices []string, stateProof *core.StateProof, stateRootHash string) ([]byte, error) {
	var proof [][]byte
	for _, p := range stateProof.AccountProof {
		bytes := util.DecodeHex(p)
		proof = append(proof, bytes)
	}

	db := mpt.NewFromProof(proof)
	root := util.DecodeHex(stateRootHash)
	accountBaseBytes, err := mpt.Verify([]byte(contractAddr), db, root)
	if err != nil {
		return nil, err
	}

	accountBase, err2 := core.AccountBaseFromBytes(accountBaseBytes)
	if err2 != nil {
		msg := fmt.Sprintf("%s - %s", "parse accountBase error", err2.Error())
		return nil, errors.New(msg)
	}

	var proof2 [][]byte
	for _, p := range stateProof.StateProof {
		bytes := util.DecodeHex(p)
		proof2 = append(proof2, bytes)
	}

	db2 := mpt.NewFromProof(proof2)
	storageKey := core.GenerateStorageKey(contractAddr, vname, indices)
	hashedStorageKey := util.Sha256(storageKey)
	value, err3 := mpt.Verify([]byte((util.EncodeHex(hashedStorageKey))), db2, accountBase.StorageRoot)
	if err3 != nil {
		msg := fmt.Sprintf("%s - %s", "get value error", err3.Error())
		return nil, errors.New(msg)
	}

	return value, nil
}
