package workpool

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

type WorkerPool struct {
	maxWorkers int64
	sem        *semaphore.Weighted
	takes      map[string]Task
	ids        []string
}

func NewWorkPool(max int64) *WorkerPool {
	m := semaphore.NewWeighted(max)
	ts := make(map[string]Task)
	ids := make([]string, 0)
	return &WorkerPool{
		maxWorkers: max,
		sem:        m,
		takes:      ts,
		ids:        ids,
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.ids = append(wp.ids, task.UUID())
	wp.takes[task.UUID()] = task
}

func (wp *WorkerPool) Top() Task {
	if len(wp.ids) == 0 {
		return nil
	}

	id := wp.ids[0]
	t := wp.takes[id]

	delete(wp.takes, id)
	wp.ids = wp.ids[1:]
	return t

}

func (wp *WorkerPool) Poll(ctx context.Context, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			fmt.Println("quit now..")
			break
		default:
			task := wp.Top()
			if task == nil {
				time.Sleep(time.Second * 3)
				fmt.Println("no task for now...")
			} else {
				if err := wp.sem.Acquire(ctx, 1); err != nil {
					break
				}
				go func() {
					defer wp.sem.Release(1)
					task.Run()
				}()
			}
		}

	}
}
