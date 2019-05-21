package testing

import (
	"fmt"
	"strconv"
	"testing"
)

// Distinct returns distinct array.
func Distinct(args []string) []string {
	m := make(map[string]bool, len(args))
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

var benchmarks = []struct {
	count string
}{
	{
		count: "10",
	},
	{
		count: "100",
	},
	{
		count: "10000",
	},
}

func buildArray(c string) []string {
	var a []string
	n, _ := strconv.Atoi(c)
	for i := 0; i < n; i++ {
		e := "hoge" + c
		for j := 0; j < n; j++ {
			a = append(a, e)
		}
	}
	return a
}

// BenchmarkDistinct checks performance.
func BenchmarkDistinct(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.count, func(b *testing.B) {
			a := buildArray(bm.count)
			fmt.Printf("len(a) = %+v\n", len(a))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Distinct(a)
			}
		})
	}
}

// BenchmarkDistinctOld checks performance.
func BenchmarkDistinctOld(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.count, func(b *testing.B) {
			a := buildArray(bm.count)
			fmt.Printf("len(a) = %+v\n", len(a))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				DistinctOld(a)
			}
		})
	}
}
