package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {

	// creating map
	var m map[string]int
	fmt.Println("m : ", m) // map[]

	// MAP using make - string, string
	m1 := make(map[string]string)
	m1["name"] = "golang"
	m1["area"] = "backend"
	fmt.Println("m1 : ", m1)
	fmt.Println(`m1["not_existing"] : `, m1["not_existing"]) // zero value = ""
	// IMP: if key doesn't exist in map then it returns zero value
	// zero value means -> string ""   ;  integer = 0  ;  boolean = false

	// creating map using make
	m2 := make(map[string]int)
	m2["Age"] = 30
	m2["Salary"] = 50
	fmt.Println(`m2["Age"] : `, m2["Age"])       // m2["Age"] :  30
	fmt.Println(`m2["Salary"] : `, m2["Salary"]) // m2["Salary"] :  50
	fmt.Println(`m2["Leptop"] : `, m2["Leptop"]) // m2["Leptop"] :  0
	fmt.Println("m2 : ", m2)                     // m2 :  map[Age:30 Salary:50]
	fmt.Println("Length of map m2 : ", len(m2))  // Length of map m2 :  2
	fmt.Println(`delete(m2, "Age") : `)          //  delete key "Age"
	delete(m2, "Age")
	fmt.Println("m2 : ", m2)    // m2 :  map[Salary:50]     -> "Age" is deleted
	fmt.Println(`clear(m2) : `) // clear(m2) :  map[] -> empty map
	clear(m2)
	fmt.Println("m2 : ", m2) // m2 :  map[]

	// creating map using map literal
	m3 := map[string]int{"price": 500000, "phone": 40}
	fmt.Println(m3) // map[phone:40 price:500000]

	// useful feature
	v, ok := m3["price"]
	fmt.Println("v : ", v) // v : 500000
	if ok {
		fmt.Println("element found") // element found
	} else {
		fmt.Println("element NOT found")
	}

	//comparison
	m4 := map[string]int{"mango": 60, "apple": 120}
	m5 := map[string]int{"mango": 620, "apple": 120}
	fmt.Println(`maps.Equal(m4, m5) : `, maps.Equal(m4, m5)) // maps.Equal(m4, m5) :  false

}
