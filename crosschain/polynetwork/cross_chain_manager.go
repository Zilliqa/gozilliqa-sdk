/*
 * Copyright (C) 2020 Zilliqa
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package polynetwork

import (
	"errors"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/v3/account"
	"github.com/Zilliqa/gozilliqa-sdk/v3/bech32"
	"github.com/Zilliqa/gozilliqa-sdk/v3/contract"
	"github.com/Zilliqa/gozilliqa-sdk/v3/core"
	"github.com/Zilliqa/gozilliqa-sdk/v3/provider"
	"github.com/Zilliqa/gozilliqa-sdk/v3/transaction"
	"github.com/Zilliqa/gozilliqa-sdk/v3/util"
	"github.com/Zilliqa/gozilliqa-sdk/v3/validator"
	"strconv"
)

// Proxy for cross chain manager smart contract
type Proxy struct {
	ProxyAddr  string
	ImplAddr   string
	Wallet     *account.Wallet
	Client     *provider.Provider
	ChainId    int
	MsgVersion int
}

type ProofPair struct {
	Key  string
	Hash string
}

type ProofEntity struct {
	Proof string
	Pair  []ProofPair
}

func (p *Proxy) callWithNonce(args []core.ContractValue, transition string, nonce string) (*transaction.Transaction, error) {
	var bech32Addr string
	if validator.IsBech32(p.ProxyAddr) {
		bech32Addr = p.ProxyAddr
	} else {
		addr, err := bech32.ToBech32Address(p.ProxyAddr)
		if err != nil {
			return nil, err
		}
		bech32Addr = addr
	}

	gasPrice, err1 := p.Client.GetMinimumGasPrice()
	if err1 != nil {
		return nil, err1
	}

	c := contract.Contract{
		Address:  bech32Addr,
		Signer:   p.Wallet,
		Provider: p.Client,
	}

	params := contract.CallParams{
		Version:      strconv.FormatInt(int64(util.Pack(p.ChainId, p.MsgVersion)), 10),
		GasPrice:     gasPrice,
		GasLimit:     "40000",
		Amount:       "0",
		Nonce:        nonce,
		SenderPubKey: util.EncodeHex(p.Wallet.DefaultAccount.PublicKey),
	}

	tx, err2 := c.Call(transition, args, params, true)
	if err2 != nil {
		return nil, err2
	}

	return tx, nil
}

func (p *Proxy) call(args []core.ContractValue, transition string) (*transaction.Transaction, error) {
	var bech32Addr string
	if validator.IsBech32(p.ProxyAddr) {
		bech32Addr = p.ProxyAddr
	} else {
		addr, err := bech32.ToBech32Address(p.ProxyAddr)
		if err != nil {
			return nil, err
		}
		bech32Addr = addr
	}

	gasPrice, err1 := p.Client.GetMinimumGasPrice()
	if err1 != nil {
		return nil, err1
	}

	c := contract.Contract{
		Address:  bech32Addr,
		Signer:   p.Wallet,
		Provider: p.Client,
	}

	params := contract.CallParams{
		Version:      strconv.FormatInt(int64(util.Pack(p.ChainId, p.MsgVersion)), 10),
		GasPrice:     gasPrice,
		GasLimit:     "40000",
		Amount:       "0",
		SenderPubKey: util.EncodeHex(p.Wallet.DefaultAccount.PublicKey),
	}

	tx, err2 := c.Call(transition, args, params, true)
	if err2 != nil {
		return nil, err2
	}
	tx.Confirm(tx.ID, 1000, 10, p.Client)
	if tx.Status == core.Confirmed {
		return tx, nil
	}

	return nil, errors.New(fmt.Sprintf("call %s failed", transition))
}

func (p *Proxy) UpgradeTo() (*transaction.Transaction, error) {

	args := []core.ContractValue{
		{
			"new_crosschain_manager",
			"ByStr20",
			"0x" + p.ImplAddr,
		},
	}

	return p.call(args, "UpgradeTo")
}

func (p *Proxy) Unpause() (*transaction.Transaction, error) {
	args := []core.ContractValue{}

	return p.call(args, "UnPause")
}

func (p *Proxy) InitGenesisBlock(rawHeader string, pubKeys []string) (*transaction.Transaction, error) {
	var keys []core.ParamConstructor
	for _, key := range pubKeys {
		keys = append(keys, core.ParamConstructor{
			Constructor: "Polynetwork.Pubkey",
			ArgTypes:    make([]interface{}, 0),
			Arguments:   []interface{}{key},
		})
	}

	args := []core.ContractValue{
		{
			"rawHeader",
			"ByStr",
			rawHeader,
		},
		{
			"pubkeys",
			"List Polynetwork.Pubkey",
			keys,
		},
	}

	return p.call(args, "InitGenesisBlock")
}

func (p *Proxy) ChangeBookKeeper(rawHeader string, pubKeys []string, sigList []string) (*transaction.Transaction, error) {
	var keys []core.ParamConstructor
	for _, key := range pubKeys {
		keys = append(keys, core.ParamConstructor{
			Constructor: "Polynetwork.Pubkey",
			ArgTypes:    make([]interface{}, 0),
			Arguments:   []interface{}{key},
		})
	}

	var sigs []core.ParamConstructor
	for _, sig := range sigList {
		sigs = append(sigs, core.ParamConstructor{
			Constructor: "Polynetwork.Signature",
			ArgTypes:    make([]interface{}, 0),
			Arguments:   []interface{}{sig},
		})
	}
	args := []core.ContractValue{
		{
			"rawHeader",
			"ByStr",
			rawHeader,
		},
		{
			"pubkeys",
			"List Polynetwork.Pubkey",
			keys,
		},
		{
			"sigList",
			"List Polynetwork.Signature",
			sigs,
		},
	}

	return p.call(args, "ChangeBookKeeper")
}

func (p *Proxy) ChangeBookKeeperWithNonce(rawHeader string, pubKeys []string, sigList []string, nonce string) (*transaction.Transaction, error) {
	var keys []core.ParamConstructor
	for _, key := range pubKeys {
		keys = append(keys, core.ParamConstructor{
			Constructor: "Polynetwork.Pubkey",
			ArgTypes:    make([]interface{}, 0),
			Arguments:   []interface{}{key},
		})
	}

	var sigs []core.ParamConstructor
	for _, sig := range sigList {
		sigs = append(sigs, core.ParamConstructor{
			Constructor: "Polynetwork.Signature",
			ArgTypes:    make([]interface{}, 0),
			Arguments:   []interface{}{sig},
		})
	}
	args := []core.ContractValue{
		{
			"rawHeader",
			"ByStr",
			rawHeader,
		},
		{
			"pubkeys",
			"List Polynetwork.Pubkey",
			keys,
		},
		{
			"sigList",
			"List Polynetwork.Signature",
			sigs,
		},
	}

	return p.callWithNonce(args, "ChangeBookKeeper", nonce)
}

func (p *Proxy) VerifyHeaderAndExecuteTx(proof *ProofEntity, rawHeader string, headerProof *ProofEntity, curRawHeader string, headerSig []string) (*transaction.Transaction, error) {
	pairs := []core.ParamConstructor{}

	for _, pair := range proof.Pair {
		pairs = append(pairs, core.ParamConstructor{
			Constructor: "Pair",
			ArgTypes:    []interface{}{"ByStr1", "ByStr32"},
			Arguments:   []interface{}{"0x" + pair.Key, "0x" + pair.Hash},
		})
	}
	proofArguments := make([]interface{}, 0)
	proofArguments = append(proofArguments, "0x"+proof.Proof)
	proofArguments = append(proofArguments, pairs)
	proofConstructor := core.ParamConstructor{
		Constructor: "Polynetwork.Proof",
		ArgTypes:    make([]interface{}, 0),
		Arguments:   proofArguments,
	}

	headerPairs := []core.ParamConstructor{}
	for _, pair := range headerProof.Pair {
		headerPairs = append(headerPairs, core.ParamConstructor{
			Constructor: "Pair",
			ArgTypes:    []interface{}{"ByStr1", "ByStr32"},
			Arguments:   []interface{}{"0x" + pair.Key, "0x" + pair.Hash},
		})
	}

	headProofArguments := make([]interface{}, 0)
	headProofArguments = append(headProofArguments, "0x"+headerProof.Proof)
	headProofArguments = append(headProofArguments, headerPairs)
	headProofConstructor := core.ParamConstructor{
		Constructor: "Polynetwork.Proof",
		ArgTypes:    make([]interface{}, 0),
		Arguments:   headProofArguments,
	}

	var sigs []core.ParamConstructor
	for _, sig := range headerSig {
		sigs = append(sigs, core.ParamConstructor{
			Constructor: "Polynetwork.Signature",
			ArgTypes:    make([]interface{}, 0),
			Arguments:   []interface{}{sig},
		})
	}

	args := []core.ContractValue{
		{
			"proof",
			"Polynetwork.Proof",
			proofConstructor,
		},
		{
			"rawHeader",
			"ByStr",
			rawHeader,
		},
		{
			"headerProof",
			"Polynetwork.Proof",
			headProofConstructor,
		},
		{
			"curRawHeader",
			"ByStr",
			curRawHeader,
		},
		{
			"headerSig",
			"List Polynetwork.Signature",
			sigs,
		},
	}

	return p.call(args, "VerifyHeaderAndExecuteTx")
}

func (p *Proxy) VerifyHeaderAndExecuteTxWithNonce(proof *ProofEntity, rawHeader string, headerProof *ProofEntity, curRawHeader string, headerSig []string, nonce string) (*transaction.Transaction, error) {
	pairs := []core.ParamConstructor{}

	for _, pair := range proof.Pair {
		pairs = append(pairs, core.ParamConstructor{
			Constructor: "Pair",
			ArgTypes:    []interface{}{"ByStr1", "ByStr32"},
			Arguments:   []interface{}{"0x" + pair.Key, "0x" + pair.Hash},
		})
	}
	proofArguments := make([]interface{}, 0)
	proofArguments = append(proofArguments, "0x"+proof.Proof)
	proofArguments = append(proofArguments, pairs)
	proofConstructor := core.ParamConstructor{
		Constructor: "Polynetwork.Proof",
		ArgTypes:    make([]interface{}, 0),
		Arguments:   proofArguments,
	}

	headerPairs := []core.ParamConstructor{}
	for _, pair := range headerProof.Pair {
		headerPairs = append(headerPairs, core.ParamConstructor{
			Constructor: "Pair",
			ArgTypes:    []interface{}{"ByStr1", "ByStr32"},
			Arguments:   []interface{}{"0x" + pair.Key, "0x" + pair.Hash},
		})
	}

	headProofArguments := make([]interface{}, 0)
	headProofArguments = append(headProofArguments, "0x"+headerProof.Proof)
	headProofArguments = append(headProofArguments, headerPairs)
	headProofConstructor := core.ParamConstructor{
		Constructor: "Polynetwork.Proof",
		ArgTypes:    make([]interface{}, 0),
		Arguments:   headProofArguments,
	}

	var sigs []core.ParamConstructor
	for _, sig := range headerSig {
		sigs = append(sigs, core.ParamConstructor{
			Constructor: "Polynetwork.Signature",
			ArgTypes:    make([]interface{}, 0),
			Arguments:   []interface{}{sig},
		})
	}

	args := []core.ContractValue{
		{
			"proof",
			"Polynetwork.Proof",
			proofConstructor,
		},
		{
			"rawHeader",
			"ByStr",
			rawHeader,
		},
		{
			"headerProof",
			"Polynetwork.Proof",
			headProofConstructor,
		},
		{
			"curRawHeader",
			"ByStr",
			curRawHeader,
		},
		{
			"headerSig",
			"List Polynetwork.Signature",
			sigs,
		},
	}

	return p.callWithNonce(args, "VerifyHeaderAndExecuteTx", nonce)
}

func (p *Proxy) PopulateWhiteListFromContract(addr, val string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "addr",
			Type:  "ByStr20",
			Value: addr,
		}, {
			VName: "val",
			Type:  "Bool",
			Value: val,
		},
	}

	return p.call(args, "PopulateWhiteListFromContract")
}

func (p *Proxy) PopulateWhiteListToContract(addr, val string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "addr",
			Type:  "ByStr20",
			Value: addr,
		}, {
			VName: "val",
			Type:  "Bool",
			Value: val,
		},
	}

	return p.call(args, "PopulateWhiteListToContract")
}

func (p *Proxy) PopulateWhiteListMethod(method, val string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "method",
			Type:  "String",
			Value: method,
		}, {
			VName: "val",
			Type:  "Bool",
			Value: val,
		},
	}

	return p.call(args, "PopulateWhiteListMethod")
}

func (p *Proxy) PopulateConKeepersPublicKeyList(keepers []string) (*transaction.Transaction, error) {
	args := []core.ContractValue{{
		VName: "keepers",
		Type:  "List ByStr20",
		Value: keepers,
	}}

	return p.call(args, "PopulateConKeepersPublicKeyList")
}

func (p *Proxy) PopulateCurEpochStartHeight(height string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "height",
			Type:  "Uint32",
			Value: height,
		},
	}

	return p.call(args, "PopulateCurEpochStartHeight")
}

func (p *Proxy) PopulateZilToPolyTxHashMap(index, val string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "index",
			Type:  "Uint256",
			Value: index,
		},
		{
			VName: "val",
			Type:  "ByStr32",
			Value: val,
		},
	}

	return p.call(args, "PopulateZilToPolyTxHashMap")
}

func (p *Proxy) PopulateZilToPolyTxHashIndex(index string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "index",
			Type:  "Uint256",
			Value: index,
		},
	}

	return p.call(args, "PopulateZilToPolyTxHashIndex")
}

func (p *Proxy) PopulateFromChainTxExist(chainId, txId string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "chainId",
			Type:  "Uint64",
			Value: chainId,
		},
		{
			VName: "txId",
			Type:  "ByStr32",
			Value: txId,
		},
	}

	return p.call(args, "PopulateFromChainTxExist")
}

func (p *Proxy) PopulateFromChainTxExistWithNonce(chainId, txId, nonce string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			VName: "chainId",
			Type:  "Uint64",
			Value: chainId,
		},
		{
			VName: "txId",
			Type:  "ByStr32",
			Value: txId,
		},
	}

	return p.callWithNonce(args, "PopulateFromChainTxExist", nonce)
}
