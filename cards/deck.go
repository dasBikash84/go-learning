package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Card struct {
	suit  string
	value string
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.value, c.suit)
}

type deck []Card

func (d deck) toStringArray() []string {
	res := []string{}
	for _, card := range d {
		res = append(res, card.String())
	}
	return res
}

func deserializeStrToDeck(strs []string) deck {
	var d deck
	for _, str := range strs {
		data := strings.Split(str, " of ")
		suit, value := data[1], data[0]
		d = append(d, Card{suit: suit, value: value})
	}
	return d
}

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

	cardSuits := []string{"S", "H", "D", "C"}

	cardValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	var outDeck deck

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			outDeck = append(outDeck, Card{value: cardValue, suit: cardSuit})
		}
	}

	return outDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

var srtSeparator = "@$%&"

func (d deck) toByteSlice() []byte {
	return []byte(strings.Join(d.toStringArray(), srtSeparator))
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
	return deserializeStrToDeck(strings.Split(string(data), srtSeparator))
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
