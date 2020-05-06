package transaction

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionReceipt(t *testing.T) {
	js := []byte(`
		{
            "cumulative_gas": "1110",
            "epoch_num": "324",
            "event_logs": [
                {
                    "_eventname": "sorted",
                    "address": "0x297241320bfd1796efb1780e7a6732dfbe93220f",
                    "params": [
                        {
                            "type": "List (Uint32)",
                            "value": [
                                "1",
                                "2"
                            ],
                            "vname": "sorted"
                        }
                    ]
                },
                {
                    "_eventname": "unsorted",
                    "address": "0x297241320bfd1796efb1780e7a6732dfbe93220f",
                    "params": [
                        {
                            "type": "List (Uint32)",
                            "value": [
                                "2",
                                "1"
                            ],
                            "vname": "unsorted"
                        }
                    ]
                }
            ],
            "success": true
        }`)
	var receipt TransactionReceipt
	err := json.Unmarshal(js, &receipt)
	assert.Nil(t, err, err)
	t.Log(receipt)
}
