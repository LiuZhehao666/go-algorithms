package main

import "fmt"

/*   while循环   */
func whileLoop(n int) int {
	res := 0
	i := 1
	for i <= n {
		res += i
		i++
	}
	return res
}
func main() {
	fmt.Println(whileLoop(10))
}
