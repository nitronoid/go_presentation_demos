
<span class="menu-title" style="display: none">Intro</span>
# Go presentation 
### Jack Diver
#### SDAGE 2nd year 

---

<span class="menu-title" style="display: none">The go programming language</span>
# Go overview

- Go was built with simplicity in mind, for general purpose and systems programming in mind. 
- It is a compiled, statically and strongly typed language. 
- It supports concurrency and garbage collection out of the box. 
- Go does not support inheritance in the traditional sense.

---

<span class="menu-title" style="display: none">The history of go</span>
# History of Go

Go development was started in 2007 and announced in 2009. It was developed by Robert Griesemer, Rob Pike and Ken Thompson at google who all  disliked C++'s complexity. They wanted to create a language that removed common gripes with languages such as C++, Java, python etc... but keep the strong points such as being readable ("light on the page"), being scalable to large systems, and supporting networking. Go's tools are all open source.

---

<span class="menu-title" style="display: none">Basic types in go</span>
### Basic types in Go

- Go has automatic type deduction for variables, but also allows them to be explicitly typed. This is important in some cases.
- Types are specified after the variable name.
- Go is a strictly typed language, with no implicit conversions.

+++

<span class='menu-title slide-title'>types.go</span>
```golang
package main

import (
	"fmt"
	"reflect"
	)

// Global variables
var g_a int = 2
var g_b string = "str"
var g_c bool = true
var g_d rune = 'c'
var g_e uintptr = 0x0
var g_f *int = &g_a


// Declare multiple variables through one statement
var (
	y = "multiple"
	z = "variables"
)

// Same for constants
const (
	zero = 0
	pi = 3.14
)

func main() {
	// Auto determine type
	var a = 5
	// Declare (requires type) and then assign
	var b int
	b = 3
	// Var keyword isn't needed for local variables
	c := a + b
	// Print is in the fmt package
	fmt.Println(c, zero, pi)
	// Adds a space between outputs
	fmt.Println(y, z)

	// Scope works as you would expect
	{
		d := 2
		fmt.Println(d)
	}
	// fmt.Println(d) //ERROR

	// This is a float64, not 32
	e := 0.5
	fmt.Println(reflect.TypeOf(e))
	
	// types in go must exactly match the required types
	// there is no implicit conversion
	var num32 int32
	var numDef int
	//var test int64 = numDef + num32 //fails due to type mismatch
	var num64 int64 = int64(int32(numDef) + num32)
	fmt.Println(num64)

}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="8-14">Global variable definitions</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="17-21">We can declare multiple variables of different type through one statement</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="23-27">The same is done for constants, where const replaces var</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="30-36">Here we compare three different syntax used for declaring variables</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="31">Automatically determined type</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="33-34">Declaration requires type and will be given a zero value. Then we can assign later.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="36">The var keyword can be omitted completely for local variables if we use the := syntax.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="49-51">Using type deduction with floating point numbers yeilds a double precision variable. There are no float32 literals.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="53-58">Go has very strict typing with no implicit conversions, even between integral types. All conversions must be explicit.</span>

---

<span class="menu-title" style="display: none">Functions in go</span>
### Functions in Go

- Go functions return types are again specified at the end of the declaration before the function body.
- There is no void in Go, we can simply omit the return type.
- Go supports:
    * Returning multiple values
    * Variadic arguments
    * Named return values
    * Parameter lists that share type
- Go doesn't support function overloading or generics.

+++

<span class='menu-title slide-title'>functions.go</span>
```golang
package main

import "fmt"

// No overloading
func addi(a, b int) int {
	return a + b
}

func addf(a, b float32) float32 {
	return a + b
}

// Variadic args
func sum(args ...int) int {
	total := 0
	// range iterates over elements of containers, giving an index and element
	// here we use _ to say that we don't require the index
	for _, num := range args {
		total += num
	}
	return total
}

func multiReturn(a, b int) (int, int) {
	return a*2, b*2
}

// We can explicitly name our return value
func namedMult(a, b float32) (result float32) {
	result = a * b
	return // This is still necessary
}

