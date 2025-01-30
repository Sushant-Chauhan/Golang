package main

import "fmt"

var deckSize1 int = 20

// deckSize2 := 40      # we just can't assign a value to it using the := syntax
var deckSize3 int

func main() {
	//--------------------------//--------------------------
	// var card string = "Ace of Spades"
	card1 := "Ace of Spades"
	fmt.Println(card1)
	card1 = "Five of Diamonds"
	fmt.Println(card1)
	// -------------------------
	fmt.Println(deckSize1)
	// fmt.Println(deckSize2)
	deckSize3 = 30
	fmt.Println(deckSize3)
	//--------------------------//--------------------------

	card := newCard()
	fmt.Println("card :", card)

	printState()
}

func newCard() string {
	return "Five of Diamonds"
}
