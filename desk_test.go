package main

import (
	"os"
	"testing"
)

func TestNewDesk(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "A of ♠" {
		t.Errorf("Expected 1st card of A of ♠, but got %v", d[0])
	}

	if d[len(d)-1] != "4 of ♣" {
		t.Errorf("Expected last card of 4 of ♣, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeskAndNewDeckFromFile(t *testing.T) {
	os.Remove("_desktesting")

	desk := newDeck()
	desk.saveToFile("_desktesting")

	loadedDesk := newDeskFromFile("_desktesting")

	if len(loadedDesk) != 160 {
		t.Errorf("Expected 16 cards in desk, but got %v", len(loadedDesk))
	}

	os.Remove("_desktesting")
}
