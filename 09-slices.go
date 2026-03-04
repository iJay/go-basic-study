package main

import (
	"fmt"
	"slices"
)

// Slices are a important data type in Go

func main () {
	var s[] string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	// create a slice with make
	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	// len returns the length of the slice as expected.
	fmt.Println("len: ", len(s))

	// append returns a slice containing one or more new values
	s = append(s, "d") 
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// copy create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// slice[low: high] including s[2] but excluding s[5]
	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}