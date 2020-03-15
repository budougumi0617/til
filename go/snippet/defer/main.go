package main

import "fmt"

func run() (result string) {
	defer func() {
		result = "できます"
	}()

	s := "こんなことも"
	defer func(msg string) {
		fmt.Println(msg)
	}(s)

	s = "ただし"
	defer func() {
		fmt.Println(s)
	}()
	s = "deferを使うと"

	return "できません"
}

func main() {
	fmt.Println(run())
}
