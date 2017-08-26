package main

import (
	"fmt"
)

func main() {
	//These next two lines are the exact same
	// var card string = "Ace of Spades"
	//card := "Ace of Spades"
	cards := []string{newCard(), "Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
