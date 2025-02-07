package main

import "fmt"

//slice
func main() {
	//  ---------------- UnInitialized slice is nil ----------------
	/*
		var nums []int
		fmt.Println(nums)
		fmt.Println(nums==nil)
		fmt.Println(len(nums))
	*/
	// output :
	// []
	// true
	// 0

	// ----------------  make slice  ----------------
	// var nums = make([]int, 2)   // size = 2
	// fmt.Print(nums)           // [0 0]

	// ---------------------- Creating a slice with size-2 , Initial Capacity- 5   ----------------------
	/*
		var nums = make([]int, 2, 5)                                                               // length = 2 , initial capacity =  5
		fmt.Println("initial capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums) // 3rd parameter : capacity -> maximum number of elements can fit
		nums = append(nums, 10)
		nums = append(nums, 20)
		nums = append(nums, 30)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums)
		nums = append(nums, 40)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums)
		nums = append(nums, 50)
		nums = append(nums, 60)
		nums = append(nums, 70)
		nums = append(nums, 80)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums)
		nums = append(nums, 90)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums)
	*/
	// output :-
	// initial capacity = 5 , length =  2  , nums =  [0 0]
	// current capacity = 5 , length =  5  , nums =  [0 0 10 20 30]
	// current capacity = 10 , length =  6  , nums =  [0 0 10 20 30 40]
	// current capacity = 10 , length =  10  , nums =  [0 0 10 20 30 40 50 60 70 80]
	// current capacity = 20 , length =  11  , nums =  [0 0 10 20 30 40 50 60 70 80 90]

	// ---------------------- creating an empty slice with initial capacity-5 ----------------------
	/*
		var nums = make([]int, 0, 5)                                                               // length = 0 , initial capacity =  5
		fmt.Println("initial capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums) // 3rd parameter : capacity -> maximum number of elements can fit
		nums = append(nums, 10)
		nums = append(nums, 20)
		nums = append(nums, 30)
		nums = append(nums, 40)
		nums = append(nums, 50)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums) //  [0 0 10 20 30] , current capacity =  5
		nums = append(nums, 40)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums) //  [0 0 10 20 30 40] ,  current capacity =  10
		nums = append(nums, 50)
		nums = append(nums, 60)
		nums = append(nums, 70)
		nums = append(nums, 80)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums) // [0 0 10 20 30 40 50 60 70 80] ,  current capacity =  10
		nums = append(nums, 90)
		fmt.Println("current capacity =", cap(nums), ", length = ", len(nums), " , nums = ", nums) // [0 0 10 20 30 40 50 60 70 80 90] ,  current capacity =  20
	*/
	// ------- output : -------
	// initial capacity = 5 , length =  0  , nums =  []
	// current capacity = 5 , length =  5  , nums =  [10 20 30 40 50]
	// current capacity = 10 , length =  6  , nums =  [10 20 30 40 50 40]
	// current capacity = 10 , length =  10  , nums =  [10 20 30 40 50 40 50 60 70 80]
	// current capacity = 20 , length =  11  , nums =  [10 20 30 40 50 40 50 60 70 80 90]

	//------------------------- slice - other way (imp) ------------------------------------------
	/*	nums := []int{}
		nums = append(nums, 10)
		nums = append(nums, 20)
		nums[0] = 24 //replace 0th element
		fmt.Println(nums)
	*/
	//output :  [24 20]

	// -------------- copy slice --------------
	/*	var nums = make([]int, 0, 5)
		nums = append(nums, 21)

		var nums2 = make([]int, len(nums))
		copy(nums2, nums)

		fmt.Println(nums)    //[21]
		fmt.Println(nums2)   //[21]
	*/
	// -------------- slice operator --------------
	/*
		var nums = []int{1, 2, 3, 4, 5, 6}
		fmt.Println(nums[0:3]) // [1 2 3]
		fmt.Println(nums[:2])  // [1 2]
		fmt.Println(nums[1:])  // [2 3 4 5 6]
	*/
	// -------------- slice package --------------
	/*	var nums1 = []int{1, 2, 3}
		var nums2 = []int{1, 2, 3}
		fmt.Println(slices.Equal(nums1, nums2))    //true false
	*/

	// -------------- 2 D array --------------
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)

}
