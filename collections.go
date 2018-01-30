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

	// Creates a slice of 3 empty strings
	sliceStr := make([]string, 3)
	fmt.Println("Length:", len(sliceStr), "Elements:", sliceStr)

	// Map with string keys and int values
	var mapI map[string]int
	mapI = make(map[string]int)
	mapI["str"] = 12

	// Supports fairly complex types as keys by default
	var mapIArr map[[2]int]string
	mapIArr = make(map[[2]int]string)

	// Use complex key
	key := [2]int{0,1}
	mapIArr[key] = "str"

	// Get the data
	fmt.Println(mapI[mapIArr[key]])

	deduced := map[string][]int {
		"men":  {32, 55, 12, 55, 42, 53},
		"women":{44, 42, 23, 41, 65, 44},
	}

	fmt.Println(deduced["men"])
}
