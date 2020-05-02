package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 任意の数だけのgoroutineを起動するスニペット
func TestGoroutine_parallelSameAsProcsCount(t *testing.T) {
	nums := [10]int{}
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	type token struct{}
	sem := make(chan token, runtime.GOMAXPROCS(0))

	fmt.Println("goroutine", runtime.GOMAXPROCS(0))
	var wg sync.WaitGroup
	ch := make(chan int, len(nums))
	wg.Add(1)
	for _, num := range nums {
		sem <- token{}
		// tokenがinsertできれば後続の処理が進む
		num := num // use a copy to avoid data races
		go func() {
			ch <- num
			// 終わったらトークンを消費するので、別のgoroutineがinsertができる
			<-sem
			if num == len(nums)-1 {
				close(ch)
				wg.Done()
			}
		}()

	}
	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}()
	wg.Wait()
}
