package main

import (
	"fmt"
	"time"
)

func main () {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // use commas to separate multiple expression in the same case statement
		fmt.Println("It's the weekend", time.Now().Weekday())
	default:
		fmt.Println("It's the weekday", time.Now().Weekday())
	}

	t := time.Now()
	switch { // without an expression is an alternate way to express if/else logic
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int: 
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)

	whatAmI(1)

	whatAmI("hey")


}