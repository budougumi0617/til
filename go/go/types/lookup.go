package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const code = `
package p

type I interface{
	Hoge() string
}

type S struct {
}

func (s *S) Hoge() string{
	return ""
}

type Y struct {
	hoge string
}

func (y Y) Hoge() string{
	return ""
}
`

func main() {
	fset := token.NewFileSet()

	conf := types.Config{}

	f, err := parser.ParseFile(fset, "p", code, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	pkg, err := conf.Check("p", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	st := pkg.Scope().Lookup("S").Type()
	yt := pkg.Scope().Lookup("Y").Type()
	it := pkg.Scope().Lookup("I").Type()

	pst := types.NewPointer(st)

	if types.Implements(pst, it.Underlying().(*types.Interface)) {
		fmt.Println(pst, "implements", it)
	}
	if types.Implements(yt, it.Underlying().(*types.Interface)) {
		fmt.Println(yt, "implements", it)
	}
}

