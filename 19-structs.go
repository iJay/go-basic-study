package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	// Go is a garbage collected language; 
	// you can safely return a pointer to a local variable
	// it will only be cleaned up by the garbage collector when there are no active references to it.
	p := person{name: name}
	p.age = 42
	return &p
}

func main () {
	fmt.Println(person{"Bob", 20})
	// cannot use 10 (untyped int constant) as string value in struct literal
	// cannot use "Bob2" (untyped string constant) as int value in struct literal
	// fmt.Println(person{10, "Bob2"})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"}) // Omitted fields will be zero-valued.s
	fmt.Println(&person{name: "Ann", age: 40}) // An & prefix yields a pointer to the struct.s
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	// Access struct fields with a dot.
	fmt.Println(s.name)

	sp := &s
	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	fmt.Println(sp.age)

	sp.age = 51
	// Structs are mutable.
	fmt.Println(s.age)

	dog := struct { //  The value can have an anonymous struct type
		name string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}