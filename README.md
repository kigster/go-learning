# go-learning

by Konstantin Gredeskoul, with many examples borrowed from around the Internet.

## Various examples used in learning Go

This repo is a learning project, aimed at keeping together various projects, and following a consistent structure.

## Examples

### Fibonacci

This little program offers three different implementations of computing a Fibonacci sum: 

1. A non-optimized recursive algorithm that gets really slow for N > 47
2. An optimized recursive algorithm that uses memoization (a map) to keep all pre-computed values. 
3. A linear algorithm that computes Fib sum without recursion.

**Running**

```bash
❯ go run fibonacci.go 0 1 10 40
0 0
1 1
10 55
40 102334155
————————————————————————————————————————————
run_non_recursive    took 41.121µs seconds

0 0
1 1
10 55
40 102334155
————————————————————————————————————————————
run_recursive_memo   took 22.354µs seconds

0 0
1 1
10 55
40 102334155
————————————————————————————————————————————
run_recursive        took 623.745ms seconds
```


