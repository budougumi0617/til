package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MindscapeHQ/raygun4go"
)

func parent() {
	child()
}

func child() {
	an := os.Getenv("RAYGUN_APP_NAME")
	key := os.Getenv("RAYGUN_API_KEY")
	cli, err := raygun4go.New(an, key)
	if err != nil {
		log.Fatalf("build client faield %v", err)
	}
	serr := fmt.Errorf("sample error")
	if err := cli.SendError(serr); err != nil {
		log.Fatalf("sendError failed %v", err)
	}

}

func main() {
	parent()
}
