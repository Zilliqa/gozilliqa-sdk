package verifier

import (
	"container/list"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
)

type Verifier struct {
	RpcClient    *provider.Provider
	NumOfDsGuard int
}

func (v *Verifier) Verify(dsBlockNum string, dsComm *list.List) (*list.List, error) {
	dst, err := v.RpcClient.GetDsBlockVerbose(dsBlockNum)
	if err != nil {
		return nil, err
	}
	dsBlock := core.NewDsBlockFromDsBlockT(dst)

	newDsComm := v.UpdateDSCommitteeComposition("", dsComm, dsBlock)
	//todo perform verify
	return newDsComm, nil
}

func (v *Verifier) UpdateDSCommitteeComposition(selfKeyPub string, dsComm *list.List, dsBlock *core.DsBlock) *list.List {
	var dummy core.MinerInfoDSComm
	return v.updateDSCommitteeComposition(selfKeyPub, dsComm, dsBlock, dummy)
}

// inner type of dsComm is PairOfNode
func (v *Verifier) updateDSCommitteeComposition(selfKeyPub string, dsComm *list.List, dsBlock *core.DsBlock, info core.MinerInfoDSComm) *list.List {
	// 1. get the map of all pow winners from the DS block
	winners := dsBlock.BlockHeader.PoWDSWinners
	numOfWinners := len(winners)
	fmt.Println("winners: ", winners)

	// 2. get the array of all non-performant nodes to be removed
	removeDSNodePubkeys := dsBlock.BlockHeader.RemoveDSNodePubKeys
	fmt.Println("removed keys: ", removeDSNodePubkeys)

	// 3. shuffle the non-performant nodes to the back
	for _, removed := range removeDSNodePubkeys {
		current := dsComm.Front()
		for current != nil {
			pairOfNode := current.Value.(core.PairOfNode)
			if pairOfNode.PubKey == removed {
				break
			}
			current = current.Next()
		}
		dsComm.MoveToBack(current)
	}

	// 4. add new winners
	for _, winner := range winners {
		dsComm.PushFront(winner)
	}

	// 5. remove one node for every winner, maintaining the size of the DS Committee
	for i := 0; i < numOfWinners; i++ {
		back := dsComm.Back()
		dsComm.Remove(back)
	}

	return dsComm
}
