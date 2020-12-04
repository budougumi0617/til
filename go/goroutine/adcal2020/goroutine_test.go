package adcal2020

import (
	"testing"
)

func genTasks() []task {
	const n = 1000
	t := make([]task, 0, n)
	for i := 0; i < n; i++ {
		t = append(t, slowTask)
	}
	return t
}

func BenchmarkSequential(b *testing.B) {
	ts := genTasks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sequential(ts)
	}
}

func BenchmarkSimpleGoroutine(b *testing.B) {
	ts := genTasks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		simple(ts)
	}
}

func BenchmarkSmartGoroutine(b *testing.B) {
	ts := genTasks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		smart(ts)
	}
}

func BenchmarkMyWorkerPoolGoroutine(b *testing.B) {
	ts := genTasks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myWorkerPool(ts)
	}
}

func BenchmarkExampleWorkerPoolGoroutine(b *testing.B) {
	ts := genTasks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExampleWorkerPool(ts)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	ts := genTasks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sequential(ts)
	}
}
