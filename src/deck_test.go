package main

import (
	"os"
	"testing"
)

// func TestNewDeck(t *testing.T) {
// 	d := newDeck()
// 	expectedDeckLength := 24
// 	if len(d) != expectedDeckLength {
// 		t.Errorf("Expected deck length of %v but got %v", expectedDeckLength, len(d))
// 	}

// 	expectedFirstCard := "Ace of Spades"
// 	if d[0] != expectedFirstCard {
// 		t.Errorf("Expected first card to be %v but was %v", expectedFirstCard, d[0])
// 	}

// 	expectedLastCard := "Six of Clovers"
// 	if d[len(d)-1] != expectedLastCard {
// 		t.Errorf("Expected last card to be %v but was %v", expectedLastCard, d[len(d)-1])
// 	}
// }

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	deckTestingFilename := "_decktesting"

	os.Remove(deckTestingFilename)

	// saves deck to new file
	nd := newDeck()
	nd.saveToFile(deckTestingFilename)

	lengthOfNewDeck := len(nd)

	rd := readFile(deckTestingFilename)

	if rd == nil {
		t.Errorf("Failed to read/save deck to hard drive")
	}

	if lengthOfNewDeck != len(rd) {
		t.Errorf("Deck read from disk has different length than the file created")
	}

	os.Remove(deckTestingFilename)

}
