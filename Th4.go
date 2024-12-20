package main

import (
	"fmt"
)

/* 双层 for 循环 */
func nestedForLoop(n int) string {
	res := ""
	// 循环 i = 1, 2, ..., n-1, n
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			// 循环 j = 1, 2, ..., n-1, n
			res += fmt.Sprintf("(%d, %d), ", i, j)
			//res += fmt.Sprint("(", i, ", ", j, "), ")
		}
	}
	return res
}

func main() {
	fmt.Println(nestedForLoop(3))
}
