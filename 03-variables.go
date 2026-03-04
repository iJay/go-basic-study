package main

import "fmt"

// s := "apple" // syntax error: non-declaration statement outside function body

func main () {
	var a = "initial"
	fmt.Println(a)

	// var b, c int = 1, 'a' // no error!!!
	// var b, c int = 1, "a" // error:  cannot use "a" (untyped string constant) as int value in variable declaration
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // decalre and no initialize  the zero-valued for int is 0
	fmt.Println(e)

	f := "apple" // := is shorthand form var f string = "apple" , this syntax is only available inside functions
	fmt.Println(f)

	g := true
	fmt.Println(g)

	var h string
	fmt.Println(h) // the zero-valued for string is empty ""

	var i bool
	fmt.Println(i) // the zero-valued for bool is false
}