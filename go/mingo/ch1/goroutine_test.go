package main

import (
	"fmt"
	"testing"
	"time"
)

// 1.4 リスト3 並列タスクを実行時にgoroutineリークを防ぐ
func work() {
	workers := 5
	ch := make(chan *Task, workers)
	// 全部終わる前にcloseしちゃうけど…
	defer close(ch)
	for i := 0; i < workers; i++ {
		go func() {
			for task := range ch {
				task.DoSomething()
			}
		}()
	}
	for i := 0; i < 20; i++ {
		ch <- &Task{}
	}
}

// Task is sample struct.
type Task struct{}

// DoSomething is slow method.
func (t *Task) DoSomething() {
	time.Sleep(1 * time.Second)
	fmt.Println("Do")
}

func TestHelloWorld(t *testing.T) {
	work()
}
