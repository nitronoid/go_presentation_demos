package main

// Use "go get <link>" then import
import (
	"github.com/davecgh/go-spew/spew"
	"fmt"
)

type Car struct{
	name string
	id uint
}

func (c Car) Honk(){
	fmt.Println("Honk!")
}

// A normal, has-a relationship
type Driveway struct {
	parkedCar Car
}


// Below is struct embedding, which disguises composition as inheritance
// Using a non-pointer member has the usual inheritance traits
type NewCar struct{
	Car
}

// Using a pointer member has an alias relationship
type CarProxy struct{
	*Car
}


func main() {
	testCar := Car{"Junk", 42}

	// Access the parked car in the same way you would expect
	composition := Driveway{testCar}
	fmt.Println(composition.parkedCar.name)
	composition.parkedCar.Honk()

	// Access the internal car as though this is a car
	inheritanceish := NewCar{testCar}
	fmt.Println(inheritanceish.name)
	inheritanceish.Honk()

	// This behaves like the prior example however it affects the testCar
	proxy := CarProxy{&testCar}
	fmt.Println(testCar.name)
	proxy.name = "Garbage"
	fmt.Println(testCar.name)

	// We can create a collection of cars by still accessing all of the embedded structs
	cars := []Car{testCar, composition.parkedCar, inheritanceish.Car, *proxy.Car}
	spew.Printf("%v", cars)

}
