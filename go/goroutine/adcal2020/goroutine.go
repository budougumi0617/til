package adcal2020

import (
	"runtime"
	"sync"
)

type task func()

func fibonacciLoop(n int) int {
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

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func sequential(tasks []task) {
	for _, t := range tasks {
		t()
	}
}

func slowTask() {
	fibonacciLoop(1_000_000)
	//sqrt(20000000)
}

func simple(tasks []task) {
	var wg sync.WaitGroup
	for _, t := range tasks {
		wg.Add(1)
		go func(t task) {
			defer wg.Done()
			t()
		}(t)
	}
	wg.Wait()
}

func smart(tasks []task) {
	type token struct{}
	sem := make(chan token, runtime.GOMAXPROCS(0))

	for _, t := range tasks {
		sem <- token{}
		go func(t task) {
			t()
			<-sem
		}(t)
	}
}

func ExampleWorkerPool(tasks []task) {
	collector := StartDispatcher(runtime.GOMAXPROCS(0)) // start up worker pool
	for _, job := range tasks {
		collector.Work <- job
	}
}

func myWorkerPool(tasks []task) {
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
	for _, t := range tasks {
		queue <- t
		// fmt.Printf("enque! %d\n", i)
	}
	close(queue)
	wg.Wait()
}
