package main

import (
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const (
	text1 = `
yaml:
  key: value
  array:
  - null_value:
  - boolean: false
  - integer: 1
  - alias: aliases are like variables
  - alias: aliases are like variables
`
	text2 = `
json:
  key: value
  array:
  - null_value:
  - boolean: true
  - integer: 2
  - alias: aliase is like variable
  - alias: aliases are like variables
`
)

func main() {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)

	fmt.Println(dmp.DiffPrettyText(diffs))
}
