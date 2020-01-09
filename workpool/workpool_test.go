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
