package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuilder(t *testing.T) {
	want := "foobarhoge"
	in := []string{
		"foo",
		"bar",
		"hoge",
	}
	var sb strings.Builder
	// Reserve spaces.
	sb.Grow(40)
	for _, s := range in {
		fmt.Fprint(&sb, s)
	}

	got := sb.String()

	if got != want {
		t.Errorf("want %s, but got = %s\n", want, got)
	}
}

func TestUpdateFileName(t *testing.T) {
	mut := func(path string) string {
		// 文字列を最後の'/'で分割する
		parts := strings.LastIndex(path, "/")
		// プレフィックスを追加して元の文字列と結合
		return path[:parts+1] + "prefix_" + path[parts+1:]
	}

	in := "path/to/file.csv"
	want := "path/to/prefix_file.csv"

	if got := mut(in); got != want {
		t.Errorf("want %s, but got = %s\n", want, got)
	}
}
