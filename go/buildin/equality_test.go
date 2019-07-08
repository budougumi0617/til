package main

import (
	"testing"
)

type Tags *[]string

type Article struct {
	Title string
	Tag   Tags
}

// https://budougumi0617.github.io/2019/07/07/prevent-runtime-error-by-pointer/
func TestHelloWorld(t *testing.T) {
	var a1 interface{} = Article{
		Title: "Content Title",
		Tag: &[]string{
			"go",
			"slice",
		},
	}

	var a2 interface{} = Article{
		Title: "Content Title",
		Tag: &[]string{
			"go",
			"slice",
		},
	}
	if a1 == a2 {
		t.Fatal("Same article")
	}
}
