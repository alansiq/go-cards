package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedDeckLength := 24
	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v but got %v", expectedDeckLength, len(d))
	}

	expectedFirstCard := "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v but was %v", expectedFirstCard, d[0])
	}

	expectedLastCard := "Six of Clovers"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card to be %v but was %v", expectedLastCard, d[len(d)-1])
	}
}
