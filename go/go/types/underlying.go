package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	fset := token.NewFileSet()

	code := (`
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
}

func (y Y) Hoge() string{
	return ""
}

`)
	conf := types.Config{
		Importer: importer.Default(),
		Error: func(err error) {
			fmt.Printf("!!! %#v\n", err)
		},
	}

	f, err := parser.ParseFile(fset, "p", code, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}

	pkg, err := conf.Check("p", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal(err)
	}
	ast.Inspect(f, func(n ast.Node) bool {

		switch n.(type) {
		//case *ast.StructType, *ast.InterfaceType:
		case *ast.Ident:

			// 識別子ではない場合は無視
			expr, ok := n.(ast.Expr)
			if !ok {
				return true
			}

			typ := info.TypeOf(expr)
			if typ == nil {
				return true
			}
			_, oks := typ.Underlying().(*types.Struct)
			_, oki := typ.Underlying().(*types.Interface)

			if oki || oks {
				fmt.Printf("%v\n", typ)
				fmt.Println(fset.Position(expr.Pos()))
			}
		}

		return true
	})

	st := pkg.Scope().Lookup("S").Type().Underlying().(*types.Struct)
	it := pkg.Scope().Lookup("I").Type().Underlying().(*types.Interface)
	fmt.Printf("%#v\n", st)
	ptr := types.NewPointer(pkg.Scope().Lookup("S").Type())
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", it)

	if types.Implements(ptr, it) {
		fmt.Println(ptr, "implements", it)
	}
}

