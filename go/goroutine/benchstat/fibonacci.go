package benchstat

import (
	"runtime"
	"sync"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	f1, f0 := 1, 0
	fn := f1 + f0
	for i := n; i >= 2; i-- {
		fn = f1 + f0
		f1, f0 = fn, f1
	}
	return fn
}

func fibonacciTask() int {
	return fibonacci(1_000_000) // million
}

type task func() int

func genTasks() []task {
	const n = 1_000
	ts := make([]task, 0, n)
	for i := 0; i < n; i++ {
		ts = append(ts, fibonacciTask)
	}
	return ts
}

func execute(ts []task) {
	n := runtime.GOMAXPROCS(0)
	queue := make(chan task, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case t, ok := <-queue:
					if !ok {
						return
					}
					t()
				}
			}
		}()
	}
	for _, t := range ts {
		queue <- t
	}
	close(queue)
	wg.Wait()
}
