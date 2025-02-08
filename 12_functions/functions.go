package main

import "fmt"

func add(a, b int) int { // a int, b int
	return a + b
}

func getLanguages() (string, string, string, int) {
	return "Golang", "Javascript", "C++", 1000
}

func processIt(fn func(a int) int ) {
	fn(124)
}

func main() {
	fn1 := func(a int) int {
		return 2
	}
	processIt(fn1)


	result := add(14, 23)
	fmt.Println("Result = ", result) // Result =  37

	fmt.Println(getLanguages())           // way 1: Golang Javascript C++ 1000
	lng1, lng2, lng3, _ := getLanguages() // way 2 (getting 1000 value but i don't want to use it)
	fmt.Println(lng1, lng2, lng3)         // Golang Javascript C++
}
