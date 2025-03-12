package main

import (
	"fmt"
	"reflect"
)

func sumfunc(nums ...interface{}) float64 {
	total := 0.0
	for _, num := range nums {
		switch n := num.(type) {
		case int:
			total += float64(n)
		case float64:
			total += n
		default:
			fmt.Println("Unsupported type: ", reflect.TypeOf(num))
		}
	}
	return total
}


func summ(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {

	res2 := sumfunc(3, 4, 5, 6.2, 7.5)
	fmt.Println("Result = ", res2)
 
	res1 := summ(4, 5, 6)
	fmt.Println("Result = ", res1)

	nums := []int{3,4,5,6}
	result := summ(nums...)
	fmt.Println(result)
}
