package main

import "fmt"

// We can define the required interface needed for our function
// In this case, the type must convert to a float32
type Convertible interface {
	convert() float32
}


// Define two example types which satisfy the interface
type MyInt int

func (i *MyInt) convert() float32 {
	return float32(*i)
}

type MyFloat float32

func (f *MyFloat) convert() float32 {
	return float32(*f)
}


// Create a function that takes two convertible types
// Essentially this says, we don't care what type you pass,
// as long as it is convertible to a float32
// Here we think about what we need from a passed type
func genericAdd(a, b Convertible) float32 {
	return a.convert() + b.convert()
}

func main() {
	// These are casts not constructions
	a := MyInt(5)
	b := MyFloat(5.0)
	fmt.Println(genericAdd(&a, &b))
	// Regular int
	// Error the type int doesn't satisfy the interface
	//c := 5
	// fmt.Println(genericAdd(&c, &b))
}
