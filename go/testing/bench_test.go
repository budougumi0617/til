package testing

import (
	"strconv"
	"testing"
)

// Distinct returns distinct array.
func Distinct(args []string) []string {
	// m := make(map[string]bool, len(args))
	m := map[string]bool{}
	for _, a := range args {
		m[a] = true
	}
	as := make([]string, 0, len(m))
	for k := range m {
		as = append(as, k)
	}
	return as
}

// DistinctOld returns distinct array.
func DistinctOld(args []string) []string {
	results := make([]string, 0, len(args))
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}

// https://golang.org/pkg/testing/#hdr-Benchmarks

/*
$ make bench
go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/budougumi0617/til/go/testing
BenchmarkDistinct/DenseSmall-4            500000              2432 ns/op             483 B/op          2 allocs/op
BenchmarkDistinct/DenseLarge-4             10000            171177 ns/op            7447 B/op         10 allocs/op
BenchmarkDistinct/DenseFuge-4                  1        1581833876 ns/op          882304 B/op        190 allocs/op
BenchmarkDistinct/SparseSmall-4           200000             10322 ns/op            7447 B/op         10 allocs/op
BenchmarkDistinct/SparseLarge-4             2000           1076377 ns/op          879520 B/op        172 allocs/op
BenchmarkDistinct/SparseFuge-4                 1        55133349534 ns/op       7643668608 B/op  3903707 allocs/op
BenchmarkDistinctOld/DenseSmall-4         500000              3125 ns/op            2115 B/op          2 allocs/op
BenchmarkDistinctOld/DenseLarge-4           5000            298426 ns/op          169494 B/op         10 allocs/op
BenchmarkDistinctOld/DenseFuge-4               1        2420780895 ns/op        1600718560 B/op      165 allocs/op
BenchmarkDistinctOld/SparseSmall-4        100000             12156 ns/op            7447 B/op         10 allocs/op
BenchmarkDistinctOld/SparseLarge-4          1000           1493100 ns/op          879445 B/op        172 allocs/op
BenchmarkDistinctOld/SparseFuge-4              1        63013130864 ns/op       7643594368 B/op  3903243 allocs/op
PASS
ok      github.com/budougumi0617/til/go/testing 205.438s
*/
var benchmarks = []struct {
	count string
	n     int
	algo  func(int) []string
}{
	{
		count: "DenseSmall",
		n:     10,
		algo:  buildDenseArray,
	},
	{
		count: "DenseLarge",
		n:     100,
		algo:  buildDenseArray,
	},
	{
		count: "DenseFuge",
		n:     10000,
		algo:  buildDenseArray,
	},
	{
		count: "SparseSmall",
		n:     10,
		algo:  buildSparseArray,
	},
	{
		count: "SparseLarge",
		n:     100,
		algo:  buildSparseArray,
	},
	{
		count: "SparseFuge",
		n:     10000,
		algo:  buildSparseArray,
	},
}

func buildDenseArray(n int) []string {
	var a []string
	for i := 0; i < n; i++ {
		e := "hoge" + strconv.Itoa(i)
		for j := 0; j < n; j++ {
			a = append(a, e)
		}
	}
	return a
}

func buildSparseArray(n int) []string {
	var a []string
	nn := n * n
	for i := 0; i < nn; i++ {
		e := "hoge" + strconv.Itoa(i)
		a = append(a, e)
	}
	return a
}

func bench(n int, algo func(int) []string, d func([]string) []string) func(*testing.B) {
	return func(b *testing.B) {
		a := algo(n)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			d(a)
		}
	}
}

// BenchmarkDistinct checks performance.
func BenchmarkDistinct(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.count, bench(bm.n, bm.algo, Distinct))
	}
}

// BenchmarkDistinctOld checks performance.
func BenchmarkDistinctOld(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.count, bench(bm.n, bm.algo, DistinctOld))
	}
}
