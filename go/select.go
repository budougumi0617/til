package main

import (
	"fmt"
	"time"
)

func main() {
	sample2()
}

func sample2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})

	f := func() int {
		fmt.Println("called f")
		return 1
	}

	go func() {
		fmt.Printf("f return = %d\n", <-ch1)
		done <- struct{}{}
	}()

	for {
		select {
		case ch1 <- f():
		case ch2 <- f():
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func sample1() {
	ch := make(chan int)
	cont := make(chan struct{})
	done := make(chan struct{})
	var cnt int

	f := func() int {
		cnt++
		fmt.Printf("called f() %d\n", cnt)
		return cnt
	}

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case ch <- f():
			case <-cont:
				fmt.Println("continue")
			}
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second")
	cont <- struct{}{}
	fmt.Printf("f return = %d\n", <-ch)
	fmt.Println("after <- ch")
	done <- struct{}{}
	fmt.Println("after done <- struct{}{}")

}
