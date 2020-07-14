// https://play.golang.org/p/Drl1FjauMeM
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("ReadSliceは戻り値がbufferのスライスなのでいつの間にか中身が変わってしまう")
	run(func(r *bufio.Reader) ([]byte, error) { return r.ReadSlice('g') })
	fmt.Println("\nReadBytes使ったとき")
	run(func(r *bufio.Reader) ([]byte, error) { return r.ReadBytes('g') })
}

func run(read func(*bufio.Reader) ([]byte, error)) {
	buf := bytes.NewBufferString("original_original_original message")
	r := bufio.NewReaderSize(buf, 16)

	s, err := read(r) // orig
	if err != nil {
		panic(err)
	}
	fmt.Printf("before %q\n", s)

	// bufが上書きされるまでリードする
	_, err = read(r)
	if err != nil {
		panic(err)
	}
	_, err = read(r)
	if err != nil {
		panic(err)
	}

	// sはいじってないはずだけれど…
	fmt.Printf("after%q\n", s)
}
