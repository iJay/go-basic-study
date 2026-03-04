package main

// Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

import (
	"fmt"
	"maps"
)

func main () {
	// create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)
	// Set key/value pairs name[key] = val
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	// _, prs := m["k2"] // return ture
	// fmt.Println("prs: ", prs)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Get a value for a key with name[key]
	// If the key doesn’t exist, the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	// len returns the number of key/value pairs
	fmt.Println("len: ", len(m))

	// delete removes key/value pairs from a map
	delete(m , "k2")
	fmt.Println("map: ", m)
	
	// deleting a non-existent key will not cause an error
	delete(m, "k4") 
	fmt.Println("map: ", m)

	// remove all key/value pairs from a map
	clear(m)
	fmt.Println("map: ", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map
	_, prs := m["k2"]
	// _ is a blank identifier, it is used to discard the value
	fmt.Println("prs: ", prs)

	// declare and initialize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n map: ", n)

	
	q := make(map[string]int)
	q["k1"] = 7
	fmt.Println("q map: ", q)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n and n2 are equal")
	} else {
		fmt.Println("n and n2 are not equal")
	}

	var r = n
	r["foo"] = 100
	fmt.Println("n map: ", n)
	fmt.Println("r map: ", r)
}