func main() {
	// Variadic args
	fmt.Println(sum(3, 2, 1))

	// We can convert a slice to args with ...
	args := []int{3, 2, 1}
	fmt.Println(sum(args...))

	first, second := multiReturn(1, 2)
	fmt.Println(first, second)

	fmt.Println(namedMult(5, 6))
}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="5-12">No overloading.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="14-23">Variadic arguments.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="25-27">Multiple return values.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="29-33">Named return value.</span>

---

<span class="menu-title" style="display: none">Collections in go</span>
### Collections in Go

- Go has several built-in collection types:
    * Arrays (string acts similarly)
    * Slices
    * Maps
- A slice is the ```std::vector``` of Go

+++

<span class='menu-title slide-title'>collections.go</span>
```golang
package main

import "fmt"

func main() {
	// Arrays are a fixed size
	var arrI [3]int
	// Intialised to zero
	fmt.Println(arrI)

	// Auto type deduction
	arrStr := [4]string{"This", "is", "an", "array"}
	fmt.Println(arrStr)

	// Slices are dynamic
	sliceI := []int{0, 1, 2}
	fmt.Println("Length:", len(sliceI), "Elements:", sliceI)

	// This either modifies sliceI in place, or creates a new slice when it is too big
	sliceI = append(sliceI, 3, 4, 5)
	fmt.Println("Length:", len(sliceI), "Elements:", sliceI)

	// Deleting is more tricky
	i := 2
	// Delete and preserve order
	sliceI = append(sliceI[:i], sliceI[i+1:]...)
	fmt.Println("Length:", len(sliceI), "Elements:", sliceI)

	// Delete without preserving order by assigning the last
	sliceI[i] = sliceI[len(sliceI)-1]
	sliceI = sliceI[:len(sliceI)-1]
	fmt.Println("Length:", len(sliceI), "Elements:", sliceI)

	// When a slice contains pointers we should use
	// sliceI[len(sliceI)-1] = nil // Before removing

	// Creates a slice of 3 empty strings
	sliceStr := make([]string, 3)
	fmt.Println("Length:", len(sliceStr), "Elements:", sliceStr)

	// Map with string keys and int values
	var mapI map[string]int
	mapI = make(map[string]int)
	mapI["str"] = 12

	// Supports fairly complex types as keys by default
	mapIArr := make(map[[2]int]string)

	// Use complex key
	key := [2]int{0,1}
	mapIArr[key] = "str"

	// Get the data
	fmt.Println(mapI[mapIArr[key]])

	deduced := map[string][]int {
		"men":  {32, 55, 12, 55, 42, 53},
		"women":{44, 42, 23, 41, 65, 44},
	}

	// Test if the key exists and get the value
	val, hasKey := deduced["women"]
	fmt.Println("Has key:", hasKey, " With value:", val)

	// Remove an item from the list
	delete(deduced, "women")

	// Check again
	val2, hasKey2 := deduced["women"]
	fmt.Println("Has key:", hasKey2, " With value:", val2)

	// Adding to a map is simple, just assign to a new key
	deduced["kids"] = []int{43,2}
	fmt.Println(deduced)
}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="6-13">Arrays have a fixed size, given in the declaration.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="15-17">Slices have a similar syntax but omit the length.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="19-21">Slices can easily be appended to. If the slice grows too large for it's currently allocated memory, a new one is returned, hence the assignment.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="23-27">Deleting from a slice is strangely complicated. We must cut the slice down to before the element we want to remove, then append all the elements after it.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="29-32">If the order is unimportant we can assign from back and trim the slice by one.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="34-35">A pitfall is that slices containing pointers won't get GC'd by this, as the original slice may still exist and reference it. We need to delete the references.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="37-39">We are provided with a make function to construct collections of a given size. These elements will have a zero value.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="41-44">Maps can be declared, but will give you an error if you attempt to assign to them before using the make function.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="47">The make function requires the type anyway so I always use this syntax.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="47-54">Maps can have fairly complex types as their key, here an array is used.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="56-59">Maps can be created with data like other collections.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="62-63">When accessing a value from a map, you can also test for the key.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="66">Deleting from a Map is far simpler than from a slice.</span>

---

<span class="menu-title" style="display: none">Structs and user defined types in go</span>
### Structs and user defined types in Go

- Go defines types with the type keyword similar to typedef.
- These can be copies of existing types or newly defined structs.
- Constructors and destructors don't exist, we use free functions to create new objects if setup is non-trivial.

+++

<span class='menu-title slide-title'>structs_types.go</span>
```golang
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
type basic struct {
	name string
	value int
}

