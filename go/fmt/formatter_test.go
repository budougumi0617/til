package fmt

import (
	"fmt"
	"testing"
)

type MyString string

func (ms MyString) String() string {
	return "return from String()"
}

func (ms MyString) GoString() string {
	return "return from GoString()"
}

type Root struct {
	RootField MyString
}

// Confirm interfaces in fmt package.
// Ref:
// fmt.Formatterを実装して%vや%+vをカスタマイズしたり、%3🍺みたいな書式をつくってみよう #golang
// https://qiita.com/tenntenn/items/453a09c4c6d7f580d0ab
func TestMyString(t *testing.T) {
	var ms MyString
	fmt.Println(ms)
	fmt.Printf("ms by %%s\t=\t%s\n", ms)
	fmt.Printf("ms by %%v\t=\t%+v\n", ms)
	fmt.Printf("ms by %%+v\t=\t%v\n", ms)
	fmt.Printf("ms by %%#v\t=\t%#v\n", ms)
}

func TestRoot(t *testing.T) {
	root := Root{}
	fmt.Println(root)
	fmt.Printf("root = %+v\n", root)
	fmt.Printf("root by %%s\t=\t%s\n", root)
	fmt.Printf("root by %%v\t=\t%+v\n", root)
	fmt.Printf("root by %%+v\t=\t%v\n", root)
	fmt.Printf("root by %%#v\t=\t%#v\n", root)
}

type Password string

func (p Password) String() string {
	rs := []rune(p)
	for i := 0; i < len(rs)-2; i++ {
		rs[i] = 'X'
	}
	return string(rs)
}

type Credential struct {
	ID       string
	Password Password
}

func TestPassword(t *testing.T) {
	cr := Credential{
		ID:       "budougumi0617",
		Password: "secret",
	}
	fmt.Println(cr)
	fmt.Printf("cr by %%s\t=\t%s\n", cr)
	fmt.Printf("cr by %%v\t=\t%+v\n", cr)
	fmt.Printf("cr by %%+v\t=\t%v\n", cr)
	fmt.Printf("cr by %%#v\t=\t%#v\n", cr)
}
