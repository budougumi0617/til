package main

import "fmt"

type Elem struct{ v int }

type Elems []Elem
type PElems []*Elem

// どうやっても呼び出し元はインクリメントできない
func incElem(es Elems) {
	for _, e := range es {
		e.v++
	}
}

func incElem2(es *Elems) {
	for _, e := range *es {
		// eはコピーなのでesの中身はインクリメントされない
		e.v++
	}
	// 冗長に書けば動く。
	for i := range *es {
		(*es)[i].v++
	}
}

func incPElem(es PElems) {
	for _, e := range es {
		e.v++
	}
}

func (es PElems) Min() *Elem {
	if len(es) == 0 {
		return nil
	}
	min := es[0]

	for i := 1; i < len(es); i++ {
		if es[i].v < min.v {
			min = es[i]
		}
	}

	return min
}

func run() {
	es := Elems{Elem{2}, Elem{3}}
	fmt.Println(es)
	incElem(es)
	fmt.Println(es)
	incElem2(&es)
	fmt.Println(es)

	pes := PElems{&Elem{2}, &Elem{3}, &Elem{1}}
	fmt.Println(*pes[0], *pes[1], *pes[2])
	incPElem(pes)
	fmt.Println(*pes[0], *pes[1], *pes[2])
	fmt.Println(pes.Min())
}
