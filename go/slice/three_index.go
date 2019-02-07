// https://play.golang.org/p/A5FuqHTD5Wj
package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // array(not slice)
	s := arr[2:5]                                 // slice operator

	fmt.Println("s =", s)        // s = [2 3 4]
	fmt.Println("len =", len(s)) // len = 3
	fmt.Println("cap =", cap(s)) // cap = 8

	s = arr[2:5:7]
	fmt.Println("s =", s)        // s = [2 3 4]
	fmt.Println("len =", len(s)) // len = 3
	fmt.Println("cap =", cap(s)) // cap = 5

	s = arr[:0]
	fmt.Println("s =", s)        // s = []
	fmt.Println("len =", len(s)) // len = 0
	fmt.Println("cap =", cap(s)) // cap = 10

	s = arr[:0:7]
	fmt.Println("s =", s)        // s = []
	fmt.Println("len =", len(s)) // len = 0
	fmt.Println("cap =", cap(s)) // cap = 7

	// s = arr[::7]              // middle index required in 3-index slice
	// s = arr[5:7:5]            // invalid slice index: 7 > 5

	s = arr[5:7:10]
	fmt.Println("s =", s)        // s = [5 6]
	fmt.Println("len =", len(s)) // len = 2
	fmt.Println("cap =", cap(s)) // cap = 5

	s = arr[1:8:9]
	fmt.Println("s =", s)        // s = [1 2 3 4 5 6 7]
	fmt.Println("len =", len(s)) // len = 7
	fmt.Println("cap =", cap(s)) // cap = 8

	s2 := s[2:4:5]
	fmt.Println("s2 =", s2)       // s2 = [3 4]
	fmt.Println("len =", len(s2)) // len = 2
	fmt.Println("cap =", cap(s2)) // cap = 3
}
