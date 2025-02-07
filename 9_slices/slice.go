package main

import "fmt"

//slice
func main() {
	// //uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums)        // []
	// fmt.Println(nums==nil)   // true
	// fmt.Println(len(nums))   // 0

	//// ---------------- --  make slice  ----------------
	// var nums = make([]int, 2)   // size = 2
	// fmt.Print(nums)           // [0 0]

	// ------ capacity resize ------
	var nums = make([]int, 2, 5)                  //size=2, initial capacity=5
	fmt.Println("initial capacity = ", cap(nums)) // 3rd parameter : capacity -> maximum number of elements can fit
	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)
	fmt.Println(nums, ",  current capacity = ", cap(nums)) //  [0 0 10 20 30] , current capacity =  5
	nums = append(nums, 40)
	fmt.Println(nums, ",  current capacity = ", cap(nums)) //  [0 0 10 20 30 40] ,  current capacity =  10
	nums = append(nums, 50)
	nums = append(nums, 60)
	nums = append(nums, 70)
	nums = append(nums, 80)
	fmt.Println(nums, ",  current capacity = ", cap(nums)) // [0 0 10 20 30 40 50 60 70 80] ,  current capacity =  10
	nums = append(nums, 90)
	fmt.Println(nums, ",  current capacity = ", cap(nums)) // [0 0 10 20 30 40 50 60 70 80 90] ,  current capacity =  20

}
