package verifier

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/multisig"
	"github.com/Zilliqa/gozilliqa-sdk/util"
)

type Verifier struct {
	NumOfDsGuard int
}

func (v *Verifier) AggregatedPubKeyFromDsComm(dsComm *list.List, dsBlock *core.DsBlock) ([]byte, error) {
	pubKeys, err := v.generateDsCommArray(dsComm, dsBlock)
	if err != nil {
		return nil, err
	}
	aggregatedPubKey, err := multisig.AggregatedPubKey(pubKeys)
	if err != nil {
		return nil, err
	}
	return aggregatedPubKey, nil
}

func (v *Verifier) AggregatedPubKeyFromTxComm(dsComm *list.List, txBlock *core.TxBlock) ([]byte, error) {
	pubKeys, err := v.generateDsCommArray2(dsComm, txBlock)
	if err != nil {
		return nil, err
	}
	aggregatedPubKey, err := multisig.AggregatedPubKey(pubKeys)
	if err != nil {
		return nil, err
	}
	return aggregatedPubKey, nil
}

// abstract this two methods
func (v *Verifier) generateDsCommArray(dsComm *list.List, dsBlock *core.DsBlock) ([][]byte, error) {
	if dsComm.Len() != len(dsBlock.Cosigs.B2) {
		return nil, errors.New("ds list mismatch")
	}
	bitmap := dsBlock.Cosigs.B2
	quorum := len(bitmap) / 3 * 2
	trueCount := 0
	for _, signed := range bitmap {
		if signed {
			trueCount++
		}
	}
	if !(trueCount > quorum) {
		return nil, errors.New("quorum error")
	}
	var commKeys []string
	cursor := dsComm.Front()
	for cursor != nil {
		pair := cursor.Value.(core.PairOfNode)
		cursor = cursor.Next()
		commKeys = append(commKeys, pair.PubKey)
	}

	var pubKeys [][]byte
	for index, key := range commKeys {
		if bitmap[index] {
			pubKeys = append(pubKeys, util.DecodeHex(key))
		}
	}
	return pubKeys, nil
}

func (v *Verifier) generateDsCommArray2(dsComm *list.List, txBlock *core.TxBlock) ([][]byte, error) {
	if dsComm.Len() != len(txBlock.Cosigs.B2) {
		return nil, errors.New("ds list mismatch")
	}
	bitmap := txBlock.Cosigs.B2
	quorum := len(bitmap) / 3 * 2
	trueCount := 0
	for _, signed := range bitmap {
		if signed {
			trueCount++
		}
	}
	if !(trueCount > quorum) {
		return nil, errors.New("quorum error")
	}
	var commKeys []string
	cursor := dsComm.Front()
	for cursor != nil {
		pair := cursor.Value.(core.PairOfNode)
		cursor = cursor.Next()
		commKeys = append(commKeys, pair.PubKey)
	}

	var pubKeys [][]byte
	for index, key := range commKeys {
		if txBlock.Cosigs.B2[index] {
			pubKeys = append(pubKeys, util.DecodeHex(key))
		}
	}
	return pubKeys, nil
}

// 0. verify current ds block
// 2. generate next ds committee
// return new ds comm
func (v *Verifier) VerifyDsBlock(dsBlock *core.DsBlock, dsComm *list.List) (*list.List, error) {
	newDsComm, err2 := v.UpdateDSCommitteeComposition("", dsComm, dsBlock)
	if err2 != nil {
		return nil, err2
	}
	return newDsComm, nil
}

func (v *Verifier) VerifyTxBlock(txBlock *core.TxBlock, dsComm *list.List) error {
	aggregatedPubKey, err := v.AggregatedPubKeyFromTxComm(dsComm, txBlock)
	if err != nil {
		return err
	}
	r, s := txBlock.GetRandS()
	if !multisig.MultiVerify(aggregatedPubKey, txBlock.Serialize(), r, s) {
		msg := fmt.Sprintf("verify tx block %d error", txBlock.BlockHeader.BlockNum)
		return errors.New(msg)
	}
	return nil
}

func (v *Verifier) UpdateDSCommitteeComposition(selfKeyPub string, dsComm *list.List, dsBlock *core.DsBlock) (*list.List, error) {
	var dummy core.MinerInfoDSComm
	return v.updateDSCommitteeComposition(selfKeyPub, dsComm, dsBlock, dummy)
}

// inner type of dsComm is core.PairOfNode
func (v *Verifier) updateDSCommitteeComposition(selfKeyPub string, dsComm *list.List, dsBlock *core.DsBlock, info core.MinerInfoDSComm) (*list.List, error) {
	// 0. verify ds block first
	aggregatedPubKey, err := v.AggregatedPubKeyFromDsComm(dsComm, dsBlock)
	if err != nil {
		return nil, err
	}
	headerBytes := dsBlock.Serialize()
	r, s := dsBlock.GetRandS()

	if !multisig.MultiVerify(aggregatedPubKey, headerBytes, r, s) {
		msg := fmt.Sprintf("verify ds block %d error", dsBlock.BlockHeader.BlockNum)
		return nil, errors.New(msg)
	}

	// 1. get the map of all pow winners from the DS block
	winners := dsBlock.BlockHeader.PoWDSWinners
	numOfWinners := len(winners)

	// 2. get the array of all non-performant nodes to be removed
	removeDSNodePubkeys := dsBlock.BlockHeader.RemoveDSNodePubKeys

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
		if current != nil {
			dsComm.MoveToBack(current)
		}
	}

	// 4. add new winners
	for pubKey, peer := range winners {
		w := core.PairOfNode{
			PubKey: pubKey[2:],
			Peer:   peer,
		}
		// Place the current winner node's information in front of the DS Committee
		count := v.NumOfDsGuard
		cursor := dsComm.Front()
		for count > 0 {
			count--
			cursor = cursor.Next()
		}
		dsComm.InsertBefore(w, cursor)
	}

	// 5. remove one node for every winner, maintaining the size of the DS Committee
	for i := 0; i < numOfWinners; i++ {
		back := dsComm.Back()
		dsComm.Remove(back)
	}

	return dsComm, nil
}
