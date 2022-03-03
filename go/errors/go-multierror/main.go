package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

// errorインターフェイスにはいくつかのエラーが含まれる可能性がある。
func hoo() error {
	g := multierror.Group{}
	g.Go(func() error {
		return errors.New("hoge")
	})
	g.Go(func() error {
		return errors.New("!!")
	})
	g.Go(func() error {
		return nil // エラー返さないこともある
	})
	return g.Wait()
}

func main() {
	err := hoo()
	// 普通のエラー処理
	if err != nil {
		fmt.Printf("エラー起きてた！\n%v\n", err)
	}

	fmt.Println("errorあったかもリスト")
	// マルチエラーが返ってきてたら中身見てみる。
	var merr *multierror.Error
	if errors.As(err, &merr) {
		// https://pkg.go.dev/github.com/hashicorp/go-multierror#Error.WrappedErrors
		for i, err := range merr.WrappedErrors() {
			fmt.Printf("No.%d: %v\n", i, err)
		}

	}
}
