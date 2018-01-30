package main

import "fmt"

// This is like a copy of the type, not an alias
type myInt int

// Passing a regular int to this function won't work
// The int would need to be casted to myInt
func strict(arg myInt) {
	fmt.Println(arg)
}

// Basic struct type, instances can be constructed wherever this is visible
type basic struct{
	name string
	value int
}

func main() {

	// Instance of the basic type
	obj := basic{"test", 42}
	fmt.Println(obj)

	// struct for a pair of int's
	var pairi struct{first, second int}
	pairi = struct{first, second int}{5,7}
	fmt.Println(pairi)
	fmt.Println(pairi.first)

	empty := struct{}{}
	fmt.Println(empty)

	// We can use out of order, named intialisation of members
	var bar = struct{beer, spirit string}{
		spirit : "gin",
		beer : "fosters",
		}
	fmt.Println(bar)
}
