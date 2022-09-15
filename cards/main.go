package main

import "fmt"

var deckSize int

func main() {
	// newDeck().saveToFile("my_file")
	// newDeckFromFile("my_file").print()
	newDeck := deck(newDeck()[:5])
	newDeck.print()
	fmt.Println("---------------")
	newDeck.shuffle().print()
}
