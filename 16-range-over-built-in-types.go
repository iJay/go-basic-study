package main

import "fmt"

func main () {

	// 当range后是一个整数 只产生一个值，即当前的循环次数
	for n := range 10 {
		fmt.Println("num:", n)
	}

	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a":"apple", "b": "banana"}
	// range on map iterates over key/value pairs.
	for k, v := range kvs{
		// %s：代表 String（字符串）
		fmt.Printf("%s => %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go to zoo" {
		fmt.Println(i, c)
	}
}