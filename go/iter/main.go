package main

import (
	"fmt"
)

func f(yield func(int) bool) {
	fmt.Println("print in f")
	yield(0)
}

func main() {
	for i := range f {
		fmt.Println("print in loop", i)
	}
}
