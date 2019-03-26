package main

import (
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestSETNXIfExists(t *testing.T) {
	c, err := redis.Dial("tcp", ":6379")
	defer c.Close()
	if err != nil {
		t.Fatal(err)
	}

	rep, err := c.Do("SET", []interface{}{"temp:test:0326", 10, "EX", 4000, "NX"}...)
	v, err := redis.String(rep, err)
	if err != nil {
		t.Fatalf("%#v\n", err) // redis_test.go:20: &errors.errorString{s:"redigo: nil returned"}
	}
	fmt.Println("result=", v)
}
