package main

import "fmt"

func fibo(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}
func main() {
	n := 6
	output := fibo(n)
	fmt.Printf("The %dth number in this Fibonacci sequence is: %d\n", n, output)
}