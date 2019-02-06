package main

import "fmt"

func main() {
	cards := newDeck()
	hand := cards.deal(5)
	// fmt.Println(len(newCards))
	fmt.Println(len(cards))
	fmt.Println(len(hand))
	fmt.Println("----ORIGINAL----")
	cards.print()
	fmt.Println("----SHUFFLED----")
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")
	// newDeckFromFile("my").print()
}
