package benchstat

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	ts := genTasks()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		execute(ts)
	}
}
