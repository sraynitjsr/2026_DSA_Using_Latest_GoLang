package main

// NOTE:
// This program allocates many short-lived objects to stress the garbage collector.
//
// To run it with Go 1.25's experimental "Green Tea" garbage collector, use:
//   GOEXPERIMENT=greenteagc go run main.go
//
// No source code changes are required; the GC behavior is controlled by the
// GOEXPERIMENT environment variable.

import (
	"fmt"
	"runtime"
	"time"
)

type Node struct {
	a, b, c, d int
}

func allocateLots(n int) {
	for i := 0; i < n; i++ {
		_ = &Node{
			a: i,
			b: i * 2,
			c: i * 3,
			d: i * 4,
		}
	}
}

func main() {
	start := time.Now()

	for i := 0; i < 100; i++ {
		allocateLots(1_000_000)
		runtime.GC()
	}

	fmt.Println("Elapsed:", time.Since(start))
}
