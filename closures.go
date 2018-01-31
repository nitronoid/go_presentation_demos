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