// This is perfectly legal, go performs "pointer escape analysis"
func newBasic(name string, value int) *basic {
	ret := basic{name, value}
	//return &basic{name, value} //also valid
	return &ret
}

func main() {

	// Instance of the basic type
	obj := basic{"test", 42}
	fmt.Println(obj)

	strict(myInt{42})
	//strict(42) //error due to strict typing

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

	dyn := newBasic("foo", 5)
	fmt.Println(dyn.name)
}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="6">An example of type copying. Not the same as an alias.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="10-12">This function is an exampe of strict types in Go. myInt is the same as a regular int but won't be accepted.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="33-34">Call</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="15-18">This is how structs are defined as types, no functions are written here.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="21-25">We can emulate constructors using a function that returns an object.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="24">Note that it is legal to return a pointer to a local variable. Go performs pointer escape analysis.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="30-31">Here we create an instance of the basic class, all structs members can be brace initialised.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="37-38,42">Go supports the creation of anonymous structs.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="46-49">Structs can be initialised out of order by naming the members.</span>

---

<span class="menu-title" style="display: none">Object oriented code in go</span>
### Object oriented code (ish) in Go

- Go is an object oriented language in that it allows the creation of structs that "inherit" from one and other, and also have member functions.
- Go also allows the creation of interfaces.
- Polymorphism in Go does not work like other languages, types that implement an interface automatically "inherit" from it.
- Inheritance is described by an is-a relationship, but in Go we go straight to polymorphism's acts-like relationship.

+++

<span class='menu-title slide-title'>oo_style_1.go</span>
```golang
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
func (d Dog) immutable(newStr string){
	d.noise = newStr
}

// This is allowed to modify the original object as we have used a pointer
func (d *Dog) mutable(newStr string){
	d.noise = newStr
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

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="5-9">Interfaces in Go are declared with the interface keyword, and only contain function declarations.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="16-18">Here we can see a dog member function. Defined with the special syntax before the function name.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="16">There are two flavours of this as we will soon see.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="11-18">By adding this member function, Dog now implements the Animal interface.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="34-39">This means we can store a Dog within an Animal interface variable.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="20-29">Member functions can either be defined with a pointer contract or with a value contract.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="20-24">This member function makes a copy of the object and acts on that data.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="26-29">This member function acts on the original object through a pointer.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="41-47">Hence the object is unchanged when calling the first, but modified when calling the second.</span>

---

<span class="menu-title" style="display: none">Interface contracts in go</span>
### Interface contracts in Go

- When passing an object to an interface parameter, a contract is formed.
- The function that receives the object checks how the member function calls have been bound.
- If the member functions have pointer contracts, a pointer must be passed to the interface.
- If the member functions have been bound by value, the interface can receive either a pointer or a value.


+++

<span class='menu-title slide-title'>oo_style_2.go</span>
```golang
package main

import "fmt"

type Animal interface{
	Noise() string
}

type Dog struct{}

//func (d Dog) Noise() string{
//	return "Pass by value! I mean Woof!"
//}

func (d *Dog) Noise() string {
	return "Pass by pointer! I mean Woof!"
}

// A function that takes an interface will base the contract,
// on how the functions have been implemented.
// If the interface function has been implemented with a pointer contract,
// a pointer must be provided.
// However if the function has a value contract, a pass by value is accepted,
// and a pointer is also accepted (automatically dereferenced)
func animalNoise(animal Animal) {
	fmt.Println(animal.Noise())
}

func main() {
	dog := Dog{}
	animalNoise(&dog)
}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="11-17">Here we can see the same function being bound with two different contracts.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="25-27">This function takes an interface and will determine the contract from the call to Noise.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="30-31">Here we call the function, this would not work without the &.</span>

---

<span class="menu-title" style="display: none">Composition and struct embedding in go</span>
### Composition and struct embedding in Go

