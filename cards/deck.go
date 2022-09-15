package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) copyDeck() deck {
	var newDeck deck
	for i, card := range d {
		i = i + 0
		newDeck = append(newDeck, card)
	}
	return newDeck
}

func newDeck() deck {

	cardSuits := deck{"S", "H", "D", "C"}

	cardValues := deck{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	outDeck := deck{}

	for _, cardScardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			outDeck = append(outDeck, cardValue+" of "+cardScardSuit)
		}
	}

	return outDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

var srtSeparator = "@$%&"

func (d deck) toByteSlice() []byte {
	return []byte(strings.Join([]string(d), srtSeparator))
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, d.toByteSlice(), 0666)
}

func newDeckFromFile(fileName string) deck {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), srtSeparator))
}

func (d deck) shuffle() deck {

	souce := rand.NewSource(time.Now().UnixNano())
	r := rand.New(souce)

	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
	return d
}
