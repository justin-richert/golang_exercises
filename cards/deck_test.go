package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedDeckLength := 52

	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedDeckLength, len(d))
	}

	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "King of Clubs"

	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card of %s, but got %s", expectedFirstCard, d[0])
	}

	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card of %s, but got %s", expectedLastCard, d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	expectedDeckLength := 52

	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedDeckLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
