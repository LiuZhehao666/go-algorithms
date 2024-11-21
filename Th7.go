package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 2 {
		return n - 1
	}
	res := fib(n-1) + fib(n-2)
	return res
}

func main() {
	fmt.Println(fib(10))
}