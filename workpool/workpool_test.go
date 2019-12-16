package workpool

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

type FakeTask struct {
	Id string
}

func (t FakeTask) UUID() string {
	return t.Id
}

func NewFakeTask() FakeTask {
	return FakeTask{
		Id: uuid.New().String(),
	}
}

func (t FakeTask) Run() {
	fmt.Printf("task %s begin..\n", t.UUID())
	time.Sleep(time.Second * 3)
	fmt.Printf("task %s done...\n", t.UUID())
}

func TestNewWorkPool(t *testing.T) {
	quit := make(chan struct{})
	wp := NewWorkPool(10)
	for i := 0; i < 10; i++ {
		task := NewFakeTask()
		wp.AddTask(task)
	}

	go func() {
		wp.Poll(context.TODO(), quit)
	}()

	time.Sleep(time.Second * 10)
	close(quit)
}
