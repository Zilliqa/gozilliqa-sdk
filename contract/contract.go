package contract

import (
	"errors"
	"github.com/FireStack-Lab/LaksaGo/account"
	"github.com/FireStack-Lab/LaksaGo/provider"
	"github.com/FireStack-Lab/LaksaGo/transaction"
	"strings"
)

type ContractStatus int

const (
	Deployed ContractStatus = iota
	Rejected
	Initialised
)

type Contract struct {
	Init           []Value        `json:"init"`
	Abi            string         `json:"abi"`
	State          State          `json:"state"`
	Address        string         `json:"address"`
	Code           string         `json:"code"`
	ContractStatus ContractStatus `json:"contractStatus"`

	Singer   *account.Wallet
	Provider *provider.Provider
}

type Value struct {
	VName string      `json:"vname"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type State struct {
	Value
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
		Receipt:      transaction.TransactionReceipt{},
		SenderPubKey: params.SenderPubKey,
		ToAddr:       "0000000000000000000000000000000000000000",
		Code:         strings.ReplaceAll(c.Code, "/\\", ""),
		Data:         c.Init,
		Status:       0,
	}

	err2 := c.Singer.Sign(tx, *c.Provider)
	if err2 != nil {
		return nil, err2
	}

	rsp := c.Provider.CreateTransaction(tx.ToTransactionPayload())

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

func (c *Contract) Call(transition string, args []Value, params CallParams, priority bool, attempts, interval int) (error, *transaction.Transaction) {
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
		Receipt:      transaction.TransactionReceipt{},
		SenderPubKey: params.SenderPubKey,
		ToAddr:       c.Address,
		Code:         strings.ReplaceAll(c.Code, "/\\", ""),
		Data:         data,
		Status:       0,
		Priority:     priority,
	}

	err2 := c.Singer.Sign(tx, *c.Provider)
	if err2 != nil {
		return err2, nil
	}

	rsp := c.Provider.CreateTransaction(tx.ToTransactionPayload())

	if rsp.Error != nil {
		return errors.New(rsp.Error.Message), nil
	}

	result := rsp.Result.(map[string]interface{})
	hash := result["TranID"].(string)
	tx.ID = hash

	if tx.Status == transaction.Rejected {
		c.ContractStatus = Rejected
		return nil, nil
	}

	return nil, tx

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
