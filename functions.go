package main

import "fmt"

// No overloading
func addi(a, b int) int{
	return a + b
}

func addf(a, b float32) float32{
	return a + b
}

// Variadic args
func sum(args ...int) int{
	total := 0
	// range iterates over elements of containers, giving an index and element
	// here we use _ to say that we don't require the index
	for _, num := range args {
		total += num
	}
	return total
}

func multiReturn(a, b int) (int, int){
	return a*2, b*2
}

func main() {
	// Variadic args
	fmt.Println(sum(3, 2, 1))

	// We can convert a slice to args with ...
	args := []int{3, 2, 1}
	fmt.Println(sum(args...))


	first, second := multiReturn(1, 2)
	fmt.Println(first, second)
}
