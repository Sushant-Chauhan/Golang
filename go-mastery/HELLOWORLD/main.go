package main

import "fmt"

// deckSize2 := 40      # we just can't assign a value to it using the := syntax
var deckSize1 int = 20
var deckSize3 int

func main() {
	fmt.Println("Hello, I am THE main.go Indian!")

	// var card string = "Ace of Spades"
	fmt.Println("------------1.--------------")
	card1 := "Ace of Spades"
	fmt.Println(card1)
	card1 = "Five of Diamonds"
	fmt.Println(card1)
	fmt.Println(deckSize1)
	// fmt.Println(deckSize2)
	deckSize3 = 30
	fmt.Println(deckSize3)
	// printState()   -- method coming from state.go

	//Q. ('for .. in' is not allowed in golang)
	colors := []string{"red", "green", "blue"}
	for i, val := range colors {
		fmt.Println(i, val)
	}
	fmt.Println(colors)

}
