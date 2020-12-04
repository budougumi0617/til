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
