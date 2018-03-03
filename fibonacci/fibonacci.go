package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	//"log"
)

const MaxMemoized = 60

func fib_recursive(n uint64) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib_recursive(n-1) + fib_recursive(n-2)
	}
}

func fib_linear(n uint64) uint64 {

	var a uint64 = 0
	var b uint64 = 1

	for i := uint64(0); i < n; i++ {
		temp := a
		a = b
		b = temp + b
	}
	return a
}

func fib_recursive_memoized(n int) uint64 {
	if n > MaxMemoized {
		return 0
	} else if n <= 1 {
		return uint64(n)
	} else {
		if m[n] > 0 {
			return m[n]
		} else {
			m[n] = fib_recursive_memoized(n-1) + fib_recursive_memoized(n-2)
			return m[n]
		}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("————————————————————————————————————————————\n%-20s took %s seconds\n\n", name, elapsed)
}

func run_non_recursive() {
	defer timeTrack(time.Now(), "run_non_recursive")
	for i := 1; i < len(os.Args); i++ {
		n, e := strconv.Atoi(os.Args[i])
		if e == nil {
			fmt.Println(uint64(n), fib_linear(uint64(n)))
		}
	}
}

func run_recursive() {
	defer timeTrack(time.Now(), "run_recursive")

	for i := 1; i < len(os.Args); i++ {
		n, e := strconv.Atoi(os.Args[i])
		if e == nil {
			fmt.Println(uint64(n), fib_recursive(uint64(n)))
		}
	}
}

func run_recursive_memo() {
	defer timeTrack(time.Now(), "run_recursive_memo")

	for i := 1; i < len(os.Args); i++ {
		n, e := strconv.Atoi(os.Args[i])
		if e == nil {
			fmt.Println(n, fib_recursive_memoized(n))
		}
	}
}

var m map[int]uint64

func main() {
	m = make(map[int]uint64)

	run_non_recursive()
	run_recursive_memo()
	run_recursive()
}