- Go handles composition in a similar way to most languages.
- You can add a member to a struct, by supplying a name and type.
- Go also offers us struct embedding. Which is used by only giving a type.
- Embedding can be used to emulate inheritance from other languages, but is still composition.

+++

<span class='menu-title slide-title'>oo_style_3.go</span>
```golang
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

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="9-16">A simple struct with one member function.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="19-21">Driveway has a member of type Car, standard composition.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="39-42">Access to the members data and functions is through the member.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="24-28">Now we use struct embedding by omiting the members name.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="44-47">We can access Cars data and functions directly through the NewCar type.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="30-33">We can even embed a pointer, which can be used to interesting effect.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="49-53">If we pass a pointer to create a ProxyCar, changes made to the proxy will affect the original car, through the pointer. Similar to a reference.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="55-57">We can still access the embedded part of a struct, like we would with a non-embedded one. This is similar to casting to the base class in other OO languages.</span>

---

<span class="menu-title" style="display: none">Closures in go</span>
### Closures in Go

- Go has closures, similar to lambda functions in C++.
- Closures automatically capture all local variables, and go performs analysis on this so that they aren't destroyed when we exit that scope. 

+++

<span class='menu-title slide-title'>closures.go</span>
```golang
package main

import "fmt"

func triangleNumbers() (func() int) {
	diff := 1
	num := 0
	return func() int {
		num += diff
		diff++
		return num
	}
}

func main() {

	nextVal := triangleNumbers()

	for i := 0; i < 10; i++ {
		fmt.Println(nextVal())
	}
}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="5-13">This function returns a closure.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="6-7">The closure automatically captures these two variables.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="8-12">Using the variables here changes their scope, they won't get destroyed when this closure is returned.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="17-21">We get the returned closure and can call it successively to mutate the captured variables.</span>


---

<span class="menu-title" style="display: none">Generics in go</span>
### Generics in Go

- Go doesn't have generics like Java or C++.
- We can emulate them in a very useful and safe way using interfaces.
- As Go automatically makes types "inherit" from interfaces that they implement,
we can create interfaces to specify the types that should be passed to a function.
- The key is that we can request that passed types do something, without having to modify any types that could be passed.

+++

<span class='menu-title slide-title'>generics.go</span>
```golang
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

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="5-9">Here we define the interface that variables must satisfy to be passed into our function. In this case they must convert to a float32.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="26-32">Our function signature requires two Convertible types.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="12-23">Here I defined two structs which implement the Convertible interface.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="35-38">We can pass both of these objects to the function as they satisfy the interface.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="39-42">Can't pass a regular int as it doesn't satisfy the interface.</span>

+++

- Another example would be that all types passed to a sorting algorithm implement the Sortable interface, containing a less than function and are copy-able.

---

<span class="menu-title" style="display: none">Packages in go</span>
### Packages in Go

- Packages are Go's way of organising and encapsulating code.
- The main package must contain a file with a main function to run.
- Other packages can contain types, variables and functions to be used outside of that package.
- Every directory should contain at most one package that matches the directory name.
- Things can only be accessed from outside the package if they are exported, which is done by starting that data with a capital letter.

+++

<span class='menu-title slide-title'>pkg_demos/a_pkg.go</span>
```golang
package pkg_demos

func Exported() string {
	return "This is an Exported function!"
}

func Exported2() string {
	return "This is an Exported function calling " + nonExported()
}

func nonExported() string {
	return "a non-Exported function!"
}

type ExportedType struct {ExportedMember, nonExportedMember int}

type ExportedType2 struct {X nonExportedType}

type nonExportedType struct {x int}
```


Here we can see some types and functions defined in a package.

+++

<span class='menu-title slide-title'>pkg_demo.go</span>
```golang
package main

import (
	"fmt"
	pkgDemo "presentation_demos/pkg_demos"
	)

