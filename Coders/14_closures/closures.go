package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

func main() {
	// ----------- 1. Simple Closure -----------
	// Outer function
	counter := func() func() int {
		count := 0 // Captured variable
		return func() int {
			count++ // This function "remembers" count
			return count
		}
	}()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	// ----------- 2. Closure with Parameters -----------
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(5))
	fmt.Println(triple(5))
}
