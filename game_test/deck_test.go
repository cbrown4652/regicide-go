package game_test

import (
	"testing"

	"github.com/cbrown4652/regicide-go/game"
)

func TestInitializeDeck(t *testing.T) {
	deck := game.Deck{}

	deck.NewDeck()

	expectedTavernLength := 40

	if len(deck.Tavern()) != expectedTavernLength {
		t.Errorf("Expected tavern deck length %d, but got %d", expectedTavernLength, len(deck.Tavern()))
	}

	expectedCastleLength := 12

	if len(deck.Castle()) != expectedCastleLength {
		t.Errorf("Expected castle deck length %d, but got %d", expectedCastleLength, len(deck.Castle()))
	}

	for i, card := range deck.Castle() {
		if i < 4 {
			if card.Rank != "J" {
				t.Errorf("Expected correct J castle group, but got %s", card.Rank)
			}
		} else if i < 8 {
			if card.Rank != "Q" {
				t.Errorf("Expected correct Q castle group, but got %s", card.Rank)
			}
		} else if i < 12 {
			if card.Rank != "K" {
				t.Errorf("Expected correct K castle group, but got %s", card.Rank)
			}
		}
	}
}
