package main

import (
	"fmt"
	"unsafe"
	
)

func mainn() {
	a1 := 5
	var b1 = 10
	var c1 int = 15
	var d1 int
	d1 = 20
	fmt.Println("a1, b1, c1, d1 = ", a1, b1, c1, d1)

	fmt.Println(" ------------- data types in golang -------------")
	var a int16 = 2 //2 bytes
	var b uint16    //2 bytes
	var c float64   //8 bytes
	// var d string
	// var e bool
	fmt.Println("a, b, c = ", a, b, c)
	var f = 5          // 64 bit machine - 8 bytes
	var g = int(a) + f // 2 bytes + 8 bytes = 8 bytes
	fmt.Println("g = int(a) + f  : ", g)
	var h = a + int16(f) // 2 bytes + 2 bytes = 2 bytes   (allowed but loss of bits)
	fmt.Println("h = a + int16(f)  : ", h)

	fmt.Printf("Size of a: %d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("Size of f : %d bytes\n", unsafe.Sizeof(f))

	fmt.Println(" -------------control statements-------------")
	// if else
	score := 75
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 60 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// switch
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// loops - there is only 1 keyworkd i.e. for . No while or do while
	fmt.Println("---for loop---")
	for i := 0; i < 5; i++ {
		print("for loop: ", i, ", ")
	}
	fmt.Println("---while loop---")
	j := 0
	for j < 5 {
		print("while loop: ", j, ", ")
		j++
	}
	// infinite loop
	for {
		fmt.Println("This is an infinite loop")
		break // Add break to stop the loop
	}

	fmt.Println("---break,continue---")
	for i := 1; i <= 10; i++ {
		if i == 2 {
			continue // Skip 5
		}
		if i == 7 {
			break // Stop at 8
		}
		fmt.Println(i)
	}

	var name = "sushant"
	for i, val := range name {
		fmt.Printf("Index: %d, Value: %c\n", i, val)
	}
	// jump statements
	fmt.Println("---jump statements---")
	var no1 int = 5
	var no2 int = 6
sushant:
	// fmt.Println("enter no1")
	// fmt.Println("enter no2")
	// fmt.Scan(&no1)
	// fmt.Scan(&no2)
	no3 := no1 + no2
	fmt.Println("Sum: ", no3)
	if no3 < 0 {
		goto sushant
	}

	fmt.Println(" ------------- slice ------------- ")
	//Using make()
	slice1 := make([]int, 5) // Creates a slice of length 5 with default values (0)
	fmt.Println(slice1)      // Output: [0 0 0 0 0]

	//Using Literal Declaration
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2) // Output: [1 2 3 4 5]

	slice3 := []string{"Go", "Python", "Java"}
	for i, v := range slice3 {
		fmt.Printf("Index: %d, Value: %s || ", i, v)
	}

	var data1 [8]uint // even this is a datatype
	var data2 [9]uint // this is different datatype
	fmt.Printf("\ndata1 type: %T\n", data1)
	fmt.Printf("data2 type: %T\n", data2)

	//Slices, float, int, all datatype - in Go are passed_by_value
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Before Function Call:", slice)
	appendSlice(slice)
	fmt.Println("After Function Call:", slice)

	fmt.Println(" ------------- map ------------- ")

	// Using make()
	myMap1 := make(map[int]string)
	myMap1[100] = "Pineapple"
	myMap1[150] = "Apple"
	fmt.Println(myMap1)

	var myMap2 map[string]int
	myMap2 = map[string]int{"cat": 1, "bat": 2, "rat": 3}
	fmt.Println(myMap2)

	fmt.Println(" ------------- any ------------- ")

	// var data interface{} = 42

	// Type Assertion without checking
	// fmt.Println(data.(string)) // This will cause a panic

	fmt.Println(" ------------- functions------------- ")
	msg, result := addition("summ =", 1, 2, 3, 4, 5)
	fmt.Println(msg, result)

	f1 := multiply(3, 8)
	n := f1(5)
	fmt.Println(n)

	// Anonymous function
	func(a int, b int) {
		fmt.Println("Hello from Anonymous Function!")

		fmt.Println("Sum:", a+b)
	}(3, 5)

}

// functions
func multiply(a, b int) func(int) int {
	return func(x int) int {
		x++
		return x * b // Using x instead of shadowing a.
	}
}

func addition(s string, a, b int, d ...int) (string, int) {
	sum := a + b
	for _, val := range d {
		sum = sum + val
	}
	return s, sum
}

// Arrays - Pass by Value ;
// Slice - Pass by Reference.
//
//	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func appendSlice(s []int) {
	s = append(s, 100)
	s[2] = 200
	fmt.Println("Inside Function :", s)
}
