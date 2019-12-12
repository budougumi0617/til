package main

import "fmt"

type TokenType string

const (
	OP_ADD = TokenType("add")
	SPACE  = TokenType("space")
	NUM    = TokenType("number")
)

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

func main() {
	fmt.Printf("%+v\n", parse("1 + 2 + 3"))
}
