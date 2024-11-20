package main

import "fmt"

func whileLoop2(n int) int {
	res := 0
	i := 1
	for i <= n {
		res += i
		i++
		i *= 2
	}
	return res
}

func main() {
	fmt.Println(whileLoop2(10))
}
