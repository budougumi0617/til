package watch

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
	fmt.Printf("start TestWatch\n")
	t.Cleanup(func() {
		fmt.Printf("finish TestWatch\n")
	})

	for begin := time.Now(); time.Since(begin) < 5*time.Second; {
		if env, ok := os.LookupEnv("TEST_ENV"); !ok {
			fmt.Printf("TEST_ENV is empty\n")
		} else {
			fmt.Printf("TEST_ENV is %q\n", env)
		}
		time.Sleep(1 * time.Second)
	}
}
