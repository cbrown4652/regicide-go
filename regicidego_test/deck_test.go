package regicidego_test

import (
	"testing"

	"github.com/cbrown4652/regicide-go/regicidego"
)

func TestInitializeDeck(t *testing.T) {
	deck := regicidego.Deck{}

	deck.InitializeDeck()

	expectedTavernLength := 40

	if len(deck.GetTavern()) != expectedTavernLength {
		t.Errorf("Expected tavern deck length %d, but got %d", expectedTavernLength, len(deck.GetTavern()))
	}

	expectedCastleLength := 12

	if len(deck.GetCastle()) != expectedCastleLength {
		t.Errorf("Expected castle deck length %d, but got %d", expectedCastleLength, len(deck.GetCastle()))
	}

	for i, card := range deck.GetCastle() {
		if i < 4 {
			if card.GetRank() != "J" {
				t.Errorf("Expected correct J castle group, but got %s", card.GetRank())
			}
		} else if i < 8 {
			if card.GetRank() != "Q" {
				t.Errorf("Expected correct Q castle group, but got %s", card.GetRank())
			}
		} else if i < 12 {
			if card.GetRank() != "K" {
				t.Errorf("Expected correct K castle group, but got %s", card.GetRank())
			}
		}
	}
}
