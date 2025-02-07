package main

import "fmt"

func main() {
	card1 := newCard()
	fmt.Println("card1 =", card1)

	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")
	// fmt.Println("cards = ", cards)

	cards.print()
}

func newCard() string {
	return "Five_of_Diamonds"
}
