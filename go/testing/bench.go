package testing

import "testing"

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
	name  string
	array []string
}{
	{
		name:  "Simple",
		array: []string{"hoge", "hoge2", "hoge3", "hoge4", "hoge5"},
	},
	{
		name:  "Dup",
		array: []string{"hoge", "hoge2", "hoge3", "hoge4", "hoge5", "hoge", "hoge2", "hoge3", "hoge4", "hoge5"},
	},
}

// BenchmarkDistinct checks performance.
func BenchmarkDistinct(b *testing.B) {
	// b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Distinct(bm.array)
			}
		})
	}
}

// BenchmarkDistinctOld checks performance.
func BenchmarkDistinctOld(b *testing.B) {
	// b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DistinctOld(bm.array)
			}
		})
	}
}
