package main

import "fmt"

// iterating over data structures
func main() {
	nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Print(nums[i], " ") // 6 7 8
	// }

	sum := 0
	// Iterate over slice nums
	for i, n := range nums {
		fmt.Println("num, i : ", n, i)
		sum += n
	}
	fmt.Println("sum = ", sum)

	m := map[string]string{"fname": "sushant", "lname": "chauhan"}
	// Iterate over map m
	for k, v := range m {
		fmt.Println("k,v : ", k, v)
	}

	// Iterate over string "golangg"   [unicode code point rune ]
	// starting byte of rune
	// 255 -> 1 byte , 2 byte
	for i, ch := range "golangg" {
		fmt.Println("i, str : ", i, string(ch))
	}
}

// Output :
/*

PS C:\Users\SUSHANT CHAUHAN\Desktop\code\go-mastery\11_range> go run range.go
num, i :  6 0
num, i :  7 1
num, i :  8 2
sum =  21
k,v :  fname sushant
k,v :  lname chauhan
i, str :  0 g
i, str :  1 o
i, str :  2 l
i, str :  3 a
i, str :  4 n
i, str :  5 g
i, str :  6 g

*/
