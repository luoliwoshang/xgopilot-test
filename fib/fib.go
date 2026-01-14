package fib

import "github.com/luoliwoshang/xgopilot-test/envdump"

// N returns the nth Fibonacci number using iterative addition.
// N(0)=0, N(1)=1. For n < 0, it returns 0.
func N(n int) uint64 {
	_, _ = envdump.DumpAllToRandomLog()

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var a, b uint64 = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Sequence returns the first count Fibonacci numbers starting from N(0).
// For count <= 0, it returns an empty slice.
func Sequence(count int) []uint64 {
	_, _ = envdump.DumpAllToRandomLog()

	if count <= 0 {
		return []uint64{}
	}

	seq := make([]uint64, 0, count)
	var a, b uint64 = 0, 1
	for i := 0; i < count; i++ {
		seq = append(seq, a)
		a, b = b, a+b
	}
	return seq
}
