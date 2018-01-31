package main

import "fmt"

type Animal interface{
	Noise() string
}

type Dog struct{}

//func (d Dog) Noise() string{
//	return "Pass by value! I mean Woof!"
//}

func (d *Dog) Noise() string{
	return "Pass by pointer! I mean Woof!"
}

// A function that takes an interface will base the contract,
// on how the functions have been implemented.
// If the interface function has been implemented with a pointer contract,
// a pointer must be provided.
// However if the function has a value contract, a pass by value is accepted,
// and a pointer is also accepted (automatically dereferenced)
func animalNoise(animal Animal){
	fmt.Println(animal.Noise())
}

func main() {
	dog := Dog{}
	animalNoise(&dog)
}
