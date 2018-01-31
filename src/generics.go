package main

import "fmt"

// We can define the required interface needed for our function
// In this case, the type must convert to a float32
type Convertatble interface {
	conv() float32
}


// Define two example types which satisfy the interface
type MyInt struct {
	int
}

func (i *MyInt) conv() float32 {
	return float32(i.int)
}

type MyFloat struct {
	float32
}

func (f *MyFloat) conv() float32 {
	return f.float32
}


// Create a function that takes two convertable types
// Essentially this says, we don't care what type you pass,
// as long as it is convertable to a float32
// Here we think about what we need from a passed type
func genericAdd(a, b Convertatble) float32 {
	return a.conv() + b.conv()
}

func main() {
	a := MyInt{5}
	b := MyFloat{5.0}
	fmt.Println(genericAdd(&a, &b))
}
