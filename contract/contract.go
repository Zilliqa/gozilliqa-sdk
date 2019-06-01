package contract

import (
	"encoding/json"
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
	VName string `json:"vname"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type State struct {
	Value
}

func (c *Contract) Deploy(params DeployParams, attempts, interval int) error {
	if c.Code == "" || c.Init == nil || len(c.Init) == 0 {
		return errors.New("Cannot deploy without code or initialisation parameters.")
	}

	data, err := json.Marshal(c.Init)

	if err != nil {
		return err
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
		Data:         string(data),
		Status:       0,
	}

	err2 := c.Singer.Sign(tx, *c.Provider)
	if err2 != nil {
		return err2
	}

	rsp := c.Provider.CreateTransaction(tx.ToTransactionPayload())

	if rsp.Error != nil {
		return errors.New(rsp.Error.Message)
	}

	result := rsp.Result.(map[string]interface{})
	hash := result["TranID"].(string)

	tx.TrackTx(hash, c.Provider)

	if tx.Status == transaction.Rejected {
		c.ContractStatus = Rejected
		return nil
	}

	c.ContractStatus = Deployed
	c.Address = GetAddressFromContract(tx)

	return nil

}

func (c *Contract) Call(transition string, args []Value, params CallParams, attempts, interval int) (error, *transaction.Transaction) {
	if c.Address == "" {
		_ = errors.New("Contract has not been deployed!")
	}

	data := Data{
		Tag:    transition,
		Params: args,
	}

	dataStr, _ := json.Marshal(data)

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
		Data:         strings.ReplaceAll(string(dataStr), "/\\", ""),
		Status:       0,
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
