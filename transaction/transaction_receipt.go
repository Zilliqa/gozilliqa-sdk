package transaction

type TransactionReceipt struct {
	Success       bool `json:"success"`
	CumulativeGas string `json:"cumulative_gas"`
	EpochNum      string `json:"epoch_num"`
}
