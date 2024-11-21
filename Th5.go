package main

import "fmt"

func recur(n int) int {
	if n == 1 {
		return 1
	}
	res := recur(n - 1)
	return res + n
}
func main() {
	fmt.Println(recur(10))
}
