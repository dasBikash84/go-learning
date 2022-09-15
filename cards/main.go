package main

import "fmt"

// var deckSize int

func main() {
	// newDeck().saveToFile("my_file")
	// newDeckFromFile("my_file").print()
	newDeck := deck(newDeck()).shuffle()[:5]
	newDeck.print()
	newDeck.saveToFile("my_file")
	fmt.Println("---------------")
	newDeckFromFile("my_file").print()
}
