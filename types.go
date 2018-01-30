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
