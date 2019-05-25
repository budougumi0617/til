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
BenchmarkDistinct/small-4                 500000              2067 ns/op              16 B/op          1 allocs/op
BenchmarkDistinct/large-4                   5000            208053 ns/op              16 B/op          1 allocs/op
BenchmarkDistinct/fuge-4                       1        1796819644 ns/op              16 B/op          1 allocs/op
BenchmarkDistinctOld/small-4              500000              2095 ns/op            1792 B/op          1 allocs/op
BenchmarkDistinctOld/large-4               10000            158002 ns/op          163840 B/op          1 allocs/op
BenchmarkDistinctOld/fuge-4                    1        2708319483 ns/op        1600004096 B/op        1 allocs/op
PASS
ok      github.com/budougumi0617/til/go/testing 54.878s
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
		var r []string
		a := algo(n)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r = d(a)
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
