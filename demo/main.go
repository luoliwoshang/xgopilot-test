package main

import (
	"fmt"
	"log"

	"github.com/luoliwoshang/xgopilot-test/fib"
)

func main() {
	value := fib.N(10)
	seq := fib.Sequence(8)

	log.Printf("fib.N(10) = %d", value)
	fmt.Printf("fib.Sequence(8) = %v\n", seq)
}
