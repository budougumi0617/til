package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type TokenType string

const (
	OP_ADD TokenType = "add"
	SPACE  TokenType = "space"
	NUM    TokenType = "number"
)

// Token is...
type Token struct {
	Value string
	Type  TokenType
}

// Token is...
type Token struct {
	Value string
	Type  TokenType
}

func parse(txt string) []Token {
	var ts []Token
	// ここで頑張る

	return ts
}

func TestParse(t *testing.T) {
	want := []Token{
		{Value: "1", Type: NUM},
		{Value: " ", Type: SPACE},
		{Value: "+", Type: SPACE},
		{Value: " ", Type: SPACE},
		{Value: "2", Type: NUM},
		{Value: " ", Type: SPACE},
		{Value: "+", Type: SPACE},
		{Value: " ", Type: SPACE},
		{Value: "3", Type: NUM},
	}
	got := parse("1 + 2 + 3")

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}
