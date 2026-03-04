package main

import (
	"fmt"
	"strconv"
)

func plus(a int, b int) int { // Go requires explicit returns
	return a + b
}

// 同类型参数可以省略之前的type name 只在最后一个参数后写入type name即可
func plusPlus(a, b, c int) int {
	return a+ b + c
}

// 这里只是测试连续同类型参数的type传入
func stringPlus(a, b string, c int) string {
	return a + b + strconv.Itoa(c)
}

func main () {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	res2 := stringPlus("123", "456", 789)
	fmt.Println("res2: ", res2)
}