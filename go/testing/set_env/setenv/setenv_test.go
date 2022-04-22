package setenv

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSetenv(t *testing.T) {
	fmt.Printf("start TestSetenv\n")
	t.Cleanup(func() {
		fmt.Printf("finish TestSetenv\n")
	})
	t.Setenv("TEST_ENV", "test")
	t.Run("sub", func(t *testing.T) {
		t.Setenv("TEST_ENV", "sub test")
		fmt.Printf("TEST_ENV is %q in sub\n", os.Getenv("TEST_ENV"))
	})
	fmt.Printf("TEST_ENV is %q in parent\n", os.Getenv("TEST_ENV"))
	time.Sleep(5 * time.Second)
}

func TestWatch(t *testing.T) {
	if env, ok := os.LookupEnv("TEST_ENV"); !ok {
		fmt.Printf("TEST_ENV is empty\n")
	} else {
		fmt.Printf("TEST_ENV is %q\n", env)
	}
}
