// +build go1.12

package main

import "fmt"

func main() {
	// https://golang.org/pkg/go/build/#hdr-Build_Constraints
	// Go1.12以降でビルドするとこっちのmain.goになる。
	fmt.Println("build go1.12")
}
