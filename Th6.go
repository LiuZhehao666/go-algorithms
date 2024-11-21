package main

/* 尾递归的等价循环写法 */
import "fmt"

func loop(n int) int {
	res := 0
	for n > 0 {
		res += n
		n--
	}
	return res
}

func main() {
	fmt.Println(loop(10))
}
