package main

import "fmt"

// Go is a static type language
// Reassigning variable - not use card : =
// Use card =

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")
	//cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	fmt.Println(cards)
}

func newCard() string {
	return "hello"
}

func print() {

}
