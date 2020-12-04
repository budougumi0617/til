package benchstat

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	f1, f0 := 1, 0
	fn := f1 + f0
	for i := n; i >= 2; i-- {
		fn = f1 + f0
		f1, f0 = fn, f1
	}
	return fn
}

func fibonacciTask() int {
	return fibonacci(1_000_000) // million
}

type task func() int

func genTasks() []task {
	const n = 1_000
	ts := make([]task, 0, n)
	for i := 0; i < n; i++ {
		ts = append(ts, fibonacciTask)
	}
	return ts
}

func execute(ts []task) {
	for _, t := range ts {
		t()
	}
}
