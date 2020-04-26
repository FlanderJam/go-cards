package main

import "log"

func main() {
	buttDeck, err := newDeckFromFile("MyTest.txt")
	if err != nil {
		log.Fatal(err)
	}
	buttDeck.printDeck()
	cards := newDeck()
	cards.shuffleDeck()
	var hollyDeck deck
	var kevinDeck deck
	hollyDeck, cards = dealDeck(cards, 7)
	kevinDeck, cards = dealDeck(cards, 7)
	hollyDeck.printDeck()
	kevinDeck.printDeck()
}
