/*
 * Copyright (C) 2019 Zilliqa
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
package contract

import (
	"errors"
	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/core"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strconv"
	"strings"
)

type ContractStatus int

const MainNet = "mainnet"
const TestNet = "testnet"
const Isolated = "isolated"
const TestNetHost = "https://dev-api.zilliqa.com/"
const MainNetHost = "https://api.zilliqa.com/"
const IsolatedHost = "https://zilliqa-isolated-server.zilliqa.com/"

const (
	Deployed ContractStatus = iota
	Rejected
	Initialised
)

type Contract struct {
	Init           []core.ContractValue `json:"init"`
	Abi            string               `json:"abi"`
	State          State                `json:"state"`
	Address        string               `json:"address"`
	Code           string               `json:"code"`
	ContractStatus ContractStatus       `json:"contractStatus"`

	Signer   *account.Wallet
	Provider *provider.Provider
}

type State struct {
	core.ContractValue
}

// This a shortcut function to deploy smart contract, all following parameters would remain default:
// network id, message version nonce, gasfee, gaslimit
// take a note that for gas limit, 40k is safe and recommend setting number
// value of network can be testnet, mainnet or isolated
func (c *Contract) DeployTo(network string) (*transaction.Transaction, error) {
	if network == TestNet {
		c.Provider = provider.NewProvider(TestNetHost)
		gasPrice, err := c.Provider.GetMinimumGasPrice()
		if err != nil {
			return nil, err
		}
		parameter := DeployParams{
			Version:      strconv.FormatInt(int64(util.Pack(333, 1)), 10),
			Nonce:        "",
			GasPrice:     gasPrice,
			GasLimit:     "40000",
			SenderPubKey: "",
		}
		return c.Deploy(parameter)
	} else if network == MainNet {
		c.Provider = provider.NewProvider(MainNetHost)
		gasPrice, err := c.Provider.GetMinimumGasPrice()
		if err != nil {
			return nil, err
		}
		parameter := DeployParams{
			Version:      strconv.FormatInt(int64(util.Pack(1, 1)), 10),
			Nonce:        "",
			GasPrice:     gasPrice,
			GasLimit:     "40000",
			SenderPubKey: "",
		}
		return c.Deploy(parameter)
	} else if network == Isolated {
		c.Provider = provider.NewProvider(IsolatedHost)
		gasPrice, err := c.Provider.GetMinimumGasPrice()
		if err != nil {
			return nil, err
		}
		parameter := DeployParams{
			Version:      strconv.FormatInt(int64(util.Pack(1, 1)), 10),
			Nonce:        "",
			GasPrice:     gasPrice,
			GasLimit:     "40000",
			SenderPubKey: "",
		}
		return c.Deploy(parameter)
	} else {
		return nil, errors.New("unsupported network, please use testnet, mainnet or isolated")
	}
}

func (c *Contract) Deploy(params DeployParams) (*transaction.Transaction, error) {
	if c.Code == "" || c.Init == nil || len(c.Init) == 0 {
		return nil, errors.New("Cannot deploy without code or initialisation parameters.")
	}

	tx := &transaction.Transaction{
		ID:           params.ID,
		Version:      params.Version,
		Nonce:        params.Nonce,
		Amount:       "0",
		GasPrice:     params.GasPrice,
		GasLimit:     params.GasLimit,
		Signature:    "",
		Receipt:      core.TransactionReceipt{},
		SenderPubKey: params.SenderPubKey,
		ToAddr:       "0000000000000000000000000000000000000000",
		Code:         strings.ReplaceAll(c.Code, "/\\", ""),
		Data:         c.Init,
		Status:       0,
	}

	err2 := c.Signer.Sign(tx, *c.Provider)
	if err2 != nil {
		return nil, err2
	}

	rsp, err := c.Provider.CreateTransaction(tx.ToTransactionPayload())

	if err != nil {
		return nil, err
	}

	if rsp == nil {
		return nil, errors.New("rpc response is nil")
	}

	if rsp.Error != nil {
		return nil, errors.New(rsp.Error.Message)
	}

	result := rsp.Result.(map[string]interface{})
	hash := result["TranID"].(string)
	contractAddress := result["ContractAddress"].(string)

	tx.ID = hash
	tx.ContractAddress = contractAddress
	return tx, nil

}
func (c *Contract) Sign(transition string, args []core.ContractValue, params CallParams, priority bool) (error, *transaction.Transaction) {
	if c.Address == "" {
		_ = errors.New("Contract has not been deployed!")
	}

	data := Data{
		Tag:    transition,
		Params: args,
	}

	tx := &transaction.Transaction{
		ID:           params.ID,
		Version:      params.Version,
		Nonce:        params.Nonce,
		Amount:       params.Amount,
		GasPrice:     params.GasPrice,
		GasLimit:     params.GasLimit,
		Signature:    "",
		Receipt:      core.TransactionReceipt{},
		SenderPubKey: params.SenderPubKey,
		ToAddr:       c.Address,
		Code:         strings.ReplaceAll(c.Code, "/\\", ""),
		Data:         data,
		Status:       0,
		Priority:     priority,
	}

	err2 := c.Signer.Sign(tx, *c.Provider)
	if err2 != nil {
		return err2, nil
	}

	return nil, tx
}

func (c *Contract) CallFor(transition string, args []core.ContractValue, priority bool, amount string, network string) (*transaction.Transaction, error) {
	if network == TestNet {
		c.Provider = provider.NewProvider(TestNetHost)
		gasPrice, err := c.Provider.GetMinimumGasPrice()
		if err != nil {
			return nil, err
		}
		params := CallParams{
			Version:      strconv.FormatInt(int64(util.Pack(333, 1)), 10),
			Nonce:        "",
			GasPrice:     gasPrice,
			GasLimit:     "40000",
			Amount:       amount,
			SenderPubKey: "",
		}
		return c.Call(transition, args, params, priority)
	} else if network == MainNet {
		c.Provider = provider.NewProvider(MainNetHost)
		gasPrice, err := c.Provider.GetMinimumGasPrice()
		if err != nil {
			return nil, err
		}
		params := CallParams{
			Version:      strconv.FormatInt(int64(util.Pack(1, 1)), 10),
			Nonce:        "",
			GasPrice:     gasPrice,
			GasLimit:     "40000",
			Amount:       amount,
			SenderPubKey: "",
		}
		return c.Call(transition, args, params, priority)
	} else if network == Isolated {
		c.Provider = provider.NewProvider(IsolatedHost)
		gasPrice, err := c.Provider.GetMinimumGasPrice()
		if err != nil {
			return nil, err
		}
		params := CallParams{
			Version:      strconv.FormatInt(int64(util.Pack(1, 1)), 10),
			Nonce:        "",
			GasPrice:     gasPrice,
			GasLimit:     "40000",
			Amount:       amount,
			SenderPubKey: "",
		}
		return c.Call(transition, args, params, priority)
	} else {
		return nil, errors.New("unsupported network, please use testnet, mainnet or isolated")
	}

}

func (c *Contract) Call(transition string, args []core.ContractValue, params CallParams, priority bool) (*transaction.Transaction, error) {
	if c.Address == "" {
		_ = errors.New("Contract has not been deployed!")
	}

	data := Data{
		Tag:    transition,
		Params: args,
	}

	tx := &transaction.Transaction{
		ID:           params.ID,
		Version:      params.Version,
		Nonce:        params.Nonce,
		Amount:       params.Amount,
		GasPrice:     params.GasPrice,
		GasLimit:     params.GasLimit,
		Signature:    "",
		Receipt:      core.TransactionReceipt{},
		SenderPubKey: params.SenderPubKey,
		ToAddr:       c.Address,
		Code:         strings.ReplaceAll(c.Code, "/\\", ""),
		Data:         data,
		Status:       0,
		Priority:     priority,
	}

	err2 := c.Signer.Sign(tx, *c.Provider)
	if err2 != nil {
		return tx, err2
	}
	rsp, err := c.Provider.CreateTransaction(tx.ToTransactionPayload())

	if err != nil {
		return tx, err
	}

	if rsp == nil {
		return tx, errors.New("rpc response is nil")
	}

	if rsp.Error != nil {
		return tx, errors.New(rsp.Error.Message)
	}

	result := rsp.Result.(map[string]interface{})
	hash := result["TranID"].(string)
	tx.ID = hash

	if tx.Status == core.Rejected {
		c.ContractStatus = Rejected
		return tx, nil
	}

	return tx, nil

}

func (c *Contract) IsInitialised() bool {
	return c.ContractStatus == Initialised
}

func (c *Contract) IsDeployed() bool {
	return c.ContractStatus == Deployed
}

func (c *Contract) IsRejected() bool {
	return c.ContractStatus == Rejected
}
