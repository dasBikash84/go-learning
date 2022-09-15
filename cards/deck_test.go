package main

import (
	"os"
	"testing"
)

func TestCardDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("number of cards on deck is not 52")
	}
}

func TestSaveRemove(t *testing.T) {
	fileName := "_deckSaveReadTest"

	os.Remove(fileName)

	deck := newDeck()

	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(deck) != len(loadedDeck) {
		t.Errorf("%v cards saved but %v read back", len(deck), len(loadedDeck))
	}

	os.Remove(fileName)

}
