package main

import "testing"

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
