package adcal2020

import (
	"testing"
)

func genTasks(n int) []task {
	t := make([]task, 0, n)
	for i := 0; i < n; i++ {
		t = append(t, slowTask)
	}
	return t
}

const N = 1000

func BenchmarkSimpleGoroutine(b *testing.B) {
	ts := genTasks(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		simple(ts)
	}
}

func BenchmarkSmartGroutine(b *testing.B) {
	ts := genTasks(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		smart(ts)
	}
}

func BenchmarkReuseGroutine(b *testing.B) {
	ts := genTasks(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reuse(ts)
	}
}
