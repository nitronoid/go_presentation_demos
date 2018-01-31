package main

import "fmt"

// Interfaces strictly only have functions
// No implementations
type Animal interface{
	Noise() string
}

type Dog struct{
	noise string
}

// Now Dog satisfies the animal interface
func (d Dog) Noise() string{
	return d.noise
}

// Provides the safety of a const member function
// All changes made to the object are local to this function
func (d Dog) immutable(new string){
	d.noise = new
}

// This is allowed to modify the original object as we have used a pointer
func (d *Dog) mutable(new string){
	d.noise = new
}



func main() {
	var animal Animal
	// No new statement required, just take the address of the object (gets GC'd)
	animal = &Dog{"Woof!"}

	// Note that this is the same as (*animal).Noise(), no -> operator
	fmt.Println(animal.Noise())

	dog := Dog{"original"}
	// This affects a copy so shouldn't modify dog
	dog.immutable("not original")
	fmt.Println(dog.Noise())
	// This affects and modifies the original dog
	dog.mutable("not original")
	fmt.Println(dog.Noise())
}
