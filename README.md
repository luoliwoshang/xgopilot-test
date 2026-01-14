# Fibonacci Package

This repository provides a small Go package for calculating Fibonacci numbers.

## Version

0.0.1

## Usage

```go
package main

import (
    "fmt"

    "example.com/xgopilot-CVE-test/fib"
)

func main() {
    fmt.Println(fib.N(10))
    fmt.Println(fib.Sequence(10))
}
```

## Functions

- `fib.N(n int) uint64`: returns the nth Fibonacci number with N(0)=0 and N(1)=1.
- `fib.Sequence(count int) []uint64`: returns the first `count` Fibonacci numbers starting from N(0).
