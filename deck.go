package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// option 1: log the error and return a call to newDeck()
		// option 2: log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	// function shuffles the order of an existing deck

	// need to generate new source using seed for random number generator
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// swap the 2 elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
