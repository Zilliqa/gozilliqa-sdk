package subscription

import (
	"context"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/Zilliqa/gozilliqa-sdk/workpool"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Walker struct {
	Provider     *provider.Provider
	FromBlock    uint64
	ToBlock      uint64
	CurrentBlock uint64
	Address      string
	EventLogs    map[uint64]interface{}
	WorkerNumber int64
}

func NewWalker(p *provider.Provider, from, to uint64, address string, workerNumber int64) *Walker {
	eventLogs := make(map[uint64]interface{})
	return &Walker{
		Provider:     p,
		FromBlock:    from,
		ToBlock:      to,
		Address:      address,
		EventLogs:    eventLogs,
		WorkerNumber: workerNumber,
	}
}

type GetEventReceiptTask struct {
	Provider *provider.Provider
	Id       string
	Complete *Complete
	Walker   *Walker
	BlockNum uint64
}

type Complete struct {
	sync.Mutex
	Number int
}

func (t GetEventReceiptTask) UUID() string {
	return t.Id
}

func NewGetReceiptTask(tx string, provider2 *provider.Provider, c *Complete, w *Walker, b uint64) GetEventReceiptTask {
	return GetEventReceiptTask{
		Id:       tx,
		Provider: provider2,
		Complete: c,
		Walker:   w,
		BlockNum: b,
	}
}

func (t GetEventReceiptTask) Run() {
	t.Complete.Lock()
	defer t.Complete.Unlock()
	t.Complete.Number++
	rsp := t.Provider.GetTransaction(t.Id)
	resultMap := rsp.Result.(map[string]interface{})
	receipt := resultMap["receipt"].(map[string]interface{})
	success := receipt["success"]
	if success == nil || success.(bool) == false {
		return
	} else {
		eventLogs, ok := receipt["event_logs"]
		if ok {
			els := eventLogs.([]interface{})
			for _, el := range els {
				log := el.(map[string]interface{})
				addr, ok := log["address"]
				if ok && strings.Compare(strings.ToLower(addr.(string)), strings.ToLower(t.Walker.Address)) == 0 {
					t.Walker.EventLogs[t.BlockNum] = log
				}
			}
		}
	}
}

func (w *Walker) StartTraversalBlock() {
	for i := w.FromBlock; i < w.ToBlock; i++ {
		rsp := w.Provider.GetTransactionsForTxBlock(strconv.FormatUint(i, 10))
		if rsp.Error != nil {
			fmt.Println("tx for block ", i, " = ", rsp.Error)
		} else {
			txResult := rsp.Result.([]interface{})
			var txns []string

			// flat tx hash
			for _, txs := range txResult {
				if txs == nil {
					continue
				}
				txList := txs.([]interface{})
				if len(txList) > 0 {
					for _, tx := range txList {
						txns = append(txns, tx.(string))
					}
				} else {
					continue
				}
			}

			// get detail
			fmt.Println("tx for block ", i, " = ", txns)
			totalTxNumber := len(txns)
			complete := &Complete{
				Number: 0,
			}

			wp := workpool.NewWorkPool(w.WorkerNumber)
			quit := make(chan struct{})
			for _, tx := range txns {
				task := NewGetReceiptTask(tx, w.Provider, complete, w, i)
				wp.AddTask(task)
			}
			go func(q chan struct{}) {
				wp.Poll(context.TODO(), quit)
				for {
					time.Sleep(time.Second)
					if totalTxNumber == complete.Number {
						close(q)
						break
					}
				}
			}(quit)
		}
	}
}