func main() {

	// Exported functions
	fmt.Println(pkgDemo.Exported())
	fmt.Println(pkgDemo.Exported2())

	// We can only use named intialisation with exported members
	a := pkgDemo.ExportedType{ExportedMember: 50}

	// We can't construct this member even though it is exported,
	// this is because it's type isn't exported
	b := pkgDemo.ExportedType2{}

	// We still get access to the member though
	fmt.Println(b.X)

	fmt.Println(a, b)

}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="5">We can use an alias to refer to an imported package.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="11-12">Only call functions with captial letter.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="15">We can only used named intialisation for members that are exported.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="19,22">We can't intialise a member if it an instance of a non-exported type, even if that member itself is exported. However we can still access the member.</span>

---

<span class="menu-title" style="display: none">Concurrency in go</span>
### Concurrency in Go

- Go gives us the keyword go to create subroutines.
- Placing go in-front of a statement will execute that concurrently.
- We are also given channels which can be used to synchronise routines and also send messages between them.

+++

<span class='menu-title slide-title'>concurrency.go</span>
```golang
package main

import (
	"log"
	"os"
	"fmt"
)

func simple(task int, logger *log.Logger) {
	for i := 0; i < 10; i++ {
		logger.Println(task, ":", i)
	}
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	for i := 0; i < 10; i++ {
		go simple(i, logger)
	}

	// This is to prevent the program from closing before the printing has finished.
	var input string
	fmt.Scanln(&input)
}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="9-13">Simple function to be executed concurrently.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="16">Logger is thread safe.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="17-19">Here we create 10 threads to print out consecutive integers.</span>

+++

<span class='menu-title slide-title'>channels.go</span>
```golang
package main

import (
	"fmt"
	"time"
)

func marcoPoloPrinter(receiver <-chan string, sender chan<- string, msg string) {
	// Loop endlessly
	for {
		// Wait here until we receive a message
		m := <-receiver
		// Print that message
		fmt.Println(m)
		// Send our message
		sender <- msg
		// Slow this down a bit
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// Make a channel for our polo message
	polo := make(chan string)
	// Make a channel for our marco message
	marco := make(chan string)

	// Set up to concurrent go routines to communicate
	go marcoPoloPrinter(marco, polo, "marco")
	go marcoPoloPrinter(polo, marco, "polo")

	// Trigger the start by sending the first message
	marco <- "marco"

	// So we can end at any time
	var input string
	fmt.Scanln(&input)
}

```


<span class="code-presenting-annotation fragment current-only" data-code-focus="23-26">Channels can be created using the chan keyword, the type specifier is for the message that will be sent and recieved.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="8-20">This function uses two channels, one for sending messages and one for receiving messages.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="8">Channels can be limited to only sending or only receiving messages, note the placement of <-. </span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="12">This line waits here until a message is received and then assigns it to m.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="16">This sends our message through the sender channel.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="28-30">By setting up to routines with reversed channels we can get them to ping pong against each-other.</span>
<span class="code-presenting-annotation fragment current-only" data-code-focus="32-33">This line sends the intial signal to start the pinging.</span>


---

<span class="menu-title" style="display: none">Resources</span>
### Resources

- [Golang documentation](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org)
- [Go by example](https://gobyexample.com/)
- [Stack overflow](https://stackoverflow.com/questions/tagged/go)

---

<span class="menu-title" style="display: none">References</span>
### References
- [Introduction to Go Structures and Data Instances, Karl Seguin, 2013](http://openmymind.net/Introduction-To-Go-Structures-Data-Instances/)
- [Introduction to programming in Go, Caleb Doxey, 2012](https://www.golang-book.com/books/intro)
- [Composite literals in Go, Michał Łowicki, 2016](https://medium.com/golangspec/composite-literals-in-go-10dc62eec06a)
- [Closures are the Generics for Go, Jon Bodner, 2017](https://medium.com/capital-one-developers/closures-are-the-generics-for-go-cb32021fb5b5)
- [Go interfaces & pointers, Saikiran Yerram, 2016](https://medium.com/@agileseeker/go-interfaces-pointers-4d1d98d5c9c6)
- [Types in the Go Programming Languag, Vladimir Vivien, 2016](https://medium.com/learning-the-go-programming-language/types-in-the-go-programming-language-65e945d0a692)
- [Why Go’s structs are superior to class-based inheritance, Ian Macalinao, 2016](https://medium.com/@simplyianm/why-gos-structs-are-superior-to-class-based-inheritance-b661ba897c67)