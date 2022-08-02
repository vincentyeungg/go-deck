package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver function to be used on 'instance' of deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// deal a hand of custom size
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// utility function to turn deck to string type
func (d deck) toString() string {
	// converts the deck type to a comma separated string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
