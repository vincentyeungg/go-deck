package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// verify all decks have 52 cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// verify first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first value of 'Ace of Spades', but got %v", d[0])
	}

	// verify last card is King of Clubs
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last value of 'King of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// need to remove any existing test file from prev tests if errors occurred
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	// remove the test file after assertion
	os.Remove("_decktesting")
}
