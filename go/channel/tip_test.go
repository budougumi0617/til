package channel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func hoge(ctx context.Context, c chan struct{}) error {
	select {
	case c <- struct{}{}:
	case <-time.After(1 * time.Second):
		fmt.Println("time up")
		return errors.New(
			"timeout before receiving done restore channel",
		)
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func TestSelect(t *testing.T) {
	c := make(chan struct{})
	go func() {
		err := hoge(context.TODO(), c)
		if err != nil {
			fmt.Println(err)
		}
	}()
	//<-c
	time.Sleep(2 * time.Second)
}
