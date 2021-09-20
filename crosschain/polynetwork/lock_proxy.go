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
	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
	"github.com/Zilliqa/gozilliqa-sdk/contract"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strconv"
)

type LockProxy struct {
	Addr       string
	Wallet     *account.Wallet
	Client     *provider.Provider
	ChainId    int
	MsgVersion int
}

func (l *LockProxy) BindAssetHash(fromAssetHash, toChainId, toAssetHash string) (*transaction.Transaction, error) {

	args := []core.ContractValue{
		{
			"fromAssetHash",
			"ByStr20",
			fromAssetHash,
		},
		{
			"toChainId",
			"Uint64",
			toChainId,
		},
		{
			"toAssetHash",
			"ByStr",
			toAssetHash,
		},
	}

	return l.call(args, "BindAssetHash", "0")
}

func (l *LockProxy) BindProxyHash(toChainId, targetProxyHash string) (*transaction.Transaction, error) {

	args := []core.ContractValue{
		{
			"toChainId",
			"Uint64",
			toChainId,
		},
		{
			"targetProxyHash",
			"ByStr",
			targetProxyHash,
		},
	}

	return l.call(args, "BindProxyHash", "0")
}

func (l *LockProxy) SetManager(newManager string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			"new_manager",
			"ByStr20",
			newManager,
		},
	}

	return l.call(args, "SetManager", "0")
}

func (l *LockProxy) SetManagerProxy(newManagerProxy string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			"new_manager_proxy",
			"ByStr20",
			newManagerProxy,
		},
	}

	return l.call(args, "SetManagerProxy", "0")
}

func (l *LockProxy) Lock(fromAssetHash, toChainId, toAddress, amount string) (*transaction.Transaction, error) {

	args := []core.ContractValue{
		{
			"fromAssetHash",
			"ByStr20",
			fromAssetHash,
		},
		{
			"toChainId",
			"Uint64",
			toChainId,
		},
		{
			"toAddress",
			"ByStr",
			toAddress,
		},
		{
			"amount",
			"Uint128",
			amount,
		},
	}

	return l.call(args, "lock", amount)
}

func (l *LockProxy) Unlock(txData, fromContractAddr, fromChainId string) (*transaction.Transaction, error) {

	args := []core.ContractValue{
		{
			"txData",
			"ByStr",
			txData,
		},
		{
			"fromContractAddr",
			"ByStr",
			fromContractAddr,
		},
		{
			"fromChainId",
			"Uint64",
			fromChainId,
		},
	}

	return l.call(args, "unlock", "0")
}

func (l *LockProxy) Pause() (*transaction.Transaction, error) {
	args := []core.ContractValue{}
	return l.call(args, "Pause", "0")
}

func (l *LockProxy) UnPause() (*transaction.Transaction, error) {
	args := []core.ContractValue{}
	return l.call(args, "UnPause", "0")
}

func (l *LockProxy) UpdateAdmin(newAdmin string) (*transaction.Transaction, error) {
	args := []core.ContractValue{
		{
			"newAdmin",
			"ByStr20",
			newAdmin,
		},
	}
	return l.call(args, "UpdateAdmin", "0")
}

func (l *LockProxy) ClaimAdmin() (*transaction.Transaction, error) {
	args := []core.ContractValue{}
	return l.call(args, "ClaimAdmin", "0")
}

func (l *LockProxy) call(args []core.ContractValue, transition string, amount string) (*transaction.Transaction, error) {
	bech32Addr, err := bech32.ToBech32Address(l.Addr)
	if err != nil {
		return nil, err
	}

	gasPrice, err1 := l.Client.GetMinimumGasPrice()
	if err1 != nil {
		return nil, err1
	}

	c := contract.Contract{
		Address:  bech32Addr,
		Signer:   l.Wallet,
		Provider: l.Client,
	}

	params := contract.CallParams{
		Version:      strconv.FormatInt(int64(util.Pack(l.ChainId, l.MsgVersion)), 10),
		GasPrice:     gasPrice,
		GasLimit:     "40000",
		Amount:       amount,
		SenderPubKey: util.EncodeHex(l.Wallet.DefaultAccount.PublicKey),
	}

	tx, err2 := c.Call(transition, args, params, true)
	if err2 != nil {
		return nil, err2
	}
	tx.Confirm(tx.ID, 1000, 10, l.Client)
	if tx.Status == core.Confirmed {
		return tx, nil
	}

	return nil, errors.New(fmt.Sprintf("call %s failed", transition))
}